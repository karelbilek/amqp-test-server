package server

import (
	"github.com/karelbilek/amqp-test-server/util"
	"testing"
)

func TestImmediateFail(t *testing.T) {
	tc := newTestClient(t)
	defer tc.cleanup()
	conn := tc.connect()
	ch, retChan, _ := channelHelper(tc, conn)

	ch.QueueDeclare("q1", false, false, false, false, NO_ARGS)
	ch.QueueBind("q1", "abc", "amq.direct", false, NO_ARGS)
	ch.Publish("amq.direct", "abc", false, true, TEST_TRANSIENT_MSG)

	ret := <-retChan

	if ret.ReplyCode != 313 {
		t.Fatalf("Wrong reply code with Immediate return")
	}
	if string(ret.Body) != string(TEST_TRANSIENT_MSG.Body) {
		t.Fatalf("Did not get same payload back in BasicReturn")
	}
}

func TestImmediate(t *testing.T) {
	tc := newTestClient(t)
	defer tc.cleanup()
	conn := tc.connect()
	ch, _, _ := channelHelper(tc, conn)

	ch.QueueDeclare("q1", false, false, false, false, NO_ARGS)
	ch.QueueBind("q1", "abc", "amq.direct", false, NO_ARGS)

	deliveries, err := ch.Consume("q1", util.RandomId(), false, false, false, false, NO_ARGS)
	if err != nil {
		t.Fatalf("Failed to consume")
	}
	ch.Publish("amq.direct", "abc", false, true, TEST_TRANSIENT_MSG)
	<-deliveries
}

func TestMandatory(t *testing.T) {
	tc := newTestClient(t)
	defer tc.cleanup()
	conn := tc.connect()
	ch, retChan, _ := channelHelper(tc, conn)

	ch.Publish("amq.direct", "abc", false, true, TEST_TRANSIENT_MSG)

	ret := <-retChan

	if ret.ReplyCode != 313 {
		t.Fatalf("Wrong reply code with Mandatory return")
	}
	if string(ret.Body) != string(TEST_TRANSIENT_MSG.Body) {
		t.Fatalf("Did not get same payload back in BasicReturn")
	}
}
