package domain

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
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

type comicId uuid.UUID

func NewComicId(id uuid.UUID) comicId {
	return comicId(id)
}

func ParseComicId(id string) comicId {
	if parsed := uuid.Parse(id); parsed != nil {
		return NewComicId(parsed)
	}

	return nil
}

func NewRandomComicId() comicId {
	return NewComicId(uuid.NewRandom())
}

func (id comicId) String() string {
	return uuid.UUID(id).String()
}

func (a comicId) Equal(b interface{}) bool {
	if v, ok := b.(comicId); ok {
		return a.EqualTo(v)
	}

	if p, ok := b.(*comicId); ok {
		return a.EqualTo(*p)
	}

	return false
}

func (a comicId) EqualTo(b comicId) bool {
	return bytes.Equal(a, b)
}
