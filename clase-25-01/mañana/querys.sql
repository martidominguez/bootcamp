USE movies_db;

-- título y nombre del género de todas las series
SELECT s.title AS 'serie_title', g.name AS 'genre'
FROM series s JOIN genres g ON s.genre_id = g.id;

-- título de los episodios, nombre y apellido de los actores que trabajan en cada uno de ellos
SELECT e.title AS 'episode_title', a.first_name AS 'actor_first_name', a.last_name AS 'actor_last_name'
FROM episodes e 
	JOIN actor_episode ae ON e.id = ae.episode_id
    JOIN actors a ON ae.actor_id = a.id;

-- título de todas las series y el total de temporadas que tiene cada una de ellas
SELECT s.title AS 'serie_title', COUNT(*) AS 'amount_of_seasons'
FROM series s
	JOIN seasons sa ON s.id = sa.serie_id
GROUP BY s.title;

-- nombre de todos los géneros y cantidad total de películas por cada uno, siempre que sea mayor o igual a 3
SELECT g.name AS 'genre', COUNT(*) as 'total_movies'
FROM genres g
	JOIN movies m ON g.id = m.genre_id
GROUP BY g.name
HAVING total_movies >= 3;

-- nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan
SELECT DISTINCT first_name, last_name
FROM movies m
	JOIN actor_movie am ON m.id = am.movie_id
    JOIN actors a ON am.actor_id = a.id
WHERE m.title LIKE '%La Guerra de las galaxias%'

