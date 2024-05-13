package functions

func SplitString(s string, sep string) []string {
	result := []string{}
	token := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-len(sep) && s[i:i+len(sep)] == sep {
			result = append(result, token)
			token = ""
			i = i + len(sep) - 1
		} else {
			token += string(s[i])
		}
	}
	result = append(result, token)
	return result
}
