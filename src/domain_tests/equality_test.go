package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
	"testing"
	"time"
)

type equalityTest struct {
	source      equaler
	sameAs      map[string]interface{}
	differentTo map[string]interface{}
}

// testType is used for type comparison of the object-under-test with a different type.
type testType struct {
}

func buildTests() []equalityTest {
	tests := []equalityTest{}

	sourceComicId := model.ParseComicId("ab5b2194-4090-48a4-8b8e-f66963908451")
	tests = append(tests, equalityTest{
		source: sourceComicId,
		sameAs: map[string]interface{}{
			"Same variable": sourceComicId,
			"Same value":    model.ParseComicId("ab5b2194-4090-48a4-8b8e-f66963908451")},
		differentTo: map[string]interface{}{
			"Different value": model.ParseComicId("3307b275-3a65-4587-9dc9-17c467564d16"),
			"Nil":             nil,
			"Another type":    testType{}}})

	id := uuid.NewRandom()
	pages := []string{"0.jpg", "1.jpg"}
	writtenBy := []string{"Writer 1", "Writer 2"}
	artBy := []string{"Artist 1", "Artist 2"}
	publishedDate := time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC)
	sourceComicAdded := NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, pages, publishedDate, "Blurb")
	tests = append(tests, equalityTest{
		source: sourceComicAdded,
		sameAs: map[string]interface{}{
			"Same variable":      sourceComicAdded,
			"Deferenced pointer": *sourceComicAdded,
			"Same field values":  NewComicAdded(id, "SeriesTitle", "Title", []string{"Writer 1", "Writer 2"}, []string{"Artist 1", "Artist 2"}, []string{"0.jpg", "1.jpg"}, time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC), "Blurb"),
		},
		differentTo: map[string]interface{}{
			"Different Id":               NewComicAdded(uuid.NewRandom(), "SeriesTitle", "Title", writtenBy, artBy, pages, publishedDate, "Blurb"),
			"Different SeriesTitle":      NewComicAdded(id, "Different SeriesTitle", "Title", writtenBy, artBy, pages, publishedDate, "Blurb"),
			"Different Title":            NewComicAdded(id, "SeriesTitle", "Different Title", writtenBy, artBy, pages, publishedDate, "Blurb"),
			"Different Blurb":            NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, pages, publishedDate, "Different Blurb"),
			"Different PublishedDate":    NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, pages, time.Now(), "Blurb"),
			"Different Pages":            NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, []string{"2.jpg", "3.jpg"}, publishedDate, "Blurb"),
			"Different Pages order":      NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, []string{"1.jpg", "0.jpg"}, publishedDate, "Blurb"),
			"No pages":                   NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, []string{}, publishedDate, "Blurb"),
			"Different Pages (more)":     NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, []string{"0.jpg", "1.jpg", "2.jpg"}, publishedDate, "Blurb"),
			"Different Pages (less)":     NewComicAdded(id, "SeriesTitle", "Title", writtenBy, artBy, []string{"0.jpg"}, publishedDate, "Blurb"),
			"Different WrittenBy":        NewComicAdded(id, "SeriesTitle", "Title", []string{"W1", "W2"}, artBy, pages, publishedDate, "Blurb"),
			"Different WrittenBy order":  NewComicAdded(id, "SeriesTitle", "Title", []string{"Writer 2", "Writer 1"}, artBy, pages, publishedDate, "Blurb"),
			"No WrittenBy":               NewComicAdded(id, "SeriesTitle", "Title", []string{}, artBy, pages, publishedDate, "Blurb"),
			"Different WrittenBy (more)": NewComicAdded(id, "SeriesTitle", "Title", []string{"Writer 1", "Writer 2", "Writer 3"}, artBy, pages, publishedDate, "Blurb"),
			"Different WrittenBy (less)": NewComicAdded(id, "SeriesTitle", "Title", []string{"Writer 1"}, artBy, pages, publishedDate, "Blurb"),
			"Different ArtBy":            NewComicAdded(id, "SeriesTitle", "Title", writtenBy, []string{"A1", "A2"}, pages, publishedDate, "Blurb"),
			"Different ArtBy order":      NewComicAdded(id, "SeriesTitle", "Title", writtenBy, []string{"Artist 2", "Artist 1"}, pages, publishedDate, "Blurb"),
			"No ArtBy":                   NewComicAdded(id, "SeriesTitle", "Title", writtenBy, []string{}, pages, publishedDate, "Blurb"),
			"Different ArtBy (more)":     NewComicAdded(id, "SeriesTitle", "Title", writtenBy, []string{"Artist 1", "Artist 2", "Artist 3"}, pages, publishedDate, "Blurb"),
			"Different ArtBy (less)":     NewComicAdded(id, "SeriesTitle", "Title", writtenBy, []string{"Artist 1"}, pages, publishedDate, "Blurb"),
			"Nil":          nil,
			"Another type": testType{},
		},
	})

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
				t.Errorf("%v: Expected %+v to be equal to %+v", key, compareTo, source)
			}
		}

		for key, compareTo := range differentTo {
			if source.Equal(compareTo) {
				t.Errorf("%v: Expected %+v to not be equal to %+v", key, compareTo, source)
			}
		}
	}
}
