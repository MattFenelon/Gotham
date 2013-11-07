package domainservices

import (
	"domain"
)

type EventStorer interface {
	AddEvent(event *domain.ComicAdded) error
}

// FileStorer is an interface for storing files against a key.
//
// Store makes no guarantees as to where the source files will be persisted, only that they
// will be retrievable by the specified key.
//
// The sourcePaths parameter is a list of the files to be stored against the key.
// The filename of the stored file will be changed to the match the filename at the same index
// in the filenames parameter.
//
// Implementations of Store should be atomic. Any errors should fail the entire operation.
type FileStorer interface {
	Store(key string, filenames, sourcePaths []string) error
}
