/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:55
 * @File: memorise
 * @Desc:
 */
package models

type Memorise struct {
	Ip      string `gorm:"ip" form:"ip" json:"ip"`
	Keyword string `gorm:"keyword" form:"keyword" json:"keyword"`
	Answer  string `gorm:"answer" form:"answer" json:"answer"`
}
