package alggo

import (
	"strconv"
	"strings"
)

func flattenLevelOne(arr []interface{}) string {
	var f []interface{}
	for _, e := range arr {
		switch v := e.(type) {
		case int:
			f = append(f, e)
		case []interface{}:
			for _, s := range v {
				f = append(f, s)
			}
		}
	}

	return formatFlatten(f)
}

func formatFlatten(f []interface{}) string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, e := range f {
		switch v := e.(type) {
		case int:
			sb.WriteString(strconv.Itoa(v))
			sb.WriteString(",")
		case []interface{}:
			sb.WriteString(formatFlatten(v))
		}
	}
	sb.WriteString("]")

	return sb.String()
}

func mostRepeated(sentence string) string {
	n := strings.ToUpper(sentence)
	words := strings.Fields(n)

	counter := make(map[string]uint)
	for _, v := range words {
		counter[v]++
	}

	var mr = struct {
		word   string
		amount uint
	}{
		word:   "",
		amount: 0,
	}

	for k, v := range counter {
		if v > mr.amount {
			mr.word = k
			mr.amount = v
		}
	}

	return mr.word
}

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}
