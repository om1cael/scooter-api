package repositories

import (
	"context"
	"database/sql"

	"github.com/om1cael/scooter-api/internal/models"
)

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
