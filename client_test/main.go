package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/karelbilek/amqp-test-server/server"
	"github.com/streadway/amqp"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(fmt.Errorf("end of program"))

	runServer(ctx)
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:1234")
	if nil != err {
		panic(err)
	}
	qChan, err := conn.Channel()
	if nil != err {
		panic(err)
	}

	if err := qChan.ExchangeDeclare(
		"fooexchange", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // noWait
		nil,           // arguments
	); err != nil {
		panic(err)
	}

	_, err = qChan.QueueDeclare(
		"fooqueue", // name
		true,       // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if nil != err {
		panic(err)
	}
	err = qChan.QueueBind("fooqueue", "fookey", "fooexchange", false, nil)
	if nil != err {
		panic(err)
	}

	msgs, err := qChan.Consume("fooqueue", "consumer", false, false, false, false, nil)
	if nil != err {
		panic(err)
	}
	go func() {
		for d := range msgs {
			fmt.Println(string(d.Body))
		}
	}()

	err = qChan.Publish("fooexchange", "fookey", true, true, amqp.Publishing{
		Body:            []byte("hello"),
		UserId:          "guest",
		Headers:         amqp.Table{},
		ContentType:     "text/plain",
		ContentEncoding: "",
		Priority:        0,
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)

}

func runServer(ctx context.Context) {
	serverDbPath := "dispatchd-server.db"
	msgDbPath := "messages.db"
	server := server.NewServer(serverDbPath, msgDbPath, map[string]interface{}{}, false)
	lc := net.ListenConfig{}
	ln, err := lc.Listen(ctx, "tcp", ":1234")
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				conn, err := ln.Accept()
				if err != nil {
					fmt.Println(err)
					return
				}
				go handleConnection(server, conn)
			}
		}
	}()
}

func handleConnection(server *server.Server, conn net.Conn) {
	server.OpenConnection(conn)
}
