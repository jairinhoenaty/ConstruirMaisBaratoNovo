package pageview

type PageViewRepository interface {
	Increment(path string) (*PageView, error)
	FindAll() ([]*PageView, error)
}
