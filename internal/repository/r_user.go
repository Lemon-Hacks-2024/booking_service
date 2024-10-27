package repository

import (
	"booking_service/internal/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	postgres *sqlx.DB
}

func NewUserRepository(postgres *sqlx.DB) *UserRepository {
	return &UserRepository{
		postgres: postgres,
	}
}

func (r *UserRepository) CreateUser(user entity.User) (int, error) {
	var id int

	query := `INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id;`

	row := r.postgres.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetUserByID(userID int) (entity.User, error) {
	var user entity.User

	query := `SELECT id, first_name, last_name, email, password FROM users WHERE id = $1 AND deleted_at IS NULL;`

	row := r.postgres.QueryRow(query, userID)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByLogIN(email, password string) (entity.User, error) {
	var user entity.User

	query := `SELECT id, first_name, last_name, email, password FROM users WHERE email = $1 AND password = $2 AND deleted_at IS NULL;`

	row := r.postgres.QueryRow(query, email, password)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
