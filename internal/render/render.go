package render

import (
	"bytes"
	"html/template"

	"github.com/gobuffalo/helpers"
	"github.com/gobuffalo/helpers/forms"
	"github.com/gobuffalo/helpers/forms/bootstrap"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/plush"
)

type Data map[string]interface{}

type RenderFunc func(name string, body []byte, data Data) ([]byte, error)

func Plush(r RenderFunc, help hctx.Map) RenderFunc {
	return func(name string, body []byte, data Data) ([]byte, error) {
		if !HasExt(name, ".plush") {
			return r(name, body, data)
		}

		h := helpers.ALL()
		h[forms.FormKey] = bootstrap.Form
		h[forms.FormForKey] = bootstrap.FormFor
		h["form_for"] = bootstrap.FormFor

		for k, v := range help {
			h[k] = v
		}

		for k, v := range data {
			h[k] = v
		}

		ctx := plush.NewContextWith(h)

		s, err := plush.Render(string(body), ctx)
		if err != nil {
			return nil, err
		}
		name = StripExt(name, ".plush")
		return r(name, []byte(s), data)
	}
}

func GoTmpl(r RenderFunc) RenderFunc {
	return func(name string, body []byte, data Data) ([]byte, error) {
		if !HasExt(name, ".tmpl") {
			return r(name, body, data)
		}
		tmpl, err := template.New(name).Parse(string(body))
		if err != nil {
			return nil, err
		}
		bb := &bytes.Buffer{}

		if err := tmpl.Execute(bb, data); err != nil {
			return nil, err
		}
		return r(name, bb.Bytes(), data)
	}
}

func Noop(name string, b []byte, data Data) ([]byte, error) {
	return b, nil
}
