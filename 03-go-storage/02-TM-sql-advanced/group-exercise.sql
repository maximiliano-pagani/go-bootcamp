-- Se tiene el siguiente DER que corresponde al esquema que presenta la base de datos de una “biblioteca”.
-- En base al mismo, plantear las consultas SQL para resolver los siguientes requerimientos:

DROP DATABASE IF EXISTS biblioteca;
CREATE DATABASE biblioteca;
USE biblioteca;

CREATE TABLE `libro` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `titulo` varchar(500) NOT NULL,
  `editorial` varchar(200) NOT NULL,
  `area` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `autor` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nombre` varchar(200) NOT NULL,
  `nacionalidad` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `libro_autor` (
  `id_autor` int(10) unsigned NOT NULL,
  `id_libro` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id_autor`, `id_libro`),
  KEY `FK_libro_autor_id_autor` (`id_autor`),
  KEY `FK_libro_autor_id_libro` (`id_libro`),
  CONSTRAINT `FK_libro_autor_id_autor` FOREIGN KEY (`id_autor`) REFERENCES `autor` (`id`),
  CONSTRAINT `FK_libro_autor_id_libro` FOREIGN KEY (`id_libro`) REFERENCES `libro` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `estudiante` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `nombre` varchar(200) NOT NULL,
  `apellido` varchar(200) NOT NULL,
  `direccion` varchar(500) NOT NULL,
  `carrera` varchar(100) NOT NULL,
  `edad` int(3) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

CREATE TABLE `prestamo` (
  `id_lector` int(10) unsigned NOT NULL,
  `id_libro` int(10) unsigned NOT NULL,
  `fecha_prestamo` datetime NOT NULL,
  `fecha_devolucion` datetime NULL,
  `devuelto` boolean NOT NULL DEFAULT false,
  PRIMARY KEY (`id_lector`, `id_libro`, `fecha_prestamo`),
  KEY `FK_prestamo_id_lector` (`id_lector`),
  KEY `FK_prestamo_id_libro` (`id_libro`),
  CONSTRAINT `FK_prestamo_id_lector` FOREIGN KEY (`id_lector`) REFERENCES `estudiante` (`id`),
  CONSTRAINT `FK_prestamo_id_libro` FOREIGN KEY (`id_libro`) REFERENCES `libro` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

INSERT INTO libro (titulo, editorial, area) VALUES
	("Martín Fierro", "Alfaguara", "Nacional"),
    ("Ficciones", "Salamandra", "Contemporáneo"),
    ("Harry Potter", "De Bolsillo", "Fantasía"),
    ("El Universo: Guía de viaje", "Lonely Planet", "Astronomía");
    
INSERT INTO autor (nombre, nacionalidad) VALUES
	("José Hernandez", "Argentina"),
    ("Jorge Luis Borges", "Argentina"),
    ("J.K. Rowling", "Reino unido"),
    ("Gabriel García Márquez", "Colombia"),
    ("Valerio Massimo Manfredi", "Italia"),
    ("Oliver Berry", "Inglaterra");

INSERT INTO estudiante (nombre, apellido, direccion, carrera, edad) VALUES
	("José", "Murillo", "Santa Fe 3443", "Industrial", 24),
    ("Luis", "Gonzalez", "Lope de Vega 1255", "Filosofía", 28),
    ("Filippo", "Galli", "Paseo Colón 948", "Abogacía", 17),
    ("Agustina", "Ferreyra", "San Martín 594", "Informática", 23);

INSERT INTO libro_autor VALUES
	(1, 1),
	(2, 2),
	(3, 3),
    (6, 4);
 
INSERT INTO prestamo VALUES
	(1, 2, NOW(), null, false),
    (3, 3, NOW(), NOW(), true),
    (4, 3, NOW(), null, false);

-- Listar los datos de los autores.

SELECT * FROM autor;

-- Listar nombre y edad de los estudiantes

SELECT CONCAT(nombre, " ", apellido) as nombre, edad FROM estudiante;

-- ¿Qué estudiantes pertenecen a la carrera informática?

SELECT CONCAT(nombre, " ", apellido) as nombre FROM estudiante WHERE carrera = "Informática";

-- ¿Qué autores son de nacionalidad francesa o italiana?

SELECT nombre FROM autor WHERE nacionalidad = "Italia" OR nacionalidad = "Francia";

-- ¿Qué libros no son del área de internet?

SELECT titulo FROM libro WHERE area <> "Internet";

-- Listar los libros de la editorial Salamandra.

SELECT titulo from LIBRO WHERE editorial = "Salamandra";

-- Listar los datos de los estudiantes cuya edad es mayor al promedio.

SELECT * FROM estudiante
	WHERE edad > (
		SELECT AVG(e.edad) FROM estudiante as e
    );

-- Listar los nombres de los estudiantes cuyo apellido comience con la letra G.

SELECT * FROM estudiante WHERE apellido LIKE "G%";

-- Listar los autores del libro “El Universo: Guía de viaje”. (Se debe listar solamente los nombres).

SELECT autor.nombre
	FROM autor
    WHERE autor.id IN (
		SELECT la.id_autor
			FROM libro_autor as la
            INNER JOIN libro as l
            ON la.id_libro = l.id
            WHERE l.titulo = "El Universo: Guía de viaje"
    );

-- ¿Qué libros se prestaron al lector “Filippo Galli”?

SELECT titulo
	FROM libro
    INNER JOIN prestamo ON libro.id = prestamo.id_libro
    INNER JOIN estudiante ON estudiante.id = prestamo.id_lector
    WHERE estudiante.nombre = "Filippo" AND estudiante.apellido = "Galli";

-- Listar el nombre del estudiante de menor edad.
    
SELECT nombre, apellido
	FROM estudiante
	WHERE edad = (
		SELECT MIN(e.edad) FROM estudiante as e
    );

-- Listar nombres de los estudiantes a los que se prestaron libros de Base de Datos.

SELECT nombre, apellido 
	FROM estudiante
    INNER JOIN prestamo ON estudiante.id = prestamo.id_lector
    INNER JOIN libro ON libro.id = prestamo.id_libro
    WHERE libro.area = "Base de Datos";

-- Listar los libros que pertenecen a la autora J.K. Rowling.

SELECT titulo
	FROM libro
    INNER JOIN libro_autor ON libro.id = libro_autor.id_libro
    INNER JOIN autor ON autor.id = libro_autor.id_autor
    WHERE autor.nombre = "J.K. Rowling";

-- Listar títulos de los libros que debían devolverse el 24/09/2023.

SELECT titulo
	FROM libro
    INNER JOIN prestamo ON libro.id = prestamo.id_libro
    WHERE DATE(prestamo.fecha_devolucion) = "2023-09-24";
