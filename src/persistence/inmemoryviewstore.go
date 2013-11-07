package persistence

type InMemoryViewStore struct {
	views map[string]interface{}
}

func NewInMemoryViewStore() *InMemoryViewStore {
	return &InMemoryViewStore{
		views: map[string]interface{}{},
	}
}

func (v *InMemoryViewStore) Get(key string) interface{} {
	return v.views[key]
}

func (v *InMemoryViewStore) Store(key string, view interface{}) {
	v.views[key] = view
}
