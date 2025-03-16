package models

import (
	"testing"
)

func TestISO8601String(t *testing.T) {
	ct := ISO8601{}
	expected := "0001-01-01T00:00:00Z"
	if ct.String() != expected {
		t.Errorf("Expected %s, got %s", expected, ct.String())
	}
}

func TestISO8601Time(t *testing.T) {
	ct := ISO8601{}
	expected := "0001-01-01 00:00:00 +0000 UTC"
	if ct.Time().String() != expected {
		t.Errorf("Expected %s, got %s", expected, ct.Time().String())
	}
}

func TestISO8601UnmarshalJSON(t *testing.T) {
	ct := ISO8601{}
	input := []byte(`"20250303T081000Z"`)
	expected := "2025-03-03T08:10:00Z"
	ct.UnmarshalJSON(input)
	if ct.String() != expected {
		t.Errorf("Expected %s, got %s", expected, ct.String())
	}
}
