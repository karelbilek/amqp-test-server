package msgstore

import (
	// "container/list"
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/karelbilek/amqp-test-server/amqp"
)

func TestWrite(t *testing.T) {
	// Setup
	var dbFile = "TestWrite.db"
	os.Remove(dbFile)
	defer os.Remove(dbFile)
	rhs := []amqp.MessageResourceHolder{&TestResourceHolder{}}
	ms, err := NewMessageStore(context.Background(), dbFile)
	// Create messages
	msg1 := amqp.RandomMessage(true)
	msg2 := amqp.RandomMessage(true)
	fmt.Printf("Creating ids: %d, %d\n", msg1.Id, msg2.Id)

	// Store messages and delete one
	fmt.Println("Adding message 1")
	ms.AddMessage(msg1, []string{"some-queue", "some-other-queue"})
	fmt.Println("Adding message 2")
	qm2Map, err := ms.AddMessage(msg2, []string{"some-queue"})
	_, err = ms.GetAndDecrRef(qm2Map["some-queue"][0], "some-queue", rhs)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// Close DB
	ms.db.Close()
	keys := map[int64]bool{
		msg1.Id: true,
		// msg2.Id: true,
	}
	// Assert that the DB is correct
	err = assertKeys(dbFile, keys)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// try loading from disk in the message store
	ms2, err := NewMessageStore(context.Background(), dbFile)
	_, err = ms2.LoadQueueFromDisk("some-queue")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	_, err = ms2.LoadQueueFromDisk("some-other-queue")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
