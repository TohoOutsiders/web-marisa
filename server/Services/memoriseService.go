package Services

import (
	"web-marisa/server/Datasource"
	"web-marisa/server/Models"
)

// 插入记忆
func AddMemory(data map[string]interface{}) bool {
	var db = Datasource.GetInstace().GetMysqlDB()

	memory := Models.Memorise{
		Ip: data["ip"].(string),
		Keyword: data["keyword"].(string),
		Answer: data["answer"].(string),
	}
	db.Create(memory)

	return true
}

// 读取所有记忆
func FetchAllMemory() (memorise []Models.Memorise) {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Find(&memorise)
	return
}

// 读取一条记忆
func FetchMemory(answer string) (memorise Models.Memorise) {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Where("answer = ?", answer).First(&memorise)
	return
}

// 删除一条记忆
func DeleteMemoryByAnswer(answer string) bool {
	var db = Datasource.GetInstace().GetMysqlDB()
	db.Where("answer = ?", answer).Delete(Models.Memorise{})

	return true
}