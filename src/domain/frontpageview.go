package domain

type FrontPageView struct {
	Series []FrontPageViewSeries
}

type FrontPageViewSeries struct {
	Title          string
	ImageKey       string
	PromotedBookId string
}
