package persistence

type InMemoryFileStore struct {
	stored map[string][]string
}

func NewInMemoryFileStore() *InMemoryFileStore {
	return &InMemoryFileStore{
		stored: make(map[string][]string),
	}
}

func (f *InMemoryFileStore) Store(key string, filepaths []string) error {
	f.stored[key] = filepaths
	return nil
}

func (f *InMemoryFileStore) Get(key string) []string {
	return f.stored[key]
}
