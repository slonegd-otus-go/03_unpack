package unpack

import (
	"strconv"
	"strings"
)

func Do(in string) string {
	escaped = false
	escapedNum = false
	doubleEscaped = false
	lastRune = '\\'

	var res strings.Builder
	for i, r := range in {
		res.WriteString(parse(i, r))
	}
	return res.String()
}

var (
	escaped       bool
	escapedNum    bool
	doubleEscaped bool
	lastRune      rune
)

func parse(i int, r rune) string {
	defer func() {
		escapedNum = escaped && isNum(r)
		escaped = r == '\\'
		doubleEscaped = escaped && lastRune == '\\'
		lastRune = r
	}()

	if isNum(r) && i == 0 {
		return ""
	}

	if r == '\\' && !escaped {
		return ""
	}

	if isNum(r) && isNum(lastRune) && !escaped && !escapedNum {
		return ""
	}

	if isNum(r) && !escaped ||
		isNum(r) && escapedNum ||
		isNum(r) && doubleEscaped {
		repeat, _ := strconv.Atoi(string(r))
		var res strings.Builder
		for repeat > 1 {
			res.WriteRune(lastRune)
			repeat--
		}
		return res.String()
	}

	return string(r)
}

func isNum(r rune) bool {
	return r >= '0' && r <= '9'
}
