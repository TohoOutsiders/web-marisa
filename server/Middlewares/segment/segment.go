package segment

import (
	"github.com/huichen/sego"
	"sync"
)

type Segament struct {
}

var instance *Segament
var once sync.Once
var segmenter sego.Segmenter

func Init() *Segament {
	once.Do(func() {
		segmenter.LoadDictionary("Config/dictionary.txt")
		instance = &Segament{}
	})
	return instance
}

func (s *Segament) Cut(str string) []string {
	origin := []byte(str)
	segments := segmenter.Segment(origin)
	keyword := sego.SegmentsToSlice(segments, true)

	return keyword
}

