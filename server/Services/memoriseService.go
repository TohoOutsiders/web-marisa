package Services

import (
	"web-marisa/server/Datasource"
	"web-marisa/server/Models"
)

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
