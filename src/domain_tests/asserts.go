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

	if av.Len() != ev.Len() {
		error = true
	} else {
		for i := 0; i < av.Len(); i++ {
			if isEqual(av.Index(i), ev.Index(i)) == false {
				error = true
				break
			}
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
	a := actual.Interface()
	e := expected.Interface()

	aeq, aok := a.(equaler)
	eeq, eok := e.(equaler)

	if aok && eok {
		return aeq.Equal(eeq)
	}

	return reflect.DeepEqual(a, e)
}
