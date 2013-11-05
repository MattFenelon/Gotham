package domainservices

import (
	"domain"
)

type EventStorer interface {
	AddEvent(event *domain.ComicAdded) error
}

type FileStorer interface {
	Store(key string, filepaths []string) error
}
