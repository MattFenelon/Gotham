package persistence

import (
	"errors"
	"reflect"
)

type InMemoryViewStore struct {
	views map[string]interface{}
}

func NewInMemoryViewStore() *InMemoryViewStore {
	return &InMemoryViewStore{
		views: map[string]interface{}{},
	}
}

func (v *InMemoryViewStore) Get(key string, out interface{}) error {
	view := v.views[key]
	if view == nil {
		return errors.New("View not found") // TODO: Replace with typed error
	}

	r := reflect.ValueOf(view)
	reflect.ValueOf(out).Set(r)

	return nil
}

func (v *InMemoryViewStore) Store(key string, in interface{}) error {
	v.views[key] = in
	return nil
}
