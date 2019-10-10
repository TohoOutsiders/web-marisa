package rabbitmq

import (
	"encoding/json"
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
	Repo repository.IMemoriseRepo
}

func (l *Listener) AutoFunc() {
	valueOf := reflect.ValueOf(l)
	numOfMethod := valueOf.NumMethod()
	for i := 1; i < numOfMethod; i++ {
		valueOf.Method(i).Call(nil)
	}
}

func (l *Listener) MemoryAddListener() {
	mq := &Mq{}
	mq.BaseListener(
		&MqConsumeConfig{
			Queue:     constant.QueueNsMemoryAdd,
			consumer:  "",
			autoAck:   true,
			exclusive: false,
			noLocal:   false,
			noWait:    false,
			args:      nil,
		},
		func(msg []byte) {
			var memory models.Memorise
			err := json.Unmarshal(msg, &memory)
			if err != nil {
				log.Println("[RabbitMQ Error] Listener -> MemoryAdd:", err)
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
		},
	)
	log.Println("[RabbitMQ] Listener -> MemoryAddListener")
}

func (l *Listener) MemoryDelListener() {
	mq := &Mq{}
	mq.BaseListener(
		&MqConsumeConfig{
			Queue:     constant.QueueNsMemoryDel,
			consumer:  "",
			autoAck:   true,
			exclusive: false,
			noLocal:   false,
			noWait:    false,
			args:      nil,
		},
		func(msg []byte) {
			if l.Repo.DeleteMemoryByAnswer(string(msg)) {
				log.Println("队列消息，魔理沙忘记...", string(msg))
			}
			log.Println("【Error】队列消息，魔理沙忘记失败...", string(msg))
		},
	)
	log.Println("[RabbitMQ] Listener -> MemoryDelListener")
}
