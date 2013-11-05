package domain_tests

import (
	"reflect"
	"testing"
)

type equaler interface {
	Equal(interface{}) bool
}

func AssertEquality(t *testing.T, expected interface{}, actual interface{}) {
	AssertCollectionEquality(t, []interface{}{expected}, actual)
}

func AssertCollectionIsEmpty(t *testing.T, actual interface{}) {
	a := reflect.ValueOf(actual)

	if a.Len() == 0 {
		return
	}

	t.Errorf("\tCollection was expected to be empty but contained")
	t.Errorf("\t\t%#v\n", actual)
}

func AssertCollectionEquality(t *testing.T, expected interface{}, actual interface{}) {
	error := false

	ev := reflect.ValueOf(expected)
	av := reflect.ValueOf(actual)

	for ai := 0; ai < av.Len(); ai++ {
		found := false
		for ei := 0; ei < ev.Len(); ei++ {
			if found = isEqual(av.Index(ai), ev.Index(ei)); found {
				break
			}
		}
		if found == false {
			error = true
		}
	}
	if error {
		t.Errorf("\tCollection was expected to only contain")
		t.Errorf("\t\t%v\n", expected)
		t.Errorf("\t\tbut contained\n")
		t.Errorf("\t\t%v", actual)
	}
}

func isEqual(actual, expected reflect.Value) bool {
	a := actual.Interface().(equaler)
	e := expected.Interface().(equaler)

	return a.Equal(e)
}
