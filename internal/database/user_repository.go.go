package database

import (
	"database/sql"

	"go-concurrency-sample/internal/types"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Insert(user *types.User) error {
	_, err := r.db.Exec(`
		INSERT INTO users (
			cpf,
			name,
			email,
			birthdate
		)
		VALUES (?, ?, ?, ?)
	`,
		user.CPF,
		user.Name,
		user.Email,
		user.Birthdate,
	)

	return err
}
