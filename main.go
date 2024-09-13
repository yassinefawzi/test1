package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-reloaded/myFunctions"
)

func skip_extra(s []rune, i int) int {
	for i < len(s) && (s[i] == ' ' || s[i] == '\r') {
		i++
	}
	i--
	return i
}

func check_if_closed(s []rune, i int) bool {
	i++
	for ; i < len(s); i++ {
		if s[i] == '(' {
			return false
		} else if s[i] == ')' {
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
		if runes[i] == '(' && check_if_closed(runes, i) {
			for ; i < len(runes) && runes[i] != ')'; i++ {
				holder = append(holder, runes[i])
			}
			ret = append(ret, (string(holder) + ")"))
			holder = []rune{}
		} else if runes[i] == ' ' {
			i = skip_extra(runes, i)
			if len(holder) >= 1 {
				ret = append(ret, string(holder))
				holder = []rune{}
			}
		} else if runes[i] == '\n' {
			if len(holder) >= 1 {
				ret = append(ret, string(holder))
				holder = []rune{}
			}
			ret = append(ret, "\n")
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
	for i < len(s) && (s[i] == ' ' || s[i] == '\r') {
		i++
	}
	return i
}

func copy_first(str string) string {
	s := []rune(str)
	ret := []rune{}
	jhold := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'A' {
			if i == 0 {
				if i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\r') {
					jhold = i
					i = skip_space(s, i+1)
					if i < len(s) && (s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'u' || s[i] == 'o' || s[i] == 'h') {
						ret = append(ret, s[jhold])
						ret = append(ret, 'n')
					} else if i < len(s) && (s[i] == 'A' || s[i] == 'E' || s[i] == 'I' || s[i] == 'U' || s[i] == 'O' || s[i] == 'H') {
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
			} else if i-1 > 0 && (s[i-1] == ' ' || s[i-1] == '\r') {
				if i+1 < len(s) && (s[i+1] == ' ' || s[i+1] == '\r') {
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

func check_flag(s []string) int {
	if s[0] == "(up" || s[0] == "(up)" {
		return 1
	} else if s[0] == "(low" || s[0] == "(low)" {
		return 2
	} else if s[0] == "(cap" || s[0] == "(cap)" {
		return 3
	} else if s[0] == "(hex)" {
		return 4
	} else if s[0] == "(bin)" {
		return 5
	} else {
		return -1
	}
}

func fix_punc(s []string) []string {
	var ret []string
	control := 1
	for i := 0; i < len(s); i++ {
		if len(s[i]) == 1 && s[i][0] == '\'' {
			if control%2 != 0 {
				if i+1 < len(s) {
					s[i+1] = "'" + s[i+1]
				} else {
					ret = append(ret, s[i])
				}
			} else if control%2 == 0 {
				if i-1 > 0 {
					ret[len(ret)-1] = ret[len(ret)-1] + "'"
				} else {
					ret = append(ret, s[i])
				}
			}
			control++
		} else {
			ret = append(ret, s[i])
		}
	}
	return ret
}

func ret_index(s string) int {
	holder := strings.Split(s, ",")
	if len(holder) == 1 {
		return 1
	} else {
		if holder[1][0] != ' ' || holder[1][len(holder[1])-1] != ')' {
			return -1
		}
		ret, err := strconv.Atoi(holder[1][1 : len(holder[1])-1])
		if err != nil {
			return -1
		}
		return ret
	}
}

func check_line(s []string, control int, stop int) bool {
	for ; control < stop; control++ {
		if s[control] == "\n" {
			return false
		}
	}
	return true
}

func run_it(s []string) []string {
	var ret []string
	flag := 0
	index := 0
	control := 0
	for i := 0; i < len(s); i++ {
		if s[i][0] == '(' && check_brack(s[i]) {
			index = ret_index(s[i])
			control = len(ret) - index
			if index == -1 {
				ret = append(ret, s[i])
			}
			flag = check_flag(strings.Split(s[i], ","))
			if control >= 0 && check_line(s, control, len(ret)) {
				if flag == 1 {
					for ; control < len(ret); control++ {
						ret[control] = myFunctions.Myup(ret[control])
					}
				} else if flag == 2 {
					for ; control < len(ret); control++ {
						ret[control] = myFunctions.Mylow(ret[control])
					}
				} else if flag == 3 {
					for ; control < len(ret); control++ {
						ret[control] = myFunctions.Mycap(ret[control])
					}
				} else if flag == 4 {
					ret[control] = myFunctions.Myhex(ret[len(ret)-1])
				} else if flag == 5 {
					ret[control] = myFunctions.Mybin(ret[len(ret)-1])
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
	return ret
}

func fix_mid_quote(s []string) []string {
	var ret []string
	for i := 0; i < len(s); i++ {
		if len(s[i]) > 1 && s[i][0] == '\'' {
			s[i] = s[i][1:]
			ret = append(ret, "'")
			ret = append(ret, s[i])
		} else if len(s[i]) > 1 && s[i][len(s[i])-1] == '\'' {
			s[i] = s[i][:len(s[i])-1]
			ret = append(ret, s[i])
			ret = append(ret, "'")
		} else {
			ret = append(ret, s[i])
		}
	}
	return ret
}

func fix_dot(s []string) []string {
	var ret []string
	i := 0
	for ; i < len(s); i++ {
		word := s[i]
		if len(word) > 0 && strings.ContainsAny(string(word[0]), ".,!?;:") && word[0] != '(' {
			punct := ""
			for _, char := range word {
				if strings.ContainsAny(string(char), ".,!?;:") {
					punct += string(char)
				} else {
					break
				}
			}
			if len(ret) > 0 {
				ret[len(ret)-1] += punct
			} else {
				ret = append(ret, punct)
			}
			if len(word) > len(punct) {
				ret = append(ret, word[len(punct):])
			}
		} else if len(word) > 0 && strings.ContainsAny(string(word), ".,!?;:") && word[0] != '(' {
			k := 0
			for j := 0; j < len(word); j++ {
				holder := ""
				for ; k < len(word) && !strings.ContainsAny(string(word[k]), ".,!?;:"); k++ {
					holder += string(word[k])
				}
				punct := ""
				for ; k < len(word) && strings.ContainsAny(string(word[k]), ".,!?;:"); k++ {
					punct += string(word[k])
				}
				ret = append(ret, holder+punct)
				j = k
				punct = ""
				holder = ""
			}
		} else {
			ret = append(ret, word)
		}
	}
	return ret
}

func to_one(s []string) string {
	ret := []rune{}
	for i := 0; i < len(s); i++ {
		ret = append(ret, []rune(s[i])...)
		if i+1 < len(s) && s[i][0] != '\n' {
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

func main() {
	if len(os.Args) != 3  {
		fmt.Printf("Error: Argument needed\n")
		return
	} else if !strings.Contains(string(os.Args[2]), ".txt") {
		fmt.Printf("Error: last file should be a .txt\n")
		return
	}
	contentBytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	content := string(contentBytes)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	content = copy_first(content)
	runes := []rune(content)
	splited_content := SplitWhiteSpaces(string(runes))
	splited_content = run_it(splited_content)
	splited_content = fix_mid_quote(splited_content)
	splited_content = fix_punc(splited_content)
	splited_content = fix_dot(splited_content)
	final := to_one(splited_content)
	err = os.WriteFile(os.Args[2], []byte(final), 0644)
    if err != nil {
        fmt.Println(err)
		return
    }
}
