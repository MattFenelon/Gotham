package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domainservices"
	"reflect"
	"testing"
)

func TestGetFrontpageViewWithASingleComic(t *testing.T) {
	t.Log("When a single comic is added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domainservices.NewComicDomain(es, fs, vs)
	d.AddComic(uuid.NewRandom(), "The Walking Dead", "The Walking Dead 115", []string{}, []string{})

	actual := d.GetFrontPageView()
	expected := &domainservices.FrontPageView{
		Series: []domainservices.FrontPageViewSeries{domainservices.FrontPageViewSeries{Title: "The Walking Dead"}},
	}

	t.Log("The front page view should list the comic's series")
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %+v but was %+v", expected, actual)
	}
}

func TestGetFrontpageViewWithMultipleSeries(t *testing.T) {
	t.Log("When multiple comics from different series are added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domainservices.NewComicDomain(es, fs, vs)
	d.AddComic(uuid.NewRandom(), "The Walking Dead", "The Walking Dead 115", []string{}, []string{})
	d.AddComic(uuid.NewRandom(), "Warrior", "Warrior 1", []string{}, []string{})

	actual := d.GetFrontPageView()
	expected := &domainservices.FrontPageView{
		Series: []domainservices.FrontPageViewSeries{
			domainservices.FrontPageViewSeries{Title: "Warrior"},
			domainservices.FrontPageViewSeries{Title: "The Walking Dead"},
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
	d := domainservices.NewComicDomain(es, fs, vs)
	d.AddComic(uuid.NewRandom(), "The Walking Dead", "The Walking Dead 115", []string{}, []string{})
	d.AddComic(uuid.NewRandom(), "The Walking Dead", "The Walking Dead 114", []string{}, []string{})

	actual := d.GetFrontPageView()
	expected := &domainservices.FrontPageView{
		Series: []domainservices.FrontPageViewSeries{
			domainservices.FrontPageViewSeries{Title: "The Walking Dead"},
		},
	}

	t.Log("The front page view should list the series once")
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %+v but was %+v", expected, actual)
	}
}

func TestGetFrontpageViewWithNoSeries(t *testing.T) {
	t.Log("When no comics have been added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domainservices.NewComicDomain(es, fs, vs)

	actual := d.GetFrontPageView()
	expected := &domainservices.FrontPageView{
		Series: []domainservices.FrontPageViewSeries{},
	}

	t.Log("The view's series list should be empty")
	if reflect.DeepEqual(actual, expected) == false {
		t.Errorf("Expected %+v but was %+v", expected, actual)
	}
}
