USE movies_db;

-- agregar una película a la tabla movies
INSERT INTO movies (title, rating, awards, release_date, length, genre_id)
VALUES ('Superbad', 8, 2, '2008-10-02 00:00:00', 172, 1);

-- agregar un género a la tabla genres
INSERT INTO genres (name, ranking, active)
VALUES ('Romance', 13, 1);

-- asociar a la película del punto 1 el género creado en el punto 2
UPDATE movies
SET genre_id = 13
WHERE id = 22;

-- modificar la tabla actors para que un actor tenga como favorita la película agregada en el punto 1
UPDATE actors
SET favorite_movie_id = 22
WHERE id = 1;

-- crear una tabla temporal copia de la tabla movies
CREATE TEMPORARY TABLE copia_movies AS
SELECT *
FROM movies;

-- eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards
DELETE FROM copia_movies
WHERE awards < 5;

-- obtener la lista de todos los géneros que tengan al menos una película
SELECT name 
FROM genres g
	JOIN movies m ON g.id = m.genre_id
GROUP BY name
HAVING COUNT(*) >= 1;

-- obtener la lista de actores cuya película favorita haya ganado más de 3 awards
SELECT first_name, last_name 
FROM actors a
	JOIN movies m ON a.favorite_movie_id = m.id
WHERE m.awards > 3;

-- crear un índice sobre el nombre en la tabla movies
CREATE INDEX movie_name_index ON movies (title);

-- chequear que el índice fue creado correctamente
SHOW INDEX FROM movies;









