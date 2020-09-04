package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

// DateTime represents time in RFC3339 format
type DateTime time.Time

// MarshalGQL convert to GraphGL type
func (dt DateTime) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(time.Time(dt).Format(time.RFC3339))))
}

// UnmarshalGQL convert from Graphql type
func (dt *DateTime) UnmarshalGQL(v interface{}) error {
	dateTime, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong DateTime format")
	}

	t, err := time.Parse(time.RFC3339, dateTime)
	if err != nil {
		return fmt.Errorf("wrong DateTime format")
	}

	*dt = DateTime(t)
	return nil
}
