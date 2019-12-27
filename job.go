package woker_pool


type Job interface {
	Do() error
	ID() string
}

