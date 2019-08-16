package render

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Exts(t *testing.T) {
	r := require.New(t)

	res := Exts("foo.a.b.c.d.e")
	r.Equal([]string{".a", ".b", ".c", ".d", ".e"}, res)
}

func Test_HasExt(t *testing.T) {
	r := require.New(t)

	r.True(HasExt("foo.a.b", ".b"))
	r.False(HasExt("foo.a.b", ".c"))
}

func Test_StripExt(t *testing.T) {
	r := require.New(t)

	n := StripExt("foo.a.b", ".a")
	r.Equal("foo.b", n)
}
