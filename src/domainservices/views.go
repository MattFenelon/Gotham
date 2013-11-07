package domainservices

type FrontPageView struct {
	Series []FrontPageViewSeries
}

type FrontPageViewSeries struct {
	Title string
}
