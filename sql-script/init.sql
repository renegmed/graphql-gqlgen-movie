DROP SCHEMA IF EXISTS moviedb;

CREATE SCHEMA IF NOT EXISTS moviedb;


CREATE TABLE IF NOT EXISTS directors (
  id                SERIAL PRIMARY KEY,
  full_name         VARCHAR(128),
  country           VARCHAR(128),
  CONSTRAINT uq_director UNIQUE (full_name)
);

CREATE TABLE IF NOT EXISTS movies (
  id            SERIAL             PRIMARY KEY,
  title         VARCHAR(256),
  release_year  INT,
  genre         VARCHAR(64),
  budget        FLOAT,
  trailer      VARCHAR(256),
  director_id   INT,
  CONSTRAINT uq_title UNIQUE (title),
  CONSTRAINT fk_movies_director FOREIGN KEY (director_id) REFERENCES directors(id)
);

CREATE TABLE IF NOT EXISTS actors (
  id            SERIAL PRIMARY KEY,
  full_name     VARCHAR(128),
  country       VARCHAR(128),
  male          BOOL,
  CONSTRAINT uq_actor UNIQUE (full_name)
);

CREATE TABLE movies_actors (
  movie_id          INT NOT NULL CONSTRAINT fk_movies_actors_movie REFERENCES movies(id),
  actor_id          INT NOT NULL CONSTRAINT fk_movies_actors_actor REFERENCES actors(id),
  CONSTRAINT pk_movies_actors PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE movies_rates (
  movie_id          INT NOT NULL CONSTRAINT fk_movies_rates_movie REFERENCES movies(id),
  email             VARCHAR(128) NOT NULL,
  score             INT NOT NULL,
  CONSTRAINT pk_movies_rates PRIMARY KEY(movie_id, email)
);

COMMIT;

INSERT INTO actors(full_name,country,male)
    VALUES ('Johnny Depp','USA',true);
INSERT INTO actors(full_name,country,male)
    VALUES ('Winona Ryder','USA',false);
INSERT INTO actors(full_name,country,male)
    VALUES ('Russell Crowe','Australia',true);
INSERT INTO actors(full_name,country,male)
    VALUES ('Joaquin Phoenix','USA',true);
INSERT INTO actors(full_name,country,male)
    VALUES ('Al Pacino','USA',true);
INSERT INTO actors(full_name,country,male)
    VALUES ('Robert de Niro','USA',true);

COMMIT;

INSERT INTO directors(full_name,country) VALUES ('Tim Burton', 'USA' );
INSERT INTO directors(full_name,country) VALUES ('James Cameron', 'Canada');
INSERT INTO directors(full_name,country) VALUES ('Steven Spielberg', 'USA');
INSERT INTO directors(full_name,country) VALUES ('Martin Scorsese', 'USA');
INSERT INTO directors(full_name,country) VALUES ('Alfred Hitchcock', 'UK');
INSERT INTO directors(full_name,country) VALUES ('Clint Eastwood', 'USA');
INSERT INTO directors(full_name,country) VALUES ('Ridley Scott', 'UK');
COMMIT;

INSERT INTO movies(title,release_year,genre,budget,trailer,director_id)
    VALUES ('Edward Scissorhands',1990,'SciFi',20,'https://www.youtube.com/watch?v=M94yyfWy-KI',1);

INSERT INTO movies(title,release_year,genre,budget,trailer,director_id)
    VALUES ('Gladiator',2000,'Drama',103,'https://www.youtube.com/watch?v=owK1qxDselE',7);
COMMIT;

INSERT INTO movies_actors(movie_id,actor_id)
    VALUES(1,1);
INSERT INTO movies_actors(movie_id,actor_id)
    VALUES(1,2);
INSERT INTO movies_actors(movie_id,actor_id)
    VALUES(2,3);
INSERT INTO movies_actors(movie_id,actor_id)
    VALUES(2,4);
COMMIT;


