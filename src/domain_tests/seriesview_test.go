package domain_tests

import (
	"code.google.com/p/go-uuid/uuid"
	"domain"
	"reflect"
	"testing"
	"time"
)

func TestGetSeriesViewWithASingleComic(t *testing.T) {
	t.Log("When a single comic is added with multiple pages")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	id := uuid.NewRandom()

	expectedTitle := "The Walking Dead 115"
	expectedSeriesTitle := "The Walking Dead"
	expectedBlurb := `ALL-OUT WAR BEGINS!
The biggest storyline in WALKING DEAD history – just in time to celebrate the 10th anniversary of the series! It's Rick versus Negan with a little help from everyone else!`
	expectedDate := time.Date(2013, time.October, 9, 0, 0, 0, 0, time.UTC)
	expectedWriters := []string{"Robert Kirkman"}
	expectedArtists := []string{"Charlie Adlard", "Cliff Rathburn"}

	d.AddComic(id, expectedSeriesTitle, expectedTitle, []string{"0.jpg", "1.jpg"}, []string{"source//path//0.jpg", "source//path//1.jpg"}, expectedWriters, expectedArtists, expectedDate, expectedBlurb)

	actual := d.GetSeriesView(expectedSeriesTitle)

	t.Log("The series view should list the series' title")
	if actual.Title != expectedSeriesTitle {
		t.Errorf("\tExpected %v but was %v", expectedSeriesTitle, actual.Title)
	}

	t.Log("The series view should only list the added comic")
	if len(actual.Books) != 1 {
		t.Errorf("\tExpected length of %v but was %v", 1, len(actual.Books))
	}

	t.Log("The series view should list the comic's title")
	if actual.Books[0].Title != expectedTitle {
		t.Errorf("\tExpected %v but was %v", expectedTitle, actual.Books[0].Title)
	}

	t.Log("The series view should list the comic's published date")
	if actual.Books[0].PublishedDate != expectedDate {
		t.Errorf("\tExpected %v but was %v", expectedDate, actual.Books[0].PublishedDate)
	}

	t.Log("The series view should list the comic's writers")
	if reflect.DeepEqual(actual.Books[0].WrittenBy, expectedWriters) == false {
		t.Errorf("\tExpected %v but was %v", expectedWriters, actual.Books[0].WrittenBy)
	}

	t.Log("The series view should list the comic's artists")
	if reflect.DeepEqual(actual.Books[0].ArtBy, expectedArtists) == false {
		t.Errorf("\tExpected %v but was %v", expectedArtists, actual.Books[0].ArtBy)
	}

	t.Log("The series view should list the comic's blurb")
	if actual.Books[0].Blurb != expectedBlurb {
		t.Errorf("\tExpected %v but was %v", expectedBlurb, actual.Books[0].Blurb)
	}

	t.Log("The series view should use the comic's first page for the book image")
	if actual.Books[0].ImageKey != id.String()+"/0.jpg" {
		t.Errorf("\tExpected %v but was %v", id.String()+"/0.jpg", actual.Books[0].ImageKey)
	}

	t.Log("The series view should list the comic's identifier")
	if actual.Books[0].Id != id.String() {
		t.Errorf("\tExpected %v but was %v", id.String(), actual.Books[0].Id)
	}
}

func TestGetSeriesViewWithMultipleComics(t *testing.T) {
	t.Log("When multiple comics of the same series are added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()

	d.AddComic(id2, "The Walking Dead", "The Walking Dead 114", []string{"1.jpg"}, []string{"source//path//1.jpg"}, []string{"Robert Kirkman"}, []string{"Charlie Adlard", "Cliff Rathburn"}, time.Date(2013, time.September, 11, 0, 0, 0, 0, time.UTC), "What would Jesus do?")
	d.AddComic(id1, "The Walking Dead", "The Walking Dead 115", []string{"0.jpg"}, []string{"source//path//0.jpg"}, []string{"Robert Kirkman"}, []string{"Charlie Adlard", "Cliff Rathburn"}, time.Date(2013, time.October, 9, 0, 0, 0, 0, time.UTC), "The Walking Dead 115 Blurb")

	actual := d.GetSeriesView("The Walking Dead")

	expected := domain.SeriesView{
		Title: "The Walking Dead",
		Books: []domain.SeriesViewBookView{
			domain.SeriesViewBookView{
				Id:            id1.String(),
				ImageKey:      id1.String() + "/0.jpg",
				Title:         "The Walking Dead 115",
				PublishedDate: time.Date(2013, time.October, 9, 0, 0, 0, 0, time.UTC),
				WrittenBy:     []string{"Robert Kirkman"},
				ArtBy:         []string{"Charlie Adlard", "Cliff Rathburn"},
				Blurb:         "The Walking Dead 115 Blurb",
			},
			domain.SeriesViewBookView{
				Id:            id2.String(),
				ImageKey:      id2.String() + "/1.jpg",
				Title:         "The Walking Dead 114",
				PublishedDate: time.Date(2013, time.September, 11, 0, 0, 0, 0, time.UTC),
				WrittenBy:     []string{"Robert Kirkman"},
				ArtBy:         []string{"Charlie Adlard", "Cliff Rathburn"},
				Blurb:         "What would Jesus do?",
			},
		},
	}

	t.Log("The series view should list the series' title")
	if actual.Title != expected.Title {
		t.Errorf("\tExpected %v but was %v", expected.Title, actual.Title)
	}

	t.Log("The series view should list all the added comics")
	if len(actual.Books) != len(expected.Books) {
		t.Errorf("\tExpected %v but was %v", len(expected.Books), len(actual.Books))
	}

	t.Log("The series view should list the last book added first")
	if reflect.DeepEqual(expected.Books[0], actual.Books[0]) == false {
		t.Errorf("\tExpected %#v but was %#v", expected.Books[0], actual.Books[0])
	}

	t.Log("The series view should list the first book added last")
	if reflect.DeepEqual(expected.Books[1], actual.Books[1]) == false {
		t.Errorf("\tExpected %#v but was %#v", expected.Books[1], actual.Books[1])
	}
}

func TestGetSeriesViewForDifferentSeries(t *testing.T) {
	t.Log("When multiple comics of different series are added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)
	id1 := uuid.NewRandom()
	id2 := uuid.NewRandom()

	d.AddComic(id1, "Prophet", "Prophet 31", []string{"0.jpg"}, []string{"source//path//0.jpg"}, []string{"Brandon Graham", "Simon Roy", "Giannis Milonogiannis"}, []string{"Giannis Milonogiannis"}, time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC), "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance.")
	d.AddComic(id2, "The Walking Dead", "The Walking Dead 115", []string{"0.jpg"}, []string{"source//path//0.jpg"}, []string{"Robert Kirkman"}, []string{"Charlie Adlard", "Cliff Rathburn"}, time.Date(2013, time.October, 9, 0, 0, 0, 0, time.UTC), "The Walking Dead 115 Blurb")

	actualProphet := d.GetSeriesView("Prophet")
	actualWalkingDead := d.GetSeriesView("The Walking Dead")

	expectedProphet := &domain.SeriesView{
		Title: "Prophet",
		Books: []domain.SeriesViewBookView{
			domain.SeriesViewBookView{
				Id:            id1.String(),
				ImageKey:      id1.String() + "/0.jpg",
				Title:         "Prophet 31",
				PublishedDate: time.Date(2012, time.November, 28, 0, 0, 0, 0, time.UTC),
				WrittenBy:     []string{"Brandon Graham", "Simon Roy", "Giannis Milonogiannis"},
				ArtBy:         []string{"Giannis Milonogiannis"},
				Blurb:         "Old Man Prophet goes to meet with a lost matriarchal tribe of humanity to try to form an alliance.",
			},
		},
	}

	expectedWalkingDead := &domain.SeriesView{
		Title: "The Walking Dead",
		Books: []domain.SeriesViewBookView{
			domain.SeriesViewBookView{
				Id:            id2.String(),
				ImageKey:      id2.String() + "/0.jpg",
				Title:         "The Walking Dead 115",
				PublishedDate: time.Date(2013, time.October, 9, 0, 0, 0, 0, time.UTC),
				WrittenBy:     []string{"Robert Kirkman"},
				ArtBy:         []string{"Charlie Adlard", "Cliff Rathburn"},
				Blurb:         "The Walking Dead 115 Blurb",
			},
		},
	}

	t.Log("Each series should only contain the comics for that series")
	if reflect.DeepEqual(expectedProphet, actualProphet) == false {
		t.Errorf("\tExpected %#v but was %#v", expectedProphet, actualProphet)
	}
	if reflect.DeepEqual(expectedWalkingDead, actualWalkingDead) == false {
		t.Errorf("\tExpected %#v but was %#v", expectedWalkingDead, actualWalkingDead)
	}
}

func TestGetSeriesViewWithAnUnknownSeriesTitle(t *testing.T) {
	t.Log("When no comics have been added")

	es := NewFakeEventStorer()
	fs := NewFakeFileStore()
	vs := NewFakeViewStore()
	d := domain.NewComicDomain(es, fs, vs)

	actual := d.GetSeriesView("The Walking Dead")

	if actual != nil {
		t.Errorf("Expected %v but was %#v", nil, actual)
	}
}
