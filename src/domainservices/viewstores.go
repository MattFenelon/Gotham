package domainservices

type frontPageViewStore struct {
	vs ViewGetStorer
}

func newFrontPageViewStore(vs ViewGetStorer) frontPageViewStore {
	return frontPageViewStore{vs}
}

func (vs frontPageViewStore) Store(v *FrontPageView) {
	vs.vs.Store("frontpage", v)
}

func (vs frontPageViewStore) Get() FrontPageView {
	fp, ok := vs.vs.Get("frontpage").(*FrontPageView)

	if ok == false || fp == nil {
		fp = &FrontPageView{Series: []FrontPageViewSeries{}}
	}

	return *fp
}
