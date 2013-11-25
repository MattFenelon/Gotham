package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"reflect"
	"testing"
)

func TestGetFrontpageViewWithASingleComic(t *testing.T) {
	t.Log("When a single comic is added with multiple pages")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	id := uuid.NewRandom()
	d.AddComic(id, "The Walking Dead", "The Walking Dead 115", []string{"0.jpg", "1.jpg"}, []string{"source//path//0.jpg", "source//path//1.jpg"})

	actual := d.GetFrontPageView()

	t.Log("The front page should only list the added comic")
	if len(actual.Series) != 1 {
		t.Errorf("\tExpected length of %v but was %v", 1, len(actual.Series))
	}

	t.Log("The front page should list the comic using its series' title")
	if actual.Series[0].Title != "The Walking Dead" {
		t.Errorf("\tExpected %v but was %v", "The Walking Dead", actual.Series[0].Title)
	}

	t.Log("The front page should use the comic's first page for the series image")
	if actual.Series[0].ImageKey != id.String()+"/0.jpg" {
		t.Errorf("\tExpected %v but was %v", id.String()+"/0.jpg", actual.Series[0].ImageKey)
	}
}

func TestGetFrontpageViewWithMultipleSeries(t *testing.T) {
	t.Log("When multiple comics from different series are added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	walkingDeadId := uuid.NewRandom()
	warriorId := uuid.NewRandom()
	d.AddComic(walkingDeadId, "The Walking Dead", "The Walking Dead 115", []string{"0.jpg"}, []string{"source//path//0.jpg"})
	d.AddComic(warriorId, "Warrior", "Warrior 1", []string{"0.jpg"}, []string{"source//path//0.jpg"})

	actual := d.GetFrontPageView()
	expected := &domain.FrontPageView{
		Series: []domain.FrontPageViewSeries{
			domain.FrontPageViewSeries{Title: "Warrior", ImageKey: warriorId.String() + "/0.jpg"},
			domain.FrontPageViewSeries{Title: "The Walking Dead", ImageKey: walkingDeadId.String() + "/0.jpg"},
		},
	}

	t.Log("The front page view should list all series in the order they were added")
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %+v but was %+v", expected, actual)
	}
}

func TestGetFrontpageViewWithMultipleComicsFromTheSameSeries(t *testing.T) {
	t.Log("When multiple comics from the same series are added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	firstId := uuid.NewRandom()
	lastId := uuid.NewRandom()
	d.AddComic(firstId, "The Walking Dead", "The Walking Dead 114", []string{"0.jpg"}, []string{"source//path//0.jpg"})
	d.AddComic(lastId, "The Walking Dead", "The Walking Dead 115", []string{"1.jpg"}, []string{"source//path//1.jpg"})

	actual := d.GetFrontPageView()

	t.Log("The front page should only list the series once")
	if len(actual.Series) != 1 {
		t.Errorf("\tExpected length of %v but was %v", 1, len(actual.Series))
	}

	t.Log("The front page should list the comic using its series' title")
	if actual.Series[0].Title != "The Walking Dead" {
		t.Errorf("\tExpected %v but was %v", "The Walking Dead", actual.Series[0].Title)
	}

	t.Log("The front page should use the last added comic's first page for the series image")
	if actual.Series[0].ImageKey != lastId.String()+"/1.jpg" {
		t.Errorf("\tExpected %v but was %v", lastId.String()+"/1.jpg", actual.Series[0].ImageKey)
	}
}

func TestGetFrontpageViewWithNoSeries(t *testing.T) {
	t.Log("When no comics have been added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)

	actual := d.GetFrontPageView()
	expected := &domain.FrontPageView{
		Series: []domain.FrontPageViewSeries{},
	}

	t.Log("The view's series list should be empty")
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %+v but was %+v", expected, actual)
	}
}
