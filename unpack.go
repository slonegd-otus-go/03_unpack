package unpack

import "strconv"

func Do(in string) string {
	res := ""
	var beforRune rune

	for i, r := range in {
		parse(i, r, &beforRune, &res)
	}
	return res
}

var escaped = false
var escapedNum = false
var doubleEscaped = false

func parse(i int, r rune, beforeRune *rune, res *string) {
	defer func() {
		escapedNum = escaped && isNum(r)
		escaped = r == '\\'
		doubleEscaped = escaped && *beforeRune == '\\'
		*beforeRune = r
	}()

	if isNum(r) && i == 0 {
		return
	}

	if r == '\\' && !escaped {
		return
	}

	if isNum(r) && isNum(*beforeRune) && !escaped && !escapedNum {
		return
	}

	if isNum(r) && !escaped ||
		isNum(r) && escapedNum ||
		isNum(r) && doubleEscaped {
		repeat, _ := strconv.Atoi(string(r))
		for repeat > 1 {
			*res += (string(*beforeRune))
			repeat--
		}
		return
	}

	*res += string(r)
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}
