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
// The files parameter represents each file to be stored for the specified key.
// The key of the files map will be the filename that the file will be persisted as and
// value is the sourcepath where the file contents can be retrieved. The source filename
// is discarded once the file has been read.
//
// Implementations of Store should be atomic. Any errors should fail the entire operation.
type FileStorer interface {
	Store(key string, files map[string]string) error
}
