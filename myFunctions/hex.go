package myFunctions

import "strconv"

func power(n1 int, n2 int) int {
	ret := 1
	for ; n2 > 0; n2-- {
		ret *= n1
	}
	return ret
}

func Myhex(str string) string {
	ret := 0
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			ret += int(str[i]) - '0'*power(16, (len(str)-1-i))
		} else if str[i] >= 'a' && str[i] <= 'f' {
			ret += int(str[i]-'a') - '0'*power(16, (len(str)-1-i))
		} else if str[i] >= 'A' && str[i] <= 'F' {
			ret += int(str[i]-'A') - '0'*power(16, (len(str)-1-i))
		} else {
			return ""
		}
	}
	ret_str := strconv.Itoa(ret)
	return ret_str
}
