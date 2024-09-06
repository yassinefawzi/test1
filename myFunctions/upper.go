package myFunctions

func Myup(str string) string {
	str1 := []rune(str)
	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			str1[i] -= 32
		}
	}
	return string(str1)
}
