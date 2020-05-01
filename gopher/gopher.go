package gopher

func Plural(s string, count int) string {
	if count == 1 {
		return s
	}
	return s + "s"
}
