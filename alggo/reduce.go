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
