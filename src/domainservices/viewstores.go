package domainservices

type frontPageViewStore struct {
	vs ViewGetStorer
}

func newFrontPageViewStore(vs ViewGetStorer) frontPageViewStore {
	return frontPageViewStore{vs}
}

func (vs frontPageViewStore) Store(v *FrontPageView) {
	vs.vs.Store("frontpage", v) // TODO: Check for errors
}

func (vs frontPageViewStore) Get() FrontPageView {
	fp := &FrontPageView{Series: []FrontPageViewSeries{}}
	vs.vs.Get("frontpage", fp) // TODO: Check for errors

	return *fp
}
