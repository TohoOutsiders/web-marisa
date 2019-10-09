/**
 * @Author: Tomonori
 * @Date: 2019/6/18 14:44
 * @File: tools
 * @Desc:
 */
package tools

import (
	"bytes"
	"sync"
)

type Tools struct {
}

var (
	Tool = New()
	once sync.Once
)

func New() (t *Tools) {
	once.Do(func() {
		t = &Tools{}
	})
	return t
}

func (t *Tools) DuplicateRemove(arr []string) (result []string) {
	for i := range arr {
		flag := true
		for j := range result {
			if arr[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, arr[i])
		}
	}
	return
}

func (t *Tools) Join(arr []string, cut string) string {
	if len(arr) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	for k, v := range arr {
		if k != len(arr)-1 {
			buffer.WriteString(v)
			buffer.WriteString(cut)
		} else {
			buffer.WriteString(v)
		}
	}

	return buffer.String()
}
