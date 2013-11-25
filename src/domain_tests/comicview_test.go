package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"reflect"
	"testing"
)

func TestGetComicView(t *testing.T) {
	t.Log("When a comic is added with multiple pages")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)

	id := uuid.NewRandom()
	d.AddComic(id, "The Saga of The Swamp Thing", "Swamp Thing 20", []string{"0.jpg", "1.jpg", "2.jpg"}, []string{"path//to//0.jpg", "path//to//1.jpg", "path//to//2.jpg"})

	actual := d.GetComicView(id)
	expected := &domain.ComicView{
		Pages: []string{"0.jpg", "1.jpg", "2.jpg"},
	}

	t.Log("The comic view should contain the comic's pages")
	if reflect.DeepEqual(expected, actual) == false {
		t.Errorf("\tExpected %v but was %v", expected, actual)
	}
}

func TestGetComicViewForAnUnknownComic(t *testing.T) {
	t.Log("When getting an unknown comic")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	actual := d.GetComicView(uuid.NewRandom())

	t.Log("It should return nil")
	if actual != nil {
		t.Errorf("\tExpected nil but was %v", actual)
	}
}
