package models

import (
	"strings"
	"time"
)

type ISO8601 time.Time

func (ct ISO8601) String() string {
	return time.Time(ct).Format("2006-01-02T15:04:05Z")
}

func (ct ISO8601) Time() time.Time {
	return time.Time(ct)
}

func (ct *ISO8601) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), "\"")
	if value == "" || value == "null" {
		return nil
	}
	t, err := time.Parse("20060102T150405Z", string(value))

	if err != nil {
		return err
	}
	*ct = ISO8601(t)
	return nil
}
