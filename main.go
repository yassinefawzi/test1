package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go-reloaded/myFunctions"
)

func skip_extra(s string, i int) int {
	for i < len(s) && (s[i] == ' ' || s[i] == '\t' || s[i] == '\n' || s[i] == '\r') {
		i++
	}
	i--
	return i
}

func check_if_closed(s string, i int) bool {
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
		if runes[i] == '(' && check_if_closed(s, i) {
			for ; i < len(runes) && runes[i] != ')'; i++ {
				holder = append(holder, runes[i])
			}
			ret = append(ret, (string(holder) + ")"))
			holder = []rune{}
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
	} else if s == "(hex)" {
		return 4
	} else if s == "(bin)" {
		return 5
	} else {
		return -1
	}
}

func ret_control(s string) int {
	holder := s[1 : len(s)-1]
	ret, Err := strconv.Atoi(holder)
	if Err != nil {
		return -1
	}
	return ret
}

func control_final_index(s []string, i int, control int) int {
	j := 0
	for ; j < control; i-- {
		if s[i][0] == '(' {
			j++
		}
		j++
	}
	return j
}

/*func run_it(s []string) []string {
	var ret []string
	// holder := []rune{}
	control := 0
	flag := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i][0] == '(' {
			if check_brack(s[i]) {
				run_holder := strings.Split(s[i], ",")
				flag = check_flag(run_holder[0])
				if len(run_holder) > 1 {
					control = ret_control(run_holder[1])
				} else {
					control = 1
				}
				if control > 0 {
					control = control_final_index(s, i-1, control)
					j = i - control
					if j >= 0 {
						if flag < 0 {
							ret = append(ret, s[i])
						} else if flag == 1 {
							for ; j < i; j++ {
								ret = append(ret, myFunctions.Myup(s[j]))
							}
						} else if flag == 2 {
							for ; j < i; j++ {
								if s[j][0] != '(' {
									fmt.Println("OK")
									ret = append(ret, myFunctions.Mylow(s[j]))
								}
							}
						} else if flag == 3 {
							for ; j < i; j++ {
								ret = append(ret, myFunctions.Mycap(s[j]))
							}
						} else if flag == 4 {
							ret = append(ret, myFunctions.Myhex(s[j]))
						} else if flag == 5 {
							ret = append(ret, myFunctions.Mybin(s[j]))
						}
					} else {
						fmt.Println("Error.")
						os.Exit(0)
					}
				} else {
					ret = append(ret, s[i])
				}
			}
		}
	}
	return ret
}*/

func fix_punc(s []string) []string {
	var ret []string
	control := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i][0] == '\'' && s[i+1][0] == '\'' {
			ret = append(ret, s[i])
		} else if len(s[i]) == 1 && s[i][0] == '\'' {
			if control%2 == 0 {
				if i+1 < len(s) {
					s[i+1] = "'" + s[i+1]
				} else {
					ret = append(ret, s[i])
				}
			} else if control%2 != 0 {
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

func run_it(s []string) []string {
	var ret 
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
	splited_content = fix_punc(splited_content)
	splited_content = run_it(splited_content)
	fmt.Printf("%v\n", splited_content)
}
