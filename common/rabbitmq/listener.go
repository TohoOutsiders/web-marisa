package rabbitmq

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"reflect"
	"server/common/constant"
	"server/common/segment"
	"server/common/tools"
	"server/models"
	"server/repository"
	"strings"
)

type Listener struct {
	Repo    repository.IMemoriseRepo `inject:""`
	Channel *amqp.Channel
}

func (l *Listener) AutoFunc() {
	typeOf := reflect.TypeOf(l)
	valueOf := reflect.ValueOf(l)

	numOfMethod := valueOf.NumMethod()
	for i := 0; i < numOfMethod; i++ {
		if typeOf.Method(i).Name == "AutoFunc" {
			continue
		}
		valueOf.Method(i).Call(nil)
	}
}

func (l *Listener) MemoryAddListener() {
	message, err := l.Channel.Consume(
		constant.QueueNsMemoryAdd,
		"",
		true, //消息确认机制（Acknowlege) 自动ACK：消息一旦被接收，消费者自动发送ACK
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("[RabbitMQ Error] Listener -> MemoryAdd")
	}

	forever := make(chan bool)

	go func() {
		for msg := range message {
			var memory models.Memorise
			err := json.Unmarshal(msg.Body, &memory)
			if err != nil {
				log.Println("[RabbitMQ Error] Listener -> MemoryAdd:", err)
				continue
			}

			toPpl := segment.Init().Cut(memory.Keyword)
			memorise := l.Repo.FetchAllMemory()
			var real string

			if len(memorise) == 0 {
				real = tools.New().Join(toPpl, ",")
				goto DATA
			}

			for _, v := range memorise {
				ratio := 0
				keywords := strings.Split(v.Keyword, ",")
				for _, keyword := range keywords {
					for _, ppl := range toPpl {
						if keyword == ppl {
							ratio++
						}
					}
					if float32(ratio)/float32(len(keywords)) >= 0.6 {
						keywords = append(keywords, toPpl...)
						real = tools.New().Join(keywords, ",")
						goto DATA
					} else {
						real = tools.New().Join(toPpl, ",")
						goto DATA
					}
				}
			}
		DATA:
			data := make(map[string]interface{})
			data["ip"] = memory.Ip
			data["keyword"] = real
			data["answer"] = memory.Answer
			if l.Repo.AddMemory(data) {
				log.Println("队列消息，魔理沙学习完成...", data)
			}
		}
	}()

	log.Println("[RabbitMQ] Listener -> MemoryAdd")
	<-forever
}
