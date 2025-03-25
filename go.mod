module github.com/karelbilek/amqp-test-server

go 1.23

toolchain go1.24.1

require (
	github.com/gogo/protobuf v1.2.1
	github.com/rcrowley/go-metrics v0.0.0-20190706150252-9beb055b7962
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	go.etcd.io/bbolt v1.4.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
)

require golang.org/x/sys v0.29.0 // indirect
