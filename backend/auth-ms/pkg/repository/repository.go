package repository

type Authentication interface {
}

type Repository struct {
	Authentication
}

func NewRepository() *Repository {
	return &Repository{}
}
