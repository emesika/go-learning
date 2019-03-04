package reverse

func String(input string) string {
	chars := []rune(input)
	result := ""
	for i := len(chars) - 1 ; i > -1 ; i-- {
		result += string(chars[i])
	}
	return result
}
