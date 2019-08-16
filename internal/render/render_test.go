package render

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Render(t *testing.T) {
	r := require.New(t)

	data := Data{
		"age":  42,
		"name": "mark",
	}
	rn := Plush(GoTmpl(Noop))
	b, err := rn("foo.plush.tmpl.html", []byte(html), data)
	r.NoError(err)

	r.Equal("Plush: 42 (MARK)\nGo: 42 (mark)", strings.TrimSpace(string(b)))
}

const html = `
Plush: <%= age %> (<%= upcase(name) %>)
Go: {{.age}} ({{.name}})
`
