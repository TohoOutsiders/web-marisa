/**
 * @Author: Tomonori
 * @Date: 2019/6/18 15:37
 * @File: memoriseRepo
 * @Desc:
 */
package repository

import (
	"server/datasource"
	"server/models"
)

type MemoriseRepo struct {
	Source datasource.IDb `inject:""`
}

func (m *MemoriseRepo) AddMemory(data map[string]interface{}) bool {
	var db = m.Source.DB()
	memory := models.Memorise{
		Ip:      data["ip"].(string),
		Keyword: data["keyword"].(string),
		Answer:  data["answer"].(string),
	}
	db.Create(memory)
	return true
}

func (m *MemoriseRepo) FetchAllMemory() (memorise []models.Memorise) {
	var db = m.Source.DB()
	db.Order("memoryId desc").Find(&memorise)
	return
}

func (m *MemoriseRepo) FetchMemory(answer string) (memorise models.Memorise) {
	var db = m.Source.DB()
	db.Where("answer = ?", answer).First(&memorise)
	return
}

func (m *MemoriseRepo) DeleteMemoryByAnswer(answer string) bool {
	var db = m.Source.DB()
	db.Where("answer = ?", answer).Delete(models.Memorise{})
	return true
}
