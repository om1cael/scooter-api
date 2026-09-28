package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/om1cael/scooter-api/internal/models"
)

var errNotFound = errors.New("user not found")

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id string) (*models.User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`

	return r.db.QueryRowContext(
		ctx, query, user.Name, user.Email, user.Password,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *SQLUserRepository) GetByID(ctx context.Context, id string) (*models.User, error) {
	query := `
		SELECT id, name, email 
		FROM users
		WHERE id = $1
	`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
