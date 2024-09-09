package main

import (
	"fmt"
	"strconv"
	"strings"

	//"go-reloaded/myFunctions"
	"os"
	//"strings"
)

func skip_extra(s string, i int) int {
	for i < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
		i++
	}
	i--
	return i
}

func SplitWhiteSpaces(s string) []string {
	var ret []string
	var holder []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' || runes[i] == '\t' || runes[i] == '\n' {
			i = skip_extra(string(runes), i)
			if len(holder) >= 1 {
				ret = append(ret, string(holder))
				holder = []rune{}
			}
		} else {
			holder = append(holder, runes[i])
		}
	}
	if len(holder) > 0 {
		ret = append(ret, string(holder))
	}
	return ret
}

func skip_space(s []rune, i int) int {
	for i < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
		i++
	}
	return i
}

func copy_first(str string) string {
	s := []rune(str)
	ret := []rune{}
	jhold := 0
	for i := 0; i < len(str); i++ {
		if i+1 < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
			jhold = i
			i = skip_space(s, i)
			if s[i] == '.' || s[i] == ',' || s[i] == '!' || s[i] == '?' || s[i] == ':' || s[i] == ';' {
				ret = append(ret, s[i])
			}
			if i < len(s) && (s[i] != '.' && s[i] != ',' && s[i] != '!' && s[i] != '?' && s[i] != ':' && s[i] != ';') {
				ret = append(ret, ' ')
				ret = append(ret, s[i])
			}
		} else if s[i] == 'a' || s[i] == 'A' {
			if i == 0 {
				if i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\t' || s[i+1] == '\n' || s[i+1] == '\r') {
					jhold = i
					i = skip_space(s, i+1)
					if i < len(s) && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'u' || s[i] == 'o' || s[i] == 'h') {
						ret = append(ret, s[jhold])
						ret = append(ret, 'n')
					} else {
						i = jhold
						ret = append(ret, s[i])
					}
					i = jhold
				} else {
					ret = append(ret, s[i])
				}
			} else if i-1 > 0 && (s[i-1] == ' ' || s[i-1] == '\t' || s[i-1] == '\n' || s[i-1] == '\r') {
				if i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\t' || s[i+1] == '\n' || s[i+1] == '\r') {
					jhold = i
					i = skip_space(s, i+1)
					if i < len(s) && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'u' || s[i] == 'o' || s[i] == 'h') {
						ret = append(ret, s[jhold])
						ret = append(ret, 'n')
					} else {
						i = jhold
						ret = append(ret, s[i])
					}
					i = jhold
				} else {
					ret = append(ret, s[i])
				}
			} else {
				ret = append(ret, s[i])
			}
		} else {
			ret = append(ret, s[i])
		}
	}
	return string(ret)
}

func check_brack(s string) bool {
	return s[len(s)-1] == ')'
}

func check_flag(s string) int {
	if s == "(up" || s == "(up)" {
		return 1
	} else if s == "(low" || s == "(low)" {
		return 2
	} else if s == "(cap" || s == "(cap)" {
		return 3
	} else if s == "(hex" || s == "(hex)" {
		return 4
	} else if s == "(bin" || s == "(bin)" {
		return 5
	} else {
		return -1
	}
}

func run_it(s []string) []string {
	var ret []string
	holder := []rune{}
	var run_holder []string
	flag := 0
	for i := 0; i < len(s); i++ {
		if s[i][0] == '(' {
			if check_brack(s[i]) {
				run_holder := strings.Split(s[i], ",")
				flag = check_flag(run_holder[0])
				if flag < 0 {
					ret = append(ret, s[i])
				}

			}
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Error: 3 Arguments needed\n")
		return
	}
	var content string
	Fcontent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	content = copy_first(string(Fcontent))
	splited_content := SplitWhiteSpaces(content)
	
	fmt.Printf("%s\n", splited_content)
}
