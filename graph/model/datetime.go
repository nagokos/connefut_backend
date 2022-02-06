package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

/*
	APIサーバで受け取ったdatetimeフィールドの値をtime.Time型に変換する。
*/
func UnmarshalDateTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		if len(v) == 0 {
			return time.Time{}, nil
		}
		return time.Parse("2006/01/02 15:04", v)
	case time.Time:
		return v, nil
	default:
		return time.Now(), fmt.Errorf("DateTime is invalid")
	}
}

/*
	APIサーバからJSONを返す際に、time.Time型をstringに変換する。
*/
func MarshalDateTime(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		w.Write([]byte(strconv.Quote(t.Format("2006/01/02 15:04"))))
	})
}
