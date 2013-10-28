package domain

import (
	"errors"
	"strings"
)

type trimmedString string

func NewTrimmedString(value string) trimmedString {
	return trimmedString(strings.TrimSpace(value))
}

func (t trimmedString) String() string {
	return string(t)
}

type seriesTitle trimmedString

func NewSeriesTitle(value string) (seriesTitle, error) {
	if trimmed := NewTrimmedString(value); trimmed != "" {
		return seriesTitle(trimmed), nil
	}
	return "", errors.New("Series title's cannot be empty")
}

type bookTitle trimmedString

func NewBookTitle(value string) (bookTitle, error) {
	if trimmed := NewTrimmedString(value); trimmed != "" {
		return bookTitle(trimmed), nil
	}

	return "", errors.New("Book title's cannot be empty")
}
