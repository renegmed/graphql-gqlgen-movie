package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"movie-graphql-go/graph/generated"
	"movie-graphql-go/graph/model"
	"movie-graphql-go/models"
	"strconv"
)

func (r *queryResolver) ListDirectors(ctx context.Context) ([]*model.Director, error) {
	directors, err := r.Repo.ListDirectors()
	if err != nil {
		return nil, err
	}
	output := []*model.Director{}
	for _, director := range directors {
		//log.Println("... director country:", director.Country)
		output = append(output, &model.Director{
			ID:       strconv.Itoa(director.Id),
			FullName: director.FullName,
			Country:  &director.Country,
		})
	}
	log.Println("...List directors output", *output[4].Country)
	return output, nil
}

func (r *queryResolver) ListActors(ctx context.Context) ([]*model.Actor, error) {
	actors, err := r.Repo.ListActors()
	if err != nil {
		return nil, err
	}

	output := []*model.Actor{}
	for _, actor := range actors {

		output = append(output, &model.Actor{
			ID:       strconv.Itoa(actor.Id),
			FullName: actor.FullName,
			Country:  &actor.Country,
			Gender: func() *string {
				if actor.Male {
					gender := "male"
					return &gender
				}
				gender := "female"
				return &gender
			}(),
		})
	}
	return output, nil
}

func (r *queryResolver) ListMovies(ctx context.Context) ([]*model.Movie, error) {
	movies, err := r.Repo.ListMovies()
	if err != nil {
		return nil, err
	}

	output := []*model.Movie{}
	for _, movie := range movies {
		graphMovie, err := r.toGraphMovie(&movie)
		if err != nil {
			log.Println("...Error on convert to graph movie:", err)
		}
		output = append(output, graphMovie)
	}

	return output, nil
}

func (r *queryResolver) getDirector(id int) (*models.Director, error) {
	return r.Repo.GetDirector(id)
}

func (r *queryResolver) listMovieActors(movieId int, limit int) ([]*model.Actor, error) {
	movieActors, err := r.Repo.GetActorsForMovieWithId(movieId, limit)
	if err != nil {
		return nil, err
	}

	actors := []*model.Actor{}
	for _, actor := range movieActors {
		actors = append(actors, &model.Actor{
			ID:       strconv.Itoa(actor.Id),
			FullName: actor.FullName,
			Country:  &actor.Country,
			Gender: func() *string {
				if actor.Male {
					gender := "male"
					return &gender
				}
				gender := "female"
				return &gender
			}(),
		})
	}
	return actors, nil
}
func (r *queryResolver) GetMovie(ctx context.Context, movieID string) (*model.Movie, error) {
	movieId, err := strconv.Atoi(movieID)
	if err != nil {
		log.Println("...Error on GetMovie():", err)
		return nil, err
	}
	movie, err := r.Repo.GetMovie(movieId)
	if err != nil {
		return nil, err
	}

	return r.toGraphMovie(movie)
}

func (r *queryResolver) toGraphMovie(movie *models.Movie) (*model.Movie, error) {
	dir, err := r.getDirector(movie.DirectorId)
	if err != nil {
		log.Println("...Error on get director:", err)
	}

	director := model.Director{
		ID:       strconv.Itoa(dir.Id),
		FullName: dir.FullName,
		Country:  &dir.Country,
	}

	actors, err := r.listMovieActors(movie.Id, 100)
	if err != nil {
		log.Println("...Error on create list movie actors:", err)
		return nil, err
	}

	return &model.Movie{
		ID:       strconv.Itoa(movie.Id),
		Title:    movie.Title,
		Year:     &movie.Year,
		Genre:    &movie.Genre,
		Budget:   movie.Budget,
		Trailer:  &movie.Trailer,
		Director: &director,
		Actors:   actors,
	}, nil
}
func (r *subscriptionResolver) ListenDirectorMovies(ctx context.Context, directorID string) (<-chan *model.Movie, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }
