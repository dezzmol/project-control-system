package milestone

import "project-control-system/internal/storage"

type MilestoneRepo interface {
}

type Repository struct {
	storage *storage.Storage
}

var _ MilestoneRepo = (*Repository)(nil)

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage}
}
