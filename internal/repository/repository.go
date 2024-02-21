package repository

import (
	"database/sql"
	"github.com/alibeksuleimenov/go-movies-back-end/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}
