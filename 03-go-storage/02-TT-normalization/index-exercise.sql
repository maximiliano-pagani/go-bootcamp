-- Ejercicio 1
-- Con la base de datos “movies”, se propone crear una tabla temporal llamada “TWD” y guardar en la misma los episodios de todas las temporadas de “The Walking Dead”.
-- Realizar una consulta a la tabla temporal para ver los episodios de la primera temporada.

USE movies_db;

CREATE TEMPORARY TABLE TWD
	SELECT e.id AS id_episode, e.title AS title_episode, e.rating, e.number AS number_episode,
		s.id AS id_season, s.title AS title_season, s.number as number_season
	FROM episodes AS e 
    INNER JOIN seasons AS s
    ON e.season_id = s.id
    INNER JOIN series
	ON series.id = s.serie_id
    WHERE series.title = "The Walking Dead";

SELECT * FROM TWD;

-- Ejercicio 2
-- En la base de datos “movies”, seleccionar una tabla donde crear un índice y luego chequear la creación del mismo. 
-- Analizar por qué crearía un índice en la tabla indicada y con qué criterio se elige/n el/los campos.

CREATE INDEX rating_index ON TWD (rating);

SELECT * FROM TWD WHERE rating >= 7.5;