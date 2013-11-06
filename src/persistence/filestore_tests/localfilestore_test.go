package filestore_tests

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"persistence/filestore"
	"reflect"
	"testing"
)

func TestStoringFiles(t *testing.T) {
	t.Log("When storing files")

	path, _ := ioutil.TempDir("", "filestore_test")
	t.Log(path)
	defer os.RemoveAll(path)

	fs := filestore.NewLocalFileStore(path)
	fs.Store("test_key", map[string]string{"one": "testdata\\1.txt", "two": "testdata\\2.txt"})

	keypath := path + "\\test_key"
	files, _ := ioutil.ReadDir(keypath)

	t.Log("It should store the files using the filenames specified")
	actualFiles := make([]string, 0, len(files))
	expectedFiles := []string{"one", "two"}
	for _, f := range files {
		actualFiles = append(actualFiles, f.Name())
	}
	if reflect.DeepEqual(expectedFiles, actualFiles) == false {
		t.Errorf("\tExpected %v but was %v", expectedFiles, actualFiles)
	}

	t.Log("It should store the contents of the files")
	actualContents := make([][]byte, 0, len(files))
	expectedContents := [][]byte{[]byte("1"), []byte("2")}
	for _, f := range files {
		contents, err := ioutil.ReadFile(filepath.Join(keypath, f.Name()))
		if err != nil {
			t.Error(err)
		}
		actualContents = append(actualContents, contents)
	}
	if reflect.DeepEqual(expectedContents, actualContents) == false {
		t.Errorf("\tExpected %s but was %s", expectedContents, actualContents)
	}
}
