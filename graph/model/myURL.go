package model

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"net/url"
)

func MarshalURL(u url.URL) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
	})
}

func UnmarshalURL(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		u, err := url.Parse(v)
		if err != nil {
			return url.URL{}, err
		}
		return *u, nil
	case []byte:
		u := &url.URL{}
		if err := u.UnmarshalBinary(v); err != nil {
			return url.URL{}, err
		}
		return *u, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a url.URL", v)
	}
}
