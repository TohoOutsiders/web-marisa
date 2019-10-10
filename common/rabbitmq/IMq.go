package rabbitmq

type IMq interface {
	Sender(exchange, routingkey, data string)
	Delay(exchange, routingKey, data string, delayTime int)
	Connect() error
}
