package ticket

import "project-control-system/internal/storage"

type TicketRepo interface {
}

type Repository struct {
	storage *storage.Storage
}

var _ TicketRepo = (*Repository)(nil)

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage}
}
