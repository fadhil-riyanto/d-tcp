package utils

func SeeLinebreak(src string) string {
	var buf string

	for _, char := range src {
		if char == '\n' {
			buf = buf + "\\n"
		}
		if char == '\r' {
			buf = buf + "\\r"
		}
		buf = buf + string(char)
	}
	return buf
}
