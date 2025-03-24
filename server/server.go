package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/karelbilek/amqp-test-server/amqp"
	"github.com/karelbilek/amqp-test-server/binding"
	"github.com/karelbilek/amqp-test-server/exchange"
	"github.com/karelbilek/amqp-test-server/msgstore"
	"github.com/karelbilek/amqp-test-server/queue"
)

type Server struct {
	exchanges       map[string]*exchange.Exchange
	queues          map[string]*queue.Queue
	bindings        []*binding.Binding
	idLock          sync.Mutex
	conns           map[int64]*AMQPConnection
	db              *bolt.DB
	serverLock      sync.Mutex
	msgStore        *msgstore.MessageStore
	exchangeDeleter chan *exchange.Exchange
	queueDeleter    chan *queue.Queue
	users           map[string]User
	strictMode      bool
	ctx             context.Context
}

func (server *Server) MarshalJSON() ([]byte, error) {
	conns := make(map[string]*AMQPConnection)
	for id, value := range server.conns {
		conns[fmt.Sprintf("%d", id)] = value
	}
	return json.Marshal(map[string]interface{}{
		"exchanges":     server.exchanges,
		"queues":        server.queues,
		"connections":   conns,
		"msgCount":      server.msgStore.MessageCount(),
		"msgIndexCount": server.msgStore.IndexCount(),
	})
}

func NewServer(ctx context.Context, dbPath string, msgStorePath string, userJson map[string]interface{}, strictMode bool) *Server {
	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		panic(err.Error())

	}
	msgStore, err := msgstore.NewMessageStore(ctx, msgStorePath)
	msgStore.Start()
	if err != nil {
		panic("Could not create message store!")
	}

	var server = &Server{
		exchanges:       make(map[string]*exchange.Exchange),
		queues:          make(map[string]*queue.Queue),
		bindings:        make([]*binding.Binding, 0),
		conns:           make(map[int64]*AMQPConnection),
		db:              db,
		msgStore:        msgStore,
		exchangeDeleter: make(chan *exchange.Exchange),
		queueDeleter:    make(chan *queue.Queue),
		users:           make(map[string]User),
		strictMode:      strictMode,
		ctx:             ctx,
	}

	server.init(ctx)
	server.addUsers(userJson)
	return server
}

func (server *Server) init(ctx context.Context) {
	server.msgStore.LoadMessages() //this must be before initQueues
	server.initExchanges()
	server.initQueues(ctx)
	server.initBindings() // this must be after init{Exchanges,Queues}
	go server.exchangeDeleteMonitor()
	go server.queueDeleteMonitor()
}

func (server *Server) exchangeDeleteMonitor() {
	for {
		select {
		case e := <-server.exchangeDeleter:
			var dele = &amqp.ExchangeDelete{

				Exchange: e.Name,
				NoWait:   true,
			}
			server.deleteExchange(dele)
		case <-server.ctx.Done():
			return
		}
	}
}

func (server *Server) queueDeleteMonitor() {
	for {
		select {
		case q := <-server.queueDeleter:
			var delq = &amqp.QueueDelete{
				Queue:  q.Name,
				NoWait: true,
			}
			server.deleteQueue(delq, -1)
		case <-server.ctx.Done():
			return
		}
	}
}

func (server *Server) initBindings() {
	// Load bindings
	bindings, err := binding.LoadAllBindings(server.db)
	if err != nil {
		panic("Couldn't load bindings!")
	}
	for _, b := range bindings {
		// Get Exchange
		var exchange, foundExchange = server.exchanges[b.ExchangeName]
		if !foundExchange {
			panic("Couldn't bind non-existant exchange " + b.ExchangeName)
		}
		// Add Binding
		err = exchange.AddBinding(b, -1)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (server *Server) initQueues(ctx context.Context) {
	// Load queues
	queues, err := queue.LoadAllQueues(ctx, server.db, server.msgStore, server.queueDeleter)
	if err != nil {
		panic("Couldn't load queues!")
	}
	for _, queue := range queues {
		err = server.addQueue(queue)
		if err != nil {
			panic("Couldn't load queues!")
		}
	}
	// Load queue data
	for _, queue := range server.queues {
		queue.LoadFromMsgStore(server.msgStore)
	}
}

func (server *Server) initExchanges() {
	// LOAD FROM PERSISTENT STORAGE
	exchanges, err := exchange.LoadAllExchanges(server.db, server.exchangeDeleter)
	if err != nil {
		panic("Couldn't load exchanges!")
	}
	for _, ex := range exchanges {
		err = server.addExchange(ex)
		if err != nil {
			panic("Couldn't load queues!")
		}
	}
	if err != nil {
		panic("FAILED TO LOAD EXCHANGES: " + err.Error())
	}

	// DECLARE MISSING SYSEM EXCHANGES
	server.genDefaultExchange("", exchange.EX_TYPE_DIRECT)
	server.genDefaultExchange("amq.direct", exchange.EX_TYPE_DIRECT)
	server.genDefaultExchange("amq.fanout", exchange.EX_TYPE_FANOUT)
	server.genDefaultExchange("amq.topic", exchange.EX_TYPE_TOPIC)
}

func (server *Server) genDefaultExchange(name string, typ uint8) {
	_, hasKey := server.exchanges[name]
	if !hasKey {
		var ex = exchange.NewExchange(
			name,
			exchange.EX_TYPE_TOPIC,
			true,
			false,
			false,
			amqp.NewTable(),
			true,
			server.exchangeDeleter,
		)
		// Persist
		ex.Persist(server.db)
		err := server.addExchange(ex)
		if err != nil {
			panic(err.Error())
		}
	}
}

func (server *Server) addExchange(ex *exchange.Exchange) error {
	server.serverLock.Lock()
	defer server.serverLock.Unlock()
	server.exchanges[ex.Name] = ex
	return nil
}

func (server *Server) addQueue(q *queue.Queue) error {
	server.serverLock.Lock()
	defer server.serverLock.Unlock()
	server.queues[q.Name] = q
	var defaultExchange = server.exchanges[""]
	var defaultBinding, err = binding.NewBinding(q.Name, "", q.Name, amqp.NewTable(), false)
	if err != nil {
		return err
	}
	defaultExchange.AddBinding(defaultBinding, q.ConnId)
	q.Start()
	return nil
}

func (server *Server) deleteQueuesForConn(connId int64) {
	server.serverLock.Lock()
	var queues = make([]*queue.Queue, 0)
	for _, queue := range server.queues {
		if queue.ConnId == connId {
			queues = append(queues, queue)
		}
	}
	server.serverLock.Unlock()
	for _, queue := range queues {
		var method = &amqp.QueueDelete{
			Queue: queue.Name,
		}
		server.deleteQueue(method, connId)
	}
}

func (server *Server) deleteQueue(method *amqp.QueueDelete, connId int64) (uint32, uint16, error) {
	server.serverLock.Lock()
	defer server.serverLock.Unlock()
	// Validate
	var queue, foundQueue = server.queues[method.Queue]
	if !foundQueue {
		return 0, 404, errors.New("Queue not found")
	}

	if queue.ConnId != -1 && queue.ConnId != connId {
		return 0, 405, fmt.Errorf("Queue is locked to another connection")
	}

	// Close to stop anything from changing
	queue.Close()
	// Delete for storage
	bindings := server.bindingsForQueue(queue.Name)
	server.removeBindingsForQueue(method.Queue)
	server.depersistQueue(queue, bindings)

	// Cleanup
	numPurged, err := queue.Delete(method.IfUnused, method.IfEmpty)
	delete(server.queues, method.Queue)
	if err != nil {
		return 0, 406, err
	}
	return numPurged, 0, nil

}

func (server *Server) depersistQueue(queue *queue.Queue, bindings []*binding.Binding) error {
	return server.db.Update(func(tx *bolt.Tx) error {
		for _, binding := range bindings {
			if err := binding.DepersistBoltTx(tx); err != nil {
				return err
			}
		}
		return queue.DepersistBoltTx(tx)
	})
}

func (server *Server) bindingsForQueue(queueName string) []*binding.Binding {
	ret := make([]*binding.Binding, 0)
	for _, exchange := range server.exchanges {
		ret = append(ret, exchange.BindingsForQueue(queueName)...)
	}
	return ret
}

func (server *Server) removeBindingsForQueue(queueName string) {
	for _, exchange := range server.exchanges {
		exchange.RemoveBindingsForQueue(queueName)
	}
}

func (server *Server) deleteExchange(method *amqp.ExchangeDelete) (uint16, error) {
	server.serverLock.Lock()
	defer server.serverLock.Unlock()
	exchange, found := server.exchanges[method.Exchange]
	if !found {
		return 404, fmt.Errorf("Exchange not found: '%s'", method.Exchange)
	}
	if exchange.System {
		return 530, fmt.Errorf("Cannot delete system exchange: '%s'", method.Exchange)
	}
	exchange.Close()
	exchange.Depersist(server.db)
	// Note: we don't need to delete the bindings from the queues they are
	// associated with because they are stored on the exchange.
	delete(server.exchanges, method.Exchange)
	return 0, nil
}

func (server *Server) OpenConnection(network net.Conn) {
	c := NewAMQPConnection(server.ctx, server, network)
	server.serverLock.Lock()
	server.conns[c.id] = c
	server.serverLock.Unlock()
	c.openConnection()
}

func (server *Server) returnMessage(msg *amqp.Message, code uint16, text string) *amqp.BasicReturn {
	return &amqp.BasicReturn{
		Exchange:   msg.Method.Exchange,
		RoutingKey: msg.Method.RoutingKey,
		ReplyCode:  code,
		ReplyText:  text,
	}
}

func (server *Server) publish(exchange *exchange.Exchange, msg *amqp.Message) (*amqp.BasicReturn, *amqp.AMQPError) {
	// Concurrency note: Since there is no lock we can, technically, have messages
	// published after the exchange has been closed. These couldn't be on the same
	// channel as the close is happening on, so that seems justifiable.
	if exchange.Closed {
		if msg.Method.Mandatory || msg.Method.Immediate {
			var rm = server.returnMessage(msg, 313, "Exchange closed, cannot route to queues or consumers")
			return rm, nil
		}
		return nil, nil
	}
	queues, amqpErr := exchange.QueuesForPublish(msg)
	if amqpErr != nil {
		return nil, amqpErr
	}

	if len(queues) == 0 {
		// If we got here the message was unroutable.
		if msg.Method.Mandatory || msg.Method.Immediate {
			var rm = server.returnMessage(msg, 313, "No queues available")
			return rm, nil
		}
	}

	var queueNames = make([]string, 0, len(queues))
	for k, _ := range queues {
		queueNames = append(queueNames, k)
	}

	// Immediate messages
	if msg.Method.Immediate {
		var consumed = false
		// Add message to message store
		queueMessagesByQueue, err := server.msgStore.AddMessage(msg, queueNames)
		if err != nil {
			return nil, amqp.NewSoftError(500, err.Error(), 60, 40)
		}
		// Try to immediately consumed it
		for queueName, _ := range queues {
			qms := queueMessagesByQueue[queueName]
			for _, qm := range qms {
				queue, found := server.queues[queueName]
				if !found {
					// The queue must have been deleted since the queuesForPublish call
					continue
				}
				var oneConsumed = queue.ConsumeImmediate(qm)
				var rhs = make([]amqp.MessageResourceHolder, 0)
				if !oneConsumed {
					server.msgStore.RemoveRef(qm, queueName, rhs)
				}
				consumed = oneConsumed || consumed
			}
		}
		if !consumed {
			var rm = server.returnMessage(msg, 313, "No consumers available for immediate message")
			return rm, nil
		}
		return nil, nil
	}

	// Add the message to the message store along with the queues we're about to add it to
	queueMessagesByQueue, err := server.msgStore.AddMessage(msg, queueNames)
	if err != nil {
		return nil, amqp.NewSoftError(500, err.Error(), 60, 40)
	}

	for queueName, _ := range queues {
		qms := queueMessagesByQueue[queueName]
		for _, qm := range qms {
			queue, found := server.queues[queueName]
			if !found || !queue.Add(qm) {
				// If we couldn't add it means the queue is closed and we should
				// remove the ref from the message store. The queue being closed means
				// it is going away, so worst case if the server dies we have to process
				// and discard the message on boot.
				var rhs = make([]amqp.MessageResourceHolder, 0)
				server.msgStore.RemoveRef(qm, queueName, rhs)
			}
		}
	}
	return nil, nil
}

// Close closes all open connections
func (server *Server) Close() error {
	for _, conn := range server.conns {
		conn.hardClose()
	}
	server.serverLock.Lock()
	defer server.serverLock.Unlock()
	server.conns = make(map[int64]*AMQPConnection)
	return nil
}
