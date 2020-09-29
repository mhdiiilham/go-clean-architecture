package book

type Writer interface {
	Create(e *Book) (int, error)
}

type repository interface {
	Writer
}

type Manager interface {
	repository
}
