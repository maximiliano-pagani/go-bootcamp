-- Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase.
-- Importar el archivo movies_db.sql desde PHPMyAdmin o MySQL Workbench y resolver las siguientes consultas:

-- Mostrar el título y el nombre del género de todas las series.

SELECT series.title as serie, genres.name as genre
	FROM series
	INNER JOIN genres ON series.genre_id = genres.id;

-- Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.

SELECT episodes.title as episodio, CONCAT(actors.first_name, " ", actors.last_name) as actor 
	FROM episodes
    INNER JOIN actor_episode ON actor_episode.episode_id = episodes.id
    INNER JOIN actors ON actor_episode.actor_id = actors.id;

-- Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.

SELECT series.title AS titulo, COUNT(*) as temporadas
	FROM series
    INNER JOIN seasons ON series.id = seasons.serie_id
    GROUP BY series.title;

-- Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.

SELECT genres.name AS genero, COUNT(movies.id) AS total_peliculas
	FROM genres
    LEFT JOIN movies ON genres.id = movies.genre_id
    GROUP BY genres.id
    HAVING total_peliculas >= 3;

-- Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.

SELECT DISTINCT CONCAT(actors.first_name, " ", actors.last_name) as actor
	FROM movies
    INNER JOIN actor_movie ON actor_movie.movie_id = movies.id
    INNER JOIN actors on actor_movie.actor_id = actors.id
    WHERE movies.title LIKE "%Guerra de las galaxias%";
    