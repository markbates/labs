package strongparams

import (
	"net/http"
	"testing"

	"github.com/gobuffalo/httptest"
	"github.com/stretchr/testify/require"
)

func app(f http.HandlerFunc) http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", f)
	return mux
}

func Test_Permit(t *testing.T) {
	r := require.New(t)

	var params Params
	var err error

	w := httptest.New(app(func(res http.ResponseWriter, req *http.Request) {
		params, err = Permit(req, "foo")
	}))

	w.HTML("/?foo=bar&baz=bax").Get()
	r.NoError(err)
	r.Equal("bar", params.Get("foo"))
	r.Zero(params.Get("baz"))
}

func Test_Permit_Post(t *testing.T) {
	r := require.New(t)

	var params Params
	var err error

	w := httptest.New(app(func(res http.ResponseWriter, req *http.Request) {
		params, err = Permit(req, "foo")
	}))

	w.HTML("/").Post(map[string]string{
		"foo": "bar",
		"baz": "bax",
	})
	r.NoError(err)
	r.Equal("bar", params.Get("foo"))
	r.Zero(params.Get("baz"))
}
