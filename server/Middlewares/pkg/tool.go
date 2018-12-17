package pkg

import (
	"bytes"
)

func DuplicateRemove(arr []string) (result []string) {
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
	return result
}

func Join(arr []string, cut string) (result string) {
	if len(arr) == 0 {
		return ""
	}

	var buffer bytes.Buffer
	for k, v := range arr {
		if (k != len(arr) - 1) {
			buffer.WriteString(v)
			buffer.WriteString(cut)
		} else {
			buffer.WriteString(v)
		}
	}

	return buffer.String()
}