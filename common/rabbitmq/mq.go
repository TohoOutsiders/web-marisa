package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"server/common/setting"
)

type Mq struct {
	channel *amqp.Channel
}

func (m *Mq) Connect() error {
	conf := setting.Config.RabbitMQ

	address := fmt.Sprintf("amqp://%s:%s@%s:%d", conf.User, conf.Pass, conf.Addr, conf.Port)

	conn, err := amqp.Dial(address)
	if err != nil {
		return err
	}

	log.Println("Connect RabbitMQ Success")

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	m.channel = ch
	log.Println("[RabbitMQ] Channel -> []")

	mqConcif := &MqConfig{ch}
	mqConcif.AutoFunc()

	listener := &Listener{Channel: ch}
	listener.AutoFunc()

	return nil
}

func (m *Mq) Sender(exchange, routingkey, data string) {
	err := m.channel.Publish(
		exchange,
		routingkey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		log.Println("[RabbitMQ Error] Sender ->", routingkey, ":", err)
	}
	log.Println("[RabbitMQ] Sender ->", routingkey)
}
