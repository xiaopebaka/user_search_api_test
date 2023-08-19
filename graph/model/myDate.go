package model

import (
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"time"
)

func MarshalDate(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, t.Format(time.RFC3339))
	})
}

func UnmarshalDate(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339, tmpStr)
	}
	return time.Time{}, errors.New("Date must be a string formatted as RFC3339")
}
