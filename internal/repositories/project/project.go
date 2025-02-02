package project

import (
	"project-control-system/internal/entities"
	"project-control-system/internal/storage"
)

type ProjectRepo interface {
	GetProjectById(id string) (entities.Project, error)
}

type Repository struct {
	storage *storage.Storage
}

var _ ProjectRepo = (*Repository)(nil)

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage}
}

func (repo *Repository) GetProjectById(id string) (entities.Project, error) {
	return entities.Project{}, nil
}
