package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net"
	"server/common/setting"
	"server/repository"
	"time"
)

type Mq struct {
	Repo repository.IMemoriseRepo `inject:""`
}

type MqConsumeConfig struct {
	Queue                               string
	consumer                            string
	autoAck, exclusive, noLocal, noWait bool
	args                                amqp.Table
}

var (
	senderChannel, channel *amqp.Channel
)

func (m *Mq) Connect() error {
	conf := setting.Config.RabbitMQ

	address := fmt.Sprintf("amqp://%s:%s@%s:%d", conf.User, conf.Pass, conf.Addr, conf.Port)

	conn, err := amqp.DialConfig(
		address,
		amqp.Config{
			Heartbeat: 60 * time.Second,
			Dial: func(network, addr string) (conn net.Conn, e error) {
				return net.DialTimeout(network, addr, 18000*time.Second)
			},
		},
	)
	if err != nil {
		return err
	}

	log.Println("Connect RabbitMQ Success")

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	channel = ch
	log.Println("[RabbitMQ] Consumer Channel -> []")
	senderCh, err := conn.Channel()
	if err != nil {
		return err
	}
	senderChannel = senderCh
	log.Println("[RabbitMQ] Sender Channel -> []")

	mqConcif := &MqConfig{ch}
	mqConcif.AutoFunc()

	listener := &Listener{m.Repo}
	listener.AutoFunc()

	return nil
}

func (m *Mq) BaseListener(consumeConfig *MqConsumeConfig, fun func([]byte)) {
	go m.watcher(consumeConfig, fun)
}

func (m *Mq) Sender(exchange, routingkey, data string) {
	err := senderChannel.Publish(
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

func (m *Mq) Delay(exchange, routingKey, data string, delayTime int) {
	err := senderChannel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			Headers: amqp.Table{
				"x-delay": delayTime,
			},
			ContentType: "text/plain",
			Body:        []byte(data),
		},
	)
	if err != nil {
		log.Println("[RabbitMQ Error] Sender ->", routingKey, ":", err)
	}
	log.Println("[RabbitMQ] Sender ->", routingKey, " Delay Time:", delayTime)
}

func (m *Mq) watcher(consumeConfig *MqConsumeConfig, fun func([]byte)) {
	message, err := channel.Consume(
		consumeConfig.Queue,
		consumeConfig.consumer,
		consumeConfig.autoAck,
		consumeConfig.exclusive,
		consumeConfig.noLocal,
		consumeConfig.noWait,
		consumeConfig.args,
	)
	if err != nil {
		log.Fatal("[RabbitMQ Error] Listener ->", consumeConfig.Queue, ":", err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range message {
			fun(msg.Body)
		}
	}()
	<-forever
}
