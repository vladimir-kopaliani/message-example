package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

const dateFormat = "2006-01-02"

// Date is foramt which contains year, month and day
type Date time.Time

// MarshalGQL convert to GraphGL type
func (dt Date) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(time.Time(dt).Format(dateFormat))))
}

// UnmarshalGQL convert from Graphql type
func (dt *Date) UnmarshalGQL(v interface{}) error {
	dateTime, ok := v.(string)
	if !ok {
		return fmt.Errorf("wrong DateTime format")
	}

	t, err := time.Parse(dateFormat, dateTime)
	if err != nil {
		return fmt.Errorf("wrong DateTime format")
	}

	*dt = Date(t)
	return nil
}
