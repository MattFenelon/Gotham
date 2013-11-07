package domain

import (
	"bytes"
	"code.google.com/p/go-uuid/uuid"
	"errors"
	"strings"
)

// A trimmedString represents a string that has been normalised to remove
// any whitespace before or after the string.
type trimmedString string

func NewTrimmedString(value string) trimmedString {
	return trimmedString(strings.TrimSpace(value))
}

func (t trimmedString) String() string {
	return string(t)
}

// A seriesTitle represents the name for a collection of comic books,
// e.g. Batman #1 belongs to the Batman series.
// Series titles cannot be empty nor can they be padded with whitespace.
type seriesTitle trimmedString

func NewSeriesTitle(value string) (seriesTitle, error) {
	if trimmed := NewTrimmedString(value); trimmed != "" {
		return seriesTitle(trimmed), nil
	}

	return "", errors.New("Series title cannot be empty")
}

func (s seriesTitle) String() string {
	return string(s)
}

// A bookTitle represents the title of a comic book.
// Book titles cannot be empty nor can they be padded with whitespace.
type bookTitle trimmedString

func NewBookTitle(value string) (bookTitle, error) {
	if trimmed := NewTrimmedString(value); trimmed != "" {
		return bookTitle(trimmed), nil
	}

	return "", errors.New("Book title cannot be empty")
}

func (b bookTitle) String() string {
	return string(b)
}

// A comicId represents the unique identity of a comic.
// A comicId is a 128 bit (16 byte) Universal Unique IDentifier as defined in RFC
// 4122.
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
		return a.equalTo(v)
	}

	if p, ok := b.(*comicId); ok {
		return a.equalTo(*p)
	}

	return false
}

func (a comicId) equalTo(b comicId) bool {
	return bytes.Equal(a, b)
}
