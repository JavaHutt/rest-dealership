package repository

type Repository struct {
	Vehicle
}

func NewRepository() *Repository {
	return &Repository{}
}

type Vehicle interface{}
