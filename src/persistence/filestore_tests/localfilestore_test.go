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
	defer os.RemoveAll(path)

	fs := filestore.NewLocalFileStore(path)

	fs.Store("test_key", []string{"one", "two"}, []string{"testdata\\1.txt", "testdata\\2.txt"})

	keypath := path + "\\test_key"
	filesAtExpectedLocation, _ := ioutil.ReadDir(keypath)

	actualFiles := make([]string, 0, len(filesAtExpectedLocation))
	actualContents := make([][]byte, 0, len(filesAtExpectedLocation))
	for _, f := range filesAtExpectedLocation {
		actualFiles = append(actualFiles, f.Name())

		contents, err := ioutil.ReadFile(filepath.Join(keypath, f.Name()))
		if err != nil {
			t.Error(err)
		}
		actualContents = append(actualContents, contents)
	}

	t.Log("It should store the files using the filenames specified")
	expectedFiles := []string{"one", "two"}
	if reflect.DeepEqual(expectedFiles, actualFiles) == false {
		t.Errorf("\tExpected %v but was %v", expectedFiles, actualFiles)
	}

	t.Log("It should store the contents of the files")
	expectedContents := [][]byte{[]byte("1"), []byte("2")}
	if reflect.DeepEqual(expectedContents, actualContents) == false {
		t.Errorf("\tExpected %s but was %s", expectedContents, actualContents)
	}
}
