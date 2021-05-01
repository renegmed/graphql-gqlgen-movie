init-project:
	go mod init movie-graphql-go

up:
	docker-compose up -d 
	
install:
	go get github.com/99designs/gqlgen 

init:
	gqlgen init 

generate:
	gqlgen generate 
	

run:
	POSTGRES_USER=admin POSTGRES_PASSWORD=pass123 POSTGRES_DB=moviedb DB_HOST=localhost go run -race main.go 