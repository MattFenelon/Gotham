package domain

import (
	"domain/model"
)

type frontPageViewStore struct {
	vs ViewGetStorer
}

func newFrontPageViewStore(vs ViewGetStorer) *frontPageViewStore {
	return &frontPageViewStore{vs}
}

func (viewStore *frontPageViewStore) Store(event *model.ComicAdded) {
	fp := viewStore.Get()

	for i, s := range fp.Series {
		if s.Title == event.SeriesTitle.String() {
			fp.Series[i].ImageKey = event.Id.String() + "/" + event.Pages[0]
			fp.Series[i].PromotedBookId = event.Id.String()
			viewStore.doStore(&fp)
			return
		}
	}

	newseries := []FrontPageViewSeries{
		FrontPageViewSeries{
			Title:          event.SeriesTitle.String(),
			ImageKey:       event.Id.String() + "/" + event.Pages[0],
			PromotedBookId: event.Id.String(),
		},
	}

	fp.Series = append(newseries, fp.Series...)
	viewStore.doStore(&fp)
}

func (s *frontPageViewStore) doStore(fp *FrontPageView) {
	s.vs.Store("frontpage", fp) // TODO: Check for errors
}

func (s *frontPageViewStore) Get() FrontPageView {
	fp := FrontPageView{Series: []FrontPageViewSeries{}}
	s.vs.Get("frontpage", &fp) // TODO: Check for errors
	return fp
}
