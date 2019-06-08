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

func parse(i int, r rune, beforeRune *rune, res *string) {
	defer func() { *beforeRune = r }()

	if isNum(r) && i == 0 {
		return
	}

	if isNum(r) && isNum(*beforeRune) {
		return
	}

	if isNum(r) {
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
