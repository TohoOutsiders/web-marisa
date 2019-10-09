package rabbitmq

type IMq interface {
	Sender(exchange, routingkey, data string)
	Connect() error
}
