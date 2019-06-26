package strongparams

import (
	"fmt"
	"net/http"
	"net/url"
)

type Params struct {
	url.Values
}

func Permit(req *http.Request, fields ...string) (Params, error) {
	p := New()
	req.ParseForm()
	fmt.Println("### req.PostForm ->", req.PostForm)

	for _, f := range fields {
		p.Add(f, req.URL.Query().Get(f))
		p.Add(f, req.PostForm.Get(f))
	}
	fmt.Println("### p ->", p)
	return p, nil
}

func Require(req *http.Request, fields ...string) (Params, error) {
	p := New()
	return p, nil
}

func New() Params {
	return Params{
		Values: url.Values{},
	}
}
