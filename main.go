package main

import (
	"fmt"
	"os"
	"strings"
	"go-reloaded/myFunctions"
)

func skip_extra(s string, i int) int {
	if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
		for s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			i++
		}
		i--
	}
	return i
}

func check_for_brack(s string, i int) bool {
	holder := []rune(s)
	for ; i < len(holder); i++ {
		if holder[i] == '(' {
			return false
		} else if holder[i] == ')' {
			return true
		}
	}
	return false
}

func SplitWhiteSpaces(s string) []string {
	var ret []string
	var holder []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] == '(' && check_for_brack(s, i+1) {
			for ; i < len(runes) && runes[i] != ')'; i++ {
				holder = append(holder, runes[i])
			}
			holder = append(holder, ')')
			if i+1 < len(runes) {
				if len(holder) >= 1 {
					ret = append(ret, string(holder))
					holder = []rune{}
				}
			}
		} else if runes[i] == ' ' || runes[i] == '\t' || runes[i] == '\n' {
			i = skip_extra(string(runes), i)
			if len(holder) >= 1 {
				ret = append(ret, string(holder))
				holder = []rune{}
			}
		} else {
			holder = append(holder, runes[i])
		}
	}
	ret = append(ret, string(holder))
	return ret
}

func atoi(s string) int {
	ret := 0
	sign := 1
	if s[0] == '-' {
		sign *= -1
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			ret = ret*10 + int(s[i]-48)
		}
	}
	return ret * sign
}

func ret_flag(str string) int {
	if strings.Compare(str, "(cap") == 0 {
		return 1
	} else if strings.Compare(str, "(low)") == 0 {
		return 2
	} else if strings.Compare(str, "(up)") == 0 {
		return 3
	} else if strings.Compare(str, "(bin)") == 0 {
		return 4
	} else if strings.Compare(str, "(hex)") == 0 {
		return 5
	} else {
		return -1
	}
}

func loop_func(content []string) []string {
	holder := []string{}
	flag := 0
	num := 0
	j := 0
	for i := 0; i < len(content); i++ {
		if len(content[i]) > 1 && content[i][0] == '(' && check_for_brack(content[i], 1) {
			if strings.Contains(content[i], ",") {
				holder = strings.Split(content[i], ",")
				holder[0] += ")"
				flag = ret_flag(holder[0])
				num = atoi(holder[1])
			} else {
				flag = ret_flag(content[i])
				num = 1
			}
			if flag > 0 {
				j = i - num
				if j < 0 {
					return nil
				}
				if flag == 1 {
					for ; j < i; j++ {
						content[j] = myFunctions.Mycap(content[j])
					}
				} else if flag == 2 {
					for ; j < i; j++ {
						content[j] = myFunctions.Mylow(content[j])
					}
				} else if flag == 3 {
					for ; j < i; j++ {
						content[j] = myFunctions.Myup(content[j])
					}
				} else if flag == 4 {
					content[i-1] = myFunctions.Mybin(content[j])
				} else if flag == 5 {
					content[i-1] = myFunctions.Myhex(content[j])
				}
			}
		}
	}
	return content
}

func check(str string) int {
	if strings.Contains(str, "(") && check_for_brack(str, 1) {
		if strings.Contains(str, "(hex)") || strings.Contains(str, "(low") || strings.Contains(str, "(up") || strings.Contains(str, "(cap)") || strings.Contains(str, "(bin)") {
			return -1
		}
	}
	return 1
}

func check2(str string) int {
	if str[0] == '.' || str[0] == ',' || str[0] == '!' || str[0] == '?' || str[0] == ':' || str[0] == ';' {
		return -1
	}
	return 1
}

func is_vowel(str rune) bool {
	if str == 'a' || str == 'e' || str == 'i' || str == 'o' || str == 'u' {
		return true
	} else if str == 'A' || str == 'E' || str == 'I' || str == 'O' || str == 'U' {
		return true
	}
	return false
}

func check_for_punctuation(s string) string {
	str := []rune(s)
	ret := []rune{}
	flag := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '.' || str[i] == ',' || str[i] == '!' || str[i] == '?' || str[i] == ':' || str[i] == ';' {
			if i+1 < len(str) && (str[i+1] != '.' && str[i+1] != ',' && str[i+1] != '!' && str[i+1] != '?' && str[i+1] != ':' && str[i+1] != ';') {
				ret = append(ret, str[i])
				if str[i+1] != '\'' {
					ret = append(ret, ' ')
				}
				flag++
			}
		}
		if flag > 0 {
			flag = 0
		} else {
			ret = append(ret, str[i])
		}
	}
	return string(ret)
}

func final_loop(content []string) []rune {
	counter := 0
	ret := []rune{}
	for i := 0; i < len(content); i++ {
		j := check(content[i])
		if j == 1 {
			content[i] = check_for_punctuation(content[i])
			if i < len(content)-1 && (content[i][0] == 'a' || content[i][0] == 'A') && is_vowel([]rune(content[i+1])[0]) {
				if content[i] == "A" {
					ret = append(ret, 'A')
					ret = append(ret, 'n')
				} else {
					ret = append(ret, 'a')
					ret = append(ret, 'n')
				}
				ret = append(ret, ' ')
			} else {
				ret = append(ret, []rune(content[i])...)
				if content[i][0] != '\'' && counter == 0 {
					ret = append(ret, ' ')
				} else if i+1 < len(content) && counter > 0 && content[i+1][0] != '\'' {
					ret = append(ret, ' ')
				} else if content[i][0] == '\'' {
					counter++
				}
			}
		}
	}
	return ret
}

func parse(str []byte) string {
	ret := []rune{}
	s := []rune(string(str))
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i+1] == '(' && (s[i] != ' ' && s[i] != '\t' && s[i] != '\n') {
			ret = append(ret, ' ')
		} else if i > 0 && s[i-1] == ')' && (s[i] != ' ' && s[i] != '\t' && s[i] != '\n') {
			ret = append(ret, ' ')
			ret = append(ret, s[i])
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}

func parse_quotes(str string) string {
	ret := []rune{}
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if s[i] == '\'' {
			if i > 0 && (s[i] != ' ' && s[i] != '\t' && s[i] != '\n') {
				ret = append(ret, ' ')
			} else if i+1 < len(s) && (s[i] != ' ' && s[i] != '\t' && s[i] != '\n') {
				ret = append(ret, ' ')
			}
		}
		ret = append(ret, s[i])
	}
	return string(ret)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Error: 3 Arguments needed\n")
		return
	}
	content := []string{}
	var fcontent string
	Fcontent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fcontent = parse(Fcontent)
	fcontent = parse_quotes(fcontent)
	content = SplitWhiteSpaces(fcontent)
	content = loop_func(content)
	if content == nil {
		fmt.Printf("Error.\n")
		return
	}
	ret := final_loop(content)
	fmt.Printf("%s\n", string(ret))
}
