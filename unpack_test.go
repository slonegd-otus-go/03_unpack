package unpack_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	unpack "github.com/slonegd-otus-go/03_unpack"
)

func TestDo(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{`1`, `a4bc2d5e`, `aaaabccddddde`},
		{`2`, `abcd`, `abcd`},
		{`3`, `45`, ``},
		{`4`, `qwe\4\5`, `qwe45`},
		{`5`, `qwe\45`, `qwe44444`},
		{`6`, `qwe\\5`, `qwe\\\\\`},
	}
	for _, tt := range tests {
		got := unpack.Do(tt.in)
		assert.Equal(t, tt.want, got, tt.name)

		got = unpack.Parse(tt.in)
		assert.Equal(t, tt.want, got, tt.name)
	}
}
