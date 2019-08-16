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
	rn := Plush(GoTmpl(Noop), map[string]interface{}{
		"foo": func() int {
			return 3
		},
	})
	b, err := rn("foo.plush.tmpl.html", []byte(html), data)
	r.NoError(err)

	r.Equal("Plush: 42 (MARK / 3)\nGo: 42 (mark)", strings.TrimSpace(string(b)))
}

const html = `
Plush: <%= age %> (<%= upcase(name) %> / <%= foo() %>)
Go: {{.age}} ({{.name}})
`
