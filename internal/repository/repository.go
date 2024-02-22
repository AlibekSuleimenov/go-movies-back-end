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
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error)
	OneMovie(id int) (*models.Movie, error)
	AllGenres() ([]*models.Genre, error)
}
