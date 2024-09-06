package myFunctions

func Mylow(str string) string {
	str1 := []rune(str)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			str1[i] += 32
		}
	}
	return string(str1)
}
