package myFunctions

func power(n1 int, n2 int) int {
	ret := 1
	for ; n2 > 0; n2-- {
		ret *= n1
	}
	return ret
}
