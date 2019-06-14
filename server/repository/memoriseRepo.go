/**
 * @Author: Tomonori
 * @Date: 2019/6/11 17:05
 * @File: memoriseRepo
 * @Desc:
 */
package repository

import (
	"github.com/jinzhu/gorm"
	"server/Models"
)

type IMemoriseRepo interface {
	// 插入记忆
	AddMemory(data map[string]interface{}) bool
	// 读取所有记忆
	FetchAllMemory() (memorise []Models.Memorise)
	// 读取一条记忆
	FetchMemory(answer string) (memorise Models.Memorise)
	// 删除一条记忆
	DeleteMemoryByAnswer(answer string) bool
}

// new memory-based repository
func NewMemoriseRepo(source *gorm.DB) IMemoriseRepo {
	return &memoriseRepo{source}
}

type memoriseRepo struct {
	source *gorm.DB
}

func (m *memoriseRepo) AddMemory(data map[string]interface{}) bool {
	var db = m.source
	memory := Models.Memorise{
		Ip: data["ip"].(string),
		Keyword: data["keyword"].(string),
		Answer: data["answer"].(string),
	}
	db.Create(memory)
	return true
}

func (m *memoriseRepo) FetchAllMemory() (memorise []Models.Memorise) {
	var db = m.source
	db.Find(&memorise)
	return
}

func (m *memoriseRepo) FetchMemory(answer string) (memorise Models.Memorise) {
	var db = m.source
	db.Where("answer = ?", answer).First(&memorise)
	return
}

func (m *memoriseRepo) DeleteMemoryByAnswer(answer string) bool {
	var db = m.source
	db.Where("answer = ?", answer).Delete(Models.Memorise{})
	return true
}
