package postgresdb

import (
	"database/sql"
	"fmt"
	"log"
	"movie-graphql-go/models"
	"os"

	_ "github.com/lib/pq"
)

const (
	PORT               = 5432
	maxTties           = 5
	pgListDirectors    = "SELECT * FROM directors"
	pgGetDirector      = "SELECT * FROM directors WHERE id=$1"
	pgListActors       = "SELECT * FROM actors"
	pgListMovies       = "SELECT * FROM movies"
	pgGetMovie         = "SELECT * FROM movies WHERE id=$1"
	pgGetActorsInMovie = "SELECT a.id,a.full_name,a.country,a.male FROM movies_actors ma, actors a WHERE ma.actor_id=a.id AND ma.movie_id=$1 LIMIT $2"
	pgInsertDirector   = "INSERT INTO directors(full_name,country) VALUES($1,$2) RETURNING id"
	pgDeleteDirector   = "DELETE FROM directors WHERE id=$1"
	pgInsertMovie      = "INSERT INTO movies(title,release_year,budget,genre,trailer,director_id) VALUES($1,$2,$3,$4,$5,$6) RETURNING id"
)

var ErrNoMatch = fmt.Errorf("no matching record")

// This is a repository.Respository
type PostgresDb struct {
	Conn *sql.DB
}

func NewPostgresDb(username, password, database string) (*PostgresDb, error) {
	host := os.Getenv("DB_HOST")
	db := PostgresDb{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, PORT, username, password, database)

	log.Println("dsn:", dsn)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return &db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return &db, err
	}
	log.Println("Database connection established")
	return &db, nil
}

func (db *PostgresDb) ListDirectors() ([]models.Director, error) {
	rows, err := db.Conn.Query(pgListDirectors)
	if err != nil {
		return nil, err
	}
	directors := make([]models.Director, 0)
	for rows.Next() {
		var director models.Director
		if err := rows.Scan(&director.Id, &director.FullName, &director.Country); err != nil {
			return nil, err
		}
		directors = append(directors, director)
	}

	return directors, nil
}

func (db *PostgresDb) ListActors() ([]models.Actor, error) {
	rows, err := db.Conn.Query(pgListActors)
	if err != nil {
		return nil, err
	}
	actors := make([]models.Actor, 0)
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.Id, &actor.FullName, &actor.Country, &actor.Male); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}

func (db *PostgresDb) ListMovies() ([]models.Movie, error) {
	rows, err := db.Conn.Query(pgListMovies)
	if err != nil {
		return nil, err
	}
	movies := make([]models.Movie, 0)
	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Genre, &movie.Budget, &movie.Trailer, &movie.DirectorId); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (db *PostgresDb) GetDirector(directorId int) (*models.Director, error) {
	stmt, err := db.Conn.Prepare(pgGetDirector)
	if err != nil {
		return nil, err
	}
	rows := stmt.QueryRow(directorId)
	director := models.Director{}
	if err := rows.Scan(&director.Id, &director.FullName, &director.Country); err != nil {
		return nil, err
	}
	return &director, nil
}

func (db *PostgresDb) GetMovie(movieId int) (*models.Movie, error) {
	stmt, err := db.Conn.Prepare(pgGetMovie)
	if err != nil {
		return nil, err
	}
	movie := models.Movie{}
	if err := stmt.QueryRow(movieId).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Genre, &movie.Budget, &movie.Trailer, &movie.DirectorId); err != nil {
		return nil, err
	}
	return &movie, nil
}

func (db *PostgresDb) GetActorsForMovieWithId(movieId int, limit int) ([]models.Actor, error) {
	stmt, err := db.Conn.Prepare(pgGetActorsInMovie)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(movieId, limit)
	if err != nil {
		return nil, err
	}
	actors := make([]models.Actor, 0)
	for rows.Next() {
		var actor models.Actor
		if err := rows.Scan(&actor.Id, &actor.FullName, &actor.Country, &actor.Male); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}
	return actors, nil
}

func (db *PostgresDb) InsertDirector(director models.Director) (*models.Director, error) {
	stmt, err := db.Conn.Prepare(pgInsertDirector)
	if err != nil {
		return nil, err
	}
	var directId int
	if err := stmt.QueryRow(director.FullName, director.Country).Scan(&directId); err != nil {
		return nil, err
	}
	director.Id = directId
	return &director, nil
}

func (db *PostgresDb) DeleteDirector(directorId int) error {
	stmt, err := db.Conn.Prepare(pgDeleteDirector)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(directorId)
	if err != nil {
		return err
	}
	changed, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if changed == 0 {
		return fmt.Errorf("There's not director with given id")
		// return &errors2.QueryError{
		// 	Message:   "There's not director with given id",
		// 	Path:      []interface{}{"model", "postgresql"},
		// 	Locations: []errors2.Location{{158, 21}},
		// }
	}
	return nil
}

func (db *PostgresDb) InsertMovie(movie models.Movie) (*models.Movie, error) {
	tx, _ := db.Conn.Begin()
	stmt, err := tx.Prepare(pgInsertMovie)
	if err != nil {
		return nil, err
	}
	var movieId int
	if err := stmt.QueryRow(movie.Title, movie.Year, movie.Budget, movie.Genre, movie.Trailer, movie.DirectorId).Scan(&movieId); err != nil {
		tx.Rollback()
		return nil, err
	}
	movie.Id = movieId
	tx.Commit()
	return &movie, nil
}
