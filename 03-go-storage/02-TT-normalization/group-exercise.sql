-- Tomando la base de datos movies_db.sql, se solicita:

USE movies_db;

-- Agregar una película a la tabla movies.

INSERT INTO movies (title, rating, awards, release_date, length, genre_id) 
	VALUES ("El padrino", 10.0, 14, NOW(), 270, 3);
    
-- Agregar un género a la tabla genres.

INSERT INTO genres (name, ranking, active) 
	VALUES ("Policial", 13, 1);
    
-- Asociar a la película del punto 1. genre el género creado en el punto 2.

START TRANSACTION;
SELECT @G:=genres.id FROM genres WHERE genres.name = "Policial";
SELECT @M:=movies.id FROM movies WHERE movies.title = "El Padrino";
UPDATE movies SET movies.genre_id = @G WHERE movies.id = @M;
COMMIT;

-- Modificar la tabla actors para que al menos un actor tenga como favorita la película agregada en el punto 1.

START TRANSACTION;
SELECT @A:=actors.id FROM actors LIMIT 1;
SELECT @M:=movies.id FROM movies WHERE movies.title = "El Padrino";
UPDATE actors SET actors.favorite_movie_id = @M WHERE id = @A;
COMMIT;

-- Crear una tabla temporal copia de la tabla movies.

CREATE TEMPORARY TABLE movies_copy
	SELECT * FROM movies;

SELECT * FROM movies_copy;

-- Eliminar de esa tabla temporal todas las películas que hayan ganado menos de 5 awards.

DELETE FROM movies_copy WHERE movies_copy.awards < 5;

SELECT * FROM movies_copy;

-- Obtener la lista de todos los géneros que tengan al menos una película.

SELECT genres.name
	FROM genres
    INNER JOIN movies ON genres.id = movies.genre_id
    GROUP BY genres.id;

-- Obtener la lista de actores cuya película favorita haya ganado más de 3 awards.

SELECT CONCAT(actors.first_name, " ", actors.last_name), movies.title
	FROM actors
    INNER JOIN movies ON actors.favorite_movie_id = movies.id
    WHERE movies.awards > 3;

-- Crear un índice sobre el nombre en la tabla movies.

CREATE INDEX IX_title ON movies (title);

-- Chequee que el índice fue creado correctamente.

SHOW INDEXES FROM movies;

-- En la base de datos movies ¿Existiría una mejora notable al crear índices? Analizar y justificar la respuesta.

-- Depende el índice y el tipo de consultas que se realizan habitualmente sobre esa tabla. En el caso por ejemplo del índice en
-- la columna title, es probable que la query de búsqueda por título de película se haga de forma frecuente, por lo que sería lógico
-- crearlo para mejorar la performance de dicha query.

-- ¿En qué otra tabla crearía un índice y por qué? Justificar la respuesta

-- Podría ser por ejemplo también lo referido a título o nombre de las series, actores y géneros, frecuente en búsquedas.
-- También quizás en la columna rating de movies, ya que son habituales las búsquedas de películas con filtro en su puntuación.
