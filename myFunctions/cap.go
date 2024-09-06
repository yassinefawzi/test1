package myFunctions

func Mycap(str string) string {
	str1 := []rune(str)
	if str[0] >= 'a' && str[0] <= 'z' {
		str1[0] -= 32
	}
	return string(str1)
}
