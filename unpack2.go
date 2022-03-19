package unpack

import "unicode"

func Parse(in string) string {
	var parser parser
	var result string

	for _, r := range in {
		switch parser.state {
		case start:
			parser.start(r)

		case symbol:
			result += parser.symbol(r)

		case escape:
			parser.escape(r)

		case fail:
			return ""
		}
	}

	if parser.state == symbol {
		result += string(parser.last)
	}
	return result
}

type parser struct {
	state state
	last  rune
}

type state int

const (
	start  = iota // начало алгоритма
	symbol        // генерация строки (известен символ для генерации)
	escape        // был escape символ
	fail          // некорректная строка
)

func (p *parser) start(r rune) {
	if unicode.IsDigit(r) {
		p.state = fail
		return
	}

	if r == '\\' {
		p.state = escape
		return
	}

	p.last = r
	p.state = symbol
}

func (p *parser) symbol(r rune) string {
	if unicode.IsDigit(r) {
		p.state = start
		return repeat(p.last, int(r-'0'))
	}

	res := string(p.last)
	p.start(r)
	return res
}

func (p *parser) escape(r rune) {
	p.last = r
	p.state = symbol
}

func repeat(r rune, count int) string {
	res := make([]rune, 0, count)
	for i := 0; i < count; i++ {
		res = append(res, r)
	}
	return string(res)
}
