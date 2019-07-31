/**
 * @Author: Tomonori
 * @Date: 2019/6/18 15:37
 * @File: memoriseService
 * @Desc:
 */
package service

import (
	"server/models"
	"server/repository"
	"server/segment"
	"server/tools"
	"strings"
)

type MemoriseService struct {
	Repo repository.IMemoriseRepo `inject:""`
}

func (m *MemoriseService) Add(memory models.Memorise) map[string]interface{} {
	toPpl := segment.Init().Cut(memory.Keyword)
	memorise := m.Repo.FetchAllMemory()
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
	if m.Repo.AddMemory(data) {
		return data
	}
	return nil
}

func (m *MemoriseService) Reply(memory models.Memorise) (int, map[string]interface{}) {
	data := make(map[string]interface{})
	toPpl := segment.Init().Cut(memory.Keyword)
	memorise := m.Repo.FetchAllMemory()
	var answer string

	if len(memorise) == 0 {
		data["answer"] = "唔嗯...不懂你在说什么呢...教教我吧~"
		return 200, data
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
			if float32(ratio)/float32(len(keywords)) >= 0.4 {
				answer = v.Answer
				goto DATA
			}
		}
	}
	if answer == "" {
		data["answer"] = "唔嗯...不懂你在说什么呢...教教我吧~"
		return 200, data
	}
DATA:
	temp := m.Repo.FetchMemory(answer)
	data["answer"] = temp.Answer
	return 200, data
}

func (m *MemoriseService) Forget(answer string) bool {
	if m.Repo.DeleteMemoryByAnswer(answer) {
		return true
	}
	return false
}

func (m *MemoriseService) Status() int {
	memorise := m.Repo.FetchAllMemory()
	count := len(memorise)
	return count
}
