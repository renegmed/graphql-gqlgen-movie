package models

type Movie struct {
	Id         int
	Title      string
	Year       int
	Genre      string
	Budget     float64
	Trailer    string
	DirectorId int
	Actors     []*Actor
}
