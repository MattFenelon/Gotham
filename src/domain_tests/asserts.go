package domain_tests

import (
	"domain"
	"testing"
)

type equaler interface {
	Equal(interface{}) bool
}

func AssertEquality(t *testing.T, expected *domain.ComicAdded, actual *FakeEventStorer) {
	AssertCollectionEquality(t, []*domain.ComicAdded{expected}, actual)
}

func AssertCollectionIsEmpty(t *testing.T, actual *FakeEventStorer) {
	if len(actual.GetAllEvents()) == 0 {
		return
	}

	events := actual.GetAllEvents()
	actualValues := make([]domain.ComicAdded, 0, len(events))
	for _, event := range events {
		actualValues = append(actualValues, *event)
	}

	t.Errorf("\tCollection was expected to be empty but contained")
	t.Errorf("\t\t%#v\n", actualValues)
}

func AssertCollectionEquality(t *testing.T, expected []*domain.ComicAdded, actual *FakeEventStorer) {
	actualValues := make([]equaler, 0, len(actual.GetAllEvents()))
	expectedValues := make([]equaler, 0, len(expected))

	for _, expectedEvent := range expected {
		expectedValues = append(expectedValues, equaler(*expectedEvent))
	}

	error := false
	for _, actualEvent := range actual.GetAllEvents() {
		actual := *actualEvent
		actualValues = append(actualValues, equaler(actual))

		found := false
		for _, expectedValue := range expectedValues {
			if found = actual.Equal(expectedValue); found {
				break
			}
		}
		if found == false {
			error = true
		}
	}
	if error {
		t.Errorf("\tCollection was expected to only contain")
		t.Errorf("\t\t%v\n", expectedValues)
		t.Errorf("\t\tbut contained\n")
		t.Errorf("\t\t%v", actualValues)
	}
}
