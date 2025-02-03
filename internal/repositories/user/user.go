package user

import (
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"log"
	"project-control-system/internal/entities"
	"project-control-system/internal/storage"
)

type UserRepo interface {
	ExistsByEmail(email string) (bool, error)
	ExistsByUsername(username string) (bool, error)
	Create(user entities.User) error
	GetByUsername(username string) (entities.User, error)
}

type Repository struct {
	storage *storage.Storage
}

var _ UserRepo = (*Repository)(nil)

func NewRepository(storage *storage.Storage) *Repository {
	return &Repository{storage}
}

func (r *Repository) ExistsByEmail(email string) (bool, error) {
	query, _, err := r.storage.Goqu.From("users").
		Select(goqu.COUNT("*")).
		Where(goqu.C("email").Eq(email)).
		ToSQL()

	if err != nil {
		return false, fmt.Errorf("[userRepo][ExistsByEmail]: failed to build SQL query: %w", err)
	}

	// Выполняем запрос
	var count int
	err = r.storage.Goqu.QueryRow(query).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("[userRepo][ExistsByEmail]: failed to execute query: %w", err)
	}

	return count > 0, nil
}

func (r *Repository) ExistsByUsername(username string) (bool, error) {
	query, _, err := r.storage.Goqu.From("users").
		Select(goqu.COUNT("*")).
		Where(goqu.C("username").Eq(username)).
		ToSQL()

	if err != nil {
		return false, fmt.Errorf("[userRepo][ExistsByUsername]: failed to build SQL query: %w", err)
	}

	// Выполняем запрос
	var count int
	err = r.storage.Goqu.QueryRow(query).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("[userRepo][ExistsByUsername]: failed to execute query: %w", err)
	}

	return count > 0, nil
}

func (r *Repository) Create(user entities.User) error {
	insertQuery := r.storage.Goqu.Insert("users").
		Cols("username", "email", "password").
		Vals(goqu.Vals{user.Username, user.Email, user.Password})

	_, err := insertQuery.Executor().Exec()
	if err != nil {
		return fmt.Errorf("[userRepo][create]: failed to insert user: %w", err)
	}
	log.Println("[userRepo][create]: user created")
	return nil
}

func (r *Repository) GetByUsername(username string) (entities.User, error) {
	var user entities.User
	err := r.storage.Goqu.From("users").
		Select("id", "username", "email", "password").ScanStructs(&user)

	if err != nil {
		return entities.User{}, fmt.Errorf("[userRepo][GetByUsername]: failed to get user: %w", err)
	}

	return user, nil
}
