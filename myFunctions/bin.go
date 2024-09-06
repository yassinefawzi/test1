package myFunctions

import "strconv"

func Mybin(str string) string {
	ret1 := 0
	num, _ := strconv.Atoi(str)
	ret := 0
	for num > 0 {
		ret = ret*10 + num%2
		num /= 2
	}
	for ret > 0 {
		ret1 = ret1*10 + ret%10
		ret /= 10
	}
	ret_str := strconv.Itoa(ret1)
	return ret_str
}
