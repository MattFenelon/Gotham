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

// ViewGetStorer is the interface that groups the Get and Store methods.
type ViewGetStorer interface {
	ViewGetter
	ViewStorer
}

// ViewStorer is the interface that wraps the Store method.
//
// Store takes a view with a specific key and creates or overwrites the view at that key.
type ViewStorer interface {
	Store(key string, view interface{})
}

// ViewGetter is the interface that wraps the Get method.
//
// Get retrieves a view that was previously stored at the specified key.
// If no view exists at the key, nil is returned.
type ViewGetter interface {
	Get(key string) interface{}
}
