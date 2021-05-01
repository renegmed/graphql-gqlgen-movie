package repository

import (
	"movie-graphql-go/models"
)

type AppRepo struct {
	RepoDb Repository
}

func NewAppRepository(repoDb Repository) *AppRepo {
	return &AppRepo{repoDb}
}

func (r *AppRepo) ListDirectors() ([]models.Director, error) {
	return r.RepoDb.ListDirectors()
}

func (r *AppRepo) ListActors() ([]models.Actor, error) {
	return r.RepoDb.ListActors()
}

func (r *AppRepo) ListMovies() ([]models.Movie, error) {
	return r.RepoDb.ListMovies()
}

func (r *AppRepo) GetDirector(directorId int) (*models.Director, error) {
	return r.RepoDb.GetDirector(directorId)
}

func (r *AppRepo) GetMovie(movieId int) (*models.Movie, error) {
	return r.RepoDb.GetMovie(movieId)
}

func (r *AppRepo) GetActorsForMovieWithId(movieId int, limit int) ([]models.Actor, error) {
	return r.RepoDb.GetActorsForMovieWithId(movieId, limit)
}

func (r *AppRepo) InsertDirector(director models.Director) (*models.Director, error) {
	return r.RepoDb.InsertDirector(director)
}

func (r *AppRepo) DeleteDirector(directorId int) error {
	return r.RepoDb.DeleteDirector(directorId)
}

func (r *AppRepo) InsertMovie(movie models.Movie) (*models.Movie, error) {
	return r.RepoDb.InsertMovie(movie)
}
