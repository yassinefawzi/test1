package myFunctions

import (
	"strconv"
)

func Mybin(str string) string {
	ret := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '0' && str[i] != '1' {
			return ""
		}
		ret += (int(str[i]) -'0') * power(2, len(str)-i-1)
	}
	return strconv.Itoa(ret)
}
