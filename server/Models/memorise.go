package Models

type Memorise struct {
	Ip string `gorm:"ip" form:"ip" json:"ip"`
	Keyword string `gorm:"keyword" form:"keyword" json:"keyword"`
	Answer string `gorm:"answer" form:"answer" json:"answer"`
}
