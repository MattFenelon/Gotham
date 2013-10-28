package domain

import (
	"strings"
)

type trimmedString string

func NewTrimmedString(value string) trimmedString {
	return trimmedString(strings.TrimSpace(value))
}

func (t trimmedString) String() string {
	return string(t)
}
