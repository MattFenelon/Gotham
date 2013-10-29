package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"testing"
)

type equalityTest struct {
	source      equaler
	sameAs      map[string]interface{}
	differentTo map[string]interface{}
}

type testType struct {
}

func buildTests() []equalityTest {
	tests := []equalityTest{}

	sourceComicId := domain.ParseComicId("ab5b2194-4090-48a4-8b8e-f66963908451")
	tests = append(tests, equalityTest{
		source: sourceComicId,
		sameAs: map[string]interface{}{
			"Same variable": sourceComicId,
			"Same value":    domain.ParseComicId("ab5b2194-4090-48a4-8b8e-f66963908451")},
		differentTo: map[string]interface{}{
			"Different value": domain.ParseComicId("3307b275-3a65-4587-9dc9-17c467564d16"),
			"Nil":             nil,
			"Another type":    testType{}}})

	id := uuid.NewRandom()
	sourceComicAdded := NewComicAdded(id, "SeriesTitle", "BookTitle")
	tests = append(tests, equalityTest{
		source: sourceComicAdded,
		sameAs: map[string]interface{}{
			"Same variable":      sourceComicAdded,
			"Deferenced pointer": *sourceComicAdded,
			"Same field values":  NewComicAdded(id, "SeriesTitle", "BookTitle")},
		differentTo: map[string]interface{}{
			"Different Id":          NewComicAdded(uuid.NewRandom(), "SeriesTitle", "BookTitle"),
			"Different SeriesTitle": NewComicAdded(id, "Different Series Title", "BookTitle"),
			"Different BookTitle":   NewComicAdded(id, "SeriesTitle", "Different Book Title"),
			"Nil":          nil,
			"Another type": testType{}}})

	return tests
}

func TestEquality(t *testing.T) {
	tests := buildTests()

	for _, test := range tests {
		source := test.source
		sameAs := test.sameAs
		differentTo := test.differentTo

		t.Logf("Testing equality for %T", source)
		for key, compareTo := range sameAs {
			if source.Equal(compareTo) == false {
				t.Errorf("%v: %+v and %+v should be equal but weren't", key, source, compareTo)
			}
		}

		for key, compareTo := range differentTo {
			if source.Equal(compareTo) {
				t.Errorf("%v: %v and %v should not be equal but were", key, source, compareTo)
			}
		}
	}
}
