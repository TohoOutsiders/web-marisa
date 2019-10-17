package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
	"reflect"
	"server/common/constant"
)

/**
 * 排序执行方法
 * 交换机声明一定在前
 */
type MqConfig struct {
	Channel *amqp.Channel
}

func (m *MqConfig) AutoFunc() {
	var (
		typeOf      = reflect.TypeOf(m)
		valueOf     = reflect.ValueOf(m)
		numOfMethod = valueOf.NumMethod()
	)

	for i := 0; i < numOfMethod; i++ {
		if typeOf.Method(i).Name == "AutoFunc" {
			continue
		}
		valueOf.Method(i).Call(nil)
	}
}

func (m *MqConfig) ExchangeMemory() {
	err := m.Channel.ExchangeDeclare(
		constant.ExchangeNsMemory,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Exchange Declare [%s]: %v", constant.ExchangeNsMemory, err)
	}
	log.Println("[RabbitMQ] Exchange ->", constant.ExchangeNsMemory)
}

func (m *MqConfig) QueueAdd() amqp.Queue {
	queue, err := m.Channel.QueueDeclare(
		constant.QueueNsMemoryAdd, //name
		false,                     //durable - 持久化
		false,                     //delete when unused - 队列清理
		false,                     //exclusive - 独立化
		false,                     //no-wait - 不堵塞
		nil,                       //arguments - 其他参数
	)
	if err != nil {
		log.Fatalf("Queue Declare [%s]: %v", constant.QueueNsMemoryAdd, err)
	}
	err = m.Channel.QueueBind(
		queue.Name,
		constant.QueueNsMemoryAdd,
		constant.ExchangeNsMemory,
		false,
		nil,
	)
	log.Println("[RabbitMQ] Queue ->", constant.QueueNsMemoryAdd)
	return queue
}

func (m *MqConfig) QueueDel() amqp.Queue {
	queue, err := m.Channel.QueueDeclare(
		constant.QueueNsMemoryDel, //name
		false,                     //durable - 持久化
		false,                     //delete when unused - 队列清理
		false,                     //exclusive - 独立化
		false,                     //no-wait - 不堵塞
		nil,                       //arguments - 其他参数
	)
	if err != nil {
		log.Fatalf("Queue Declare [%s]: %v", constant.QueueNsMemoryDel, err)
	}
	err = m.Channel.QueueBind(
		queue.Name,
		constant.QueueNsMemoryDel,
		constant.ExchangeNsMemory,
		false,
		nil,
	)
	log.Println("[RabbitMQ] Queue ->", constant.QueueNsMemoryDel)
	return queue
}
