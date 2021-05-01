package repository

import (
	"movie-graphql-go/models"
)

type Repository interface {
	// GetDB(config *config.DatabaseConfig) error
	ListDirectors() ([]models.Director, error)
	ListActors() ([]models.Actor, error)
	ListMovies() ([]models.Movie, error)
	GetDirector(directorId int) (*models.Director, error)
	GetMovie(movieId int) (*models.Movie, error)
	GetActorsForMovieWithId(movieId int, limit int) ([]models.Actor, error)
	InsertDirector(director models.Director) (*models.Director, error)
	DeleteDirector(directorId int) error
	InsertMovie(movie models.Movie) (*models.Movie, error)
}
