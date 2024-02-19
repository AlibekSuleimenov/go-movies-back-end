package repository

import "github.com/alibeksuleimenov/go-movies-back-end/internal/models"

type DatabaseRepo interface {
	AllMovies() ([]*models.Movie, error)
}
