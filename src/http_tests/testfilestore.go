package http_tests

import (
	"os"
	"persistence/filestore"
)

type testFileStore struct {
	path string
	*filestore.LocalFileStore
}

func newTestFileStore(path string) *testFileStore {
	return &testFileStore{
		path:           path,
		LocalFileStore: filestore.NewLocalFileStore(path),
	}
}

func (fs *testFileStore) close() {
	os.RemoveAll(fs.path)
}
