package book

type manager struct {
	repo repository
}

func NewManager(r repository) *manager {
	return &manager{
		repo: r,
	}
}

func (s *manager) Create(e *Book) (int, error) {
	return s.repo.Create(e)
}
