go-based amqp server that I want to use for testing amqp clients in go code.

Fork of github.com/ernestrc/dispatchd which is fork of github.com/dayorbyte/dispatchd

The Makefile doesn't work at all.

The context closure immediately closes all running goroutines; it's testable by uber goleak.

See client-test for example usage.