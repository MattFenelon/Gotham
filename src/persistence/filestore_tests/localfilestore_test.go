package filestore_tests

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"persistence/filestore"
	"reflect"
	"sort"
	"testing"
)

func TestStoringFiles(t *testing.T) {
	t.Log("When storing files")

	path, _ := ioutil.TempDir("", "filestore_test")
	defer os.RemoveAll(path)

	fs := filestore.NewLocalFileStore(path)
	fs.Store("test_key", []string{"one", "two"}, []string{filepath.Join("testdata", "1.txt"), filepath.Join("testdata", "2.txt")})

	t.Log("It should store the files using the specified filenames")

	f1, err := fs.Open(filepath.Join("test_key", "one"))
	defer f1.Close()
	if err != nil {
		t.Error(err)
	}
	actual, err := ioutil.ReadAll(f1)
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal([]byte("1"), actual) == false {
		t.Errorf("Expected %s but was %s", []byte("1"), actual)
	}

	f2, err := fs.Open(filepath.Join("test_key", "two"))
	defer f2.Close()
	if err != nil {
		t.Error(err)
	}
	actual, err = ioutil.ReadAll(f2)
	if err != nil {
		t.Error(err)
	}
	if bytes.Equal([]byte("2"), actual) == false {
		t.Errorf("\tExpected %s but was %s", []byte("2"), actual)
	}

	t.Log("It should list the key")
	actualKeys, _ := fs.GetAllKeys()
	expectedKeys := []string{"test_key"}
	if reflect.DeepEqual(actualKeys, expectedKeys) == false {
		t.Errorf("\tExpected %v but was %v", expectedKeys, actualKeys)
	}

	t.Log("It should list the filenames under the key")
	actualFilenames, _ := fs.GetFilenames("test_key")
	sort.Strings(actualFilenames)
	expectedFilenames := []string{"one", "two"}
	if reflect.DeepEqual(actualFilenames, expectedFilenames) == false {
		t.Errorf("\tExpected %v but was %v", expectedFilenames, actualFilenames)
	}
}
