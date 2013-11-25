package domain

import (
	"code.google.com/p/go-uuid/uuid"
	"domain/model"
)

type ComicView struct {
	Pages []string
}

type comicViewStore struct {
	vs ViewGetStorer
}

func newComicViewStore(vs ViewGetStorer) *comicViewStore {
	return &comicViewStore{vs}
}

func (c *comicViewStore) Store(event *model.ComicAdded) {
	view := &ComicView{Pages: event.Pages}
	c.vs.Store("ComicView:"+event.Id.String(), view)
}

func (c *comicViewStore) Get(id uuid.UUID) *ComicView {
	var view ComicView
	if err := c.vs.Get("ComicView:"+id.String(), &view); err != nil {
		return nil
	}
	return &view
}
