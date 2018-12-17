package segment

import (
	"github.com/yanyiwu/gojieba"
	"sync"
)

type Segament struct {
}

var instance *Segament
var once sync.Once

func Init() *Segament {
	once.Do(func() {
		instance = &Segament{}
	})
	return instance
}

func (s *Segament) Cut(str string) []string {
	origin := str
	Ppl := gojieba.NewJieba()
	keyword := Ppl.Cut(origin, true)

	return keyword
}

