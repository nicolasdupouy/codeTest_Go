package gopher

import (
	"strings"
)

func Plural(s string, count int) string {
	if count == 1 {
		return s
	}
	return s + "s"
}

func PluralWithStringsBuilder(s string, count int) string {
	if count == 1 {
		return s
	}
	builder := strings.Builder{}
	builder.WriteString(s)
	builder.WriteString("s")
	return builder.String()
}
