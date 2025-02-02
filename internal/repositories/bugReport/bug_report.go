package bugReport

import "project-control-system/internal/storage"

type BugReportRepo interface {
}

type Repository struct {
	storage *storage.Storage
}

var _ BugReportRepo = (*Repository)(nil)

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage}
}
