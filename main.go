package main

import (
	"fmt"
	//"go-reloaded/myFunctions"
	"os"
	//"strings"
)

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
			fmt.Printf("%c\n", runes[i])
			holder = append(holder, runes[i])

		}
	}
	ret = append(ret, string(holder))
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
		if i+1 < len(s) && (str[i+1] == '(') {
			if s[i] != ' ' && s[i] != '\t' && s[i] != '\n' && s[i] != '\r' {
				ret = append(ret, s[i])
				ret = append(ret, ' ')
			} else {
				ret = append(ret, s[i])
			}
		} else if i+1 < len(s) && s[i] == ')' {
			if s[i+1] != ' ' && s[i+1] != '\t' && s[i+1] != '\n' && s[i+1] != '\r' {
				ret = append(ret, s[i])
				ret = append(ret, ' ')
			} else {
				ret = append(ret, s[i])
			}
		} else if i+1 < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
			if s[i+1] != '.' && s[i+1] != ',' && s[i+1] != '!' && s[i+1] != '?' && s[i+1] != ':' && s[i+1] != ';' {
				ret = append(ret, s[i])
			}
		} else if s[i] == 'a' || s[i] == 'A' {
			if i == 0 {
				if i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\t' || s[i+1] == '\n' || s[i+1] == '\r') {
					jhold = i
					i = skip_space(s, i+1)
					if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'u' || s[i] == 'o' || s[i] == 'h' {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Error: 3 Arguments needed\n")
		return
	}
	var content string
	//var fcontent string
	Fcontent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	content = copy_first(string(Fcontent))
	splited_content := SplitWhiteSpaces(content)
	fmt.Printf("%s\n", splited_content[len(splited_content)-1])
	/*fcontent = parse(Fcontent)
	fcontent = parse_quotes(fcontent)
	content = SplitWhiteSpaces(fcontent)
	content = loop_func(content)
	if content == nil {
		fmt.Printf("Error.\n")
		return
	}
	ret := final_loop(content)
	fmt.Printf("%s\n", string(ret))*/
}
