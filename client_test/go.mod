module github.com/karelbilek/amqp-test-server/client_test

go 1.24.1

replace github.com/karelbilek/amqp-test-server => ..

require (
	github.com/karelbilek/amqp-test-server v0.0.0-00010101000000-000000000000
	github.com/streadway/amqp v1.1.0
)

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962 // indirect
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4 // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
)
