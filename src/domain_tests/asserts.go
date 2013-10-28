package domain_tests

import (
	"domain"
	"testing"
)

func AssertEquality(t *testing.T, expected *domain.ComicAdded, actual *FakeEventStorer) {
	AssertCollectionEquality(t, []*domain.ComicAdded{expected}, actual)
}

func AssertCollectionIsEmpty(t *testing.T, actual *FakeEventStorer) {
	if len(actual.events) == 0 {
		return
	}

	events := actual.events
	actualValues := make([]domain.ComicAdded, 0, len(events))
	for _, event := range events {
		actualValues = append(actualValues, *event)
	}

	t.Errorf("\tCollection was expected to be empty but contained")
	t.Errorf("\t\t%#v\n", actualValues)
}

func AssertCollectionEquality(t *testing.T, expected []*domain.ComicAdded, actual *FakeEventStorer) {
	// Write tests to test the assert methods
	actualValues := make([]domain.ComicAdded, 0, len(actual.events))
	expectedValues := make([]domain.ComicAdded, 0, len(expected))

	for _, expectedEvent := range expected {
		expectedValues = append(expectedValues, *expectedEvent)
	}

	error := false
	for _, actualEvent := range actual.events {
		actual := *actualEvent
		actualValues = append(actualValues, actual)

		found := false
		for _, expectedValue := range expectedValues {
			if actual == expectedValue {
				found = true
			}
		}
		if found == false {
			error = true
		}
	}
	if error {
		t.Errorf("\tCollection was expected to only contain")
		t.Errorf("\t\t%#v\n", expectedValues)
		t.Errorf("\t\tbut contained\n")
		t.Errorf("\t\t%#v", actualValues)
	}
}