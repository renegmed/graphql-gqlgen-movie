package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"movie-graphql-go/graph/generated"
	"movie-graphql-go/graph/model"
	"movie-graphql-go/models"
	"strconv"
)

func (r *mutationResolver) AddMovie(ctx context.Context, request *model.MovieRequest) (*model.Movie, error) {
	movie := models.Movie{
		Title:  request.Title,
		Budget: request.Budget,
	}

	if request.DirectorID != "" {
		directorId, err := strconv.Atoi(request.DirectorID)
		if err != nil {
			log.Println("...Error on add movie:", err)
			return nil, err
		}
		movie.DirectorId = directorId
	}

	if request.Year != nil {
		movie.Year = *request.Year
	}
	if request.Genre != nil {
		movie.Genre = *request.Genre
	}
	if request.Trailer != nil {
		movie.Trailer = string(*request.Trailer)
	}

	m, err := r.Repo.InsertMovie(movie)
	if err != nil {
		return nil, err
	}
	return r.modelsToGraphMovie(m)

}

func (r *mutationResolver) modelsToGraphMovie(movie *models.Movie) (*model.Movie, error) {
	return &model.Movie{
		ID:      strconv.Itoa(movie.Id),
		Title:   movie.Title,
		Year:    &movie.Year,
		Genre:   &movie.Genre,
		Budget:  movie.Budget,
		Trailer: &movie.Trailer,
	}, nil
}

func (r *mutationResolver) AddDirector(ctx context.Context, request *model.DirectorRequest) (*model.Director, error) {
	director := models.Director{
		FullName: request.FullName,
	}
	if request.Country != nil {
		director.Country = *request.Country
	}
	d, err := r.Repo.InsertDirector(director)
	if err != nil {
		return nil, err
	}
	return &model.Director{
		ID:       strconv.Itoa(d.Id),
		FullName: d.FullName,
		Country:  &d.Country,
	}, nil

	//model.Director: *d}, nil
}

func (r *mutationResolver) DeleteDirector(ctx context.Context, directorID string) ([]*model.Director, error) {
	directorId, err := strconv.Atoi(directorID)
	if err != nil {
		log.Println("...Error on add movie:", err)
		return nil, err
	}
	if err := r.Repo.DeleteDirector(directorId); err != nil {
		return nil, err
	}
	directors, err := r.Repo.ListDirectors()
	if err != nil {
		return nil, err
	}
	output := make([]*model.Director, len(directors))
	for index, modelsDirector := range directors {
		output[index] = &model.Director{
			//Director: directors[index]
			ID:       strconv.Itoa(modelsDirector.Id),
			FullName: modelsDirector.FullName,
			Country:  &modelsDirector.Country,
		}
	}
	return output, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
//func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }
