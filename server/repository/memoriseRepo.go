/**
 * @Author: Tomonori
 * @Date: 2019/6/11 17:05
 * @File: memoriseRepo
 * @Desc:
 */
package repository

import (
	"server/Datasource"
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

type MemoriseRepo struct {
}

func (m *MemoriseRepo) AddMemory(data map[string]interface{}) bool {
	var db = Datasource.GetInstace().GetMysqlDB()
	memory := Models.Memorise{
		Ip: data["ip"].(string),
		Keyword: data["keyword"].(string),
		Answer: data["answer"].(string),
	}
	db.Create(memory)
	return true
}

func (m *MemoriseRepo) FetchAllMemory() (memorise []Models.Memorise) {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Find(&memorise)
	return
}

func (m *MemoriseRepo) FetchMemory(answer string) (memorise Models.Memorise) {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Where("answer = ?", answer).First(&memorise)
	return
}

func (m *MemoriseRepo) DeleteMemoryByAnswer(answer string) bool {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Where("answer = ?", answer).Delete(Models.Memorise{})
	return true
}
