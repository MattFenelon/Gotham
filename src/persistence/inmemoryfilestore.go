package persistence

type InMemoryFileStore struct {
	stored map[string][]string
}

func NewInMemoryFileStore() *InMemoryFileStore {
	return &InMemoryFileStore{
		stored: make(map[string][]string),
	}
}

func (f *InMemoryFileStore) Store(key string, files map[string]string) error {
	// TODO: Check key and files for nil

	names := make([]string, 0, len(files))
	for key, _ := range files { // The source paths can be ignored because the inmemorystore doesn't do anything with the files.
		names = append(names, key)
	}
	f.stored[key] = names
	return nil
}

func (f *InMemoryFileStore) Get(key string) []string {
	return f.stored[key]
}
