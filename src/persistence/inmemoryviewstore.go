package persistence

import (
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

func (vs *InMemoryViewStore) Get(key string, out interface{}) error {
	v := vs.views[key]
	if v == nil {
		return nil
	}

	// Reflection must be used to set the out parameter reflection is the only
	// route open to us to change the value pointed to by out.
	rv := reflect.ValueOf(v).Elem()
	outref := reflect.ValueOf(out).Elem()
	outref.Set(rv)

	return nil
}

func (v *InMemoryViewStore) Store(key string, in interface{}) error {
	v.views[key] = in
	return nil
}
