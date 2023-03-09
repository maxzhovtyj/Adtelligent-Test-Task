package concat

import (
	"strings"
)

func Basic(str []string) string {
	result := ""

	for _, v := range str {
		result += v
	}

	return result
}

func WithJoin(sl []string) string {
	return strings.Join(sl, "")
}

func WithStringBuilder(sl []string, strLen int) string {
	var q strings.Builder

	q.Grow(len(sl) * strLen)

	for _, v := range sl {
		q.WriteString(v)
	}

	return q.String()
}
