package Models

import "web-marisa/server/Datasource"

type Memory struct {
	MemoryId int `gorm:"memoryId" form:"memoryId" json:"memoryId"`
	Ip string `gorm:"ip" form:"ip" json:"ip"`
	Keyword string `gorm:"keyword" form:"keyword" json:"keyword"`
	Answer string `gorm:"answer" form:"answer" json:"answer"`
}

func AddMemory(data map[string]interface{}) bool {
	var db = Datasource.GetInstace().GetMysqlDB()

	memory := Memory{

	}
}
