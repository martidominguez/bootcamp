USE movies_db;

-- todos los registros de la tabla movies
SELECT id, title, rating, awards, release_date, length, genre_id FROM movies;

-- nombre, apellido y rating de los actores
SELECT first_name, last_name, rating FROM actors;

-- titulo de todas las series y usar alias para que tanto el nombre de la tabla como el campo estén en español
SELECT title AS titulo FROM series AS s;

-- nombre y apellido de actores cuyo rating sea mayor a 7.5
SELECT first_name, last_name FROM actors WHERE rating > 7.5;

-- título de las películas, el rating y los premios de las películas con un rating mayor a 7.5 y con más de dos premios
SELECT title, rating, awards FROM movies WHERE rating > 7.5 AND awards > 2;

-- título de las películas y el rating ordenadas por rating en forma ascendente
SELECT title, rating FROM movies ORDER BY rating ASC;

-- títulos de las primeras tres películas en la base de datos
SELECT title FROM movies LIMIT 3;

-- top 5 de las películas con mayor rating
SELECT id, title, rating, awards, release_date, length, genre_id FROM movies ORDER BY rating DESC LIMIT 5;

-- 10 primeros actores
SELECT id, first_name, last_name, rating, favorite_movie_id FROM actors LIMIT 10;

-- título y rating de todas las películas cuyo título sea de Toy Story
SELECT title, rating FROM movies WHERE title LIKE '%Toy Story%';

-- todos los actores cuyos nombres empiezan con Sam
SELECT id, first_name, last_name, rating, favorite_movie_id FROM actors WHERE first_name LIKE 'Sam%';

-- título de las películas que salieron entre el 2004 y 2008
SELECT title FROM movies WHERE YEAR(release_date) BETWEEN 2004 AND 2008;

-- título de las películas con el rating mayor a 3, con más de 1 premio y con fecha de lanzamiento entre el año 1988 al 2009. ordenar los resultados por rating
SELECT title 
FROM movies 
WHERE rating > 3 AND awards > 1 AND YEAR(release_date) BETWEEN 1998 AND 2009 
ORDER BY rating;


