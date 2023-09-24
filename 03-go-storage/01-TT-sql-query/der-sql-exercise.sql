-- Escenario
-- Una empresa proveedora de Internet necesita una base de datos para almacenar cada uno de sus clientes junto con el plan/pack que tiene contratado.
-- Mediante un análisis previo se conoce que se tiene que almacenar la siguiente información:
-- De los clientes se debe registrar: dni, nombre, apellido, fecha de nacimiento, provincia, ciudad.
-- En cuanto a los planes de internet: identificación del plan, velocidad ofrecida en megas, precio, descuento.


-- Ejercicio 1
-- Luego del planteo de los requerimientos de la empresa, se solicita modelar los mismos mediante un DER (Diagrama Entidad-Relación).


-- Ejercicio 2 
-- Una vez modelada y planteada la base de datos, responder a las siguientes preguntas:

-- a. ¿Cuál es la primary key para la tabla de clientes? Justificar respuesta

-- Si bien el DNI puede serlo sin problema para el caso de uso actual, normalmente se utilizan ids únicas para identificar entidades o recursos por siempre a nivel negocio, ya que por ejemplo, un cliente podría cambiar la titularidad del servicio a otro DNI. Por ello la PK elegida es el atributo id.

-- b. ¿Cuál es la primary key para la tabla de planes de internet? Justificar respuesta.

-- En este caso la PK va ser un número único, invariante y autoincremental que identifique al plan. Luego, podría tener otro código único mas legible y amigable para humanos como por ejemplo 300MB_4K para la operatoria diaria.

-- c. ¿Cómo serían las relaciones entre tablas? ¿En qué tabla debería haber foreign key? ¿A qué campo de qué tabla hace referencia dicha foreign key? Justificar respuesta.

-- Dado que por enunciado deduje un cliente sólo puede tener un plan contratado, la relación entre cliente y plan es "n a 1", por lo que no hace falta una tabla pivot o intermedia. Con ubicar un atributo en la entidad de clientes que referencie al identificador de plan que posee es suficiente.

-- Ejercicio 3
-- Una vez realizado el planteo del diagrama y de haber respondido estas preguntas, utilizar PHPMyAdmin o MySQL Workbench para ejecutar lo siguiente:
-- Se solicita crear una nueva base de datos llamada “empresa_internet”. 
-- Incorporar 10 registros en la tabla de clientes y 5 en la tabla de planes de internet.
-- Realizar las asociaciones/relaciones correspondientes entre estos registros.

-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema empresa_internet
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema empresa_internet
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `empresa_internet` DEFAULT CHARACTER SET utf8 ;
USE `empresa_internet` ;

-- -----------------------------------------------------
-- Table `empresa_internet`.`plans`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`plans` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `mbps` INT NOT NULL,
  `price` FLOAT NOT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `empresa_internet`.`clients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `empresa_internet`.`clients` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `birthdate` DATETIME NOT NULL,
  `dni` INT NOT NULL,
  `state` VARCHAR(45) NOT NULL,
  `city` VARCHAR(45) NOT NULL,
  `id_plan` INT NULL,
  `discount` FLOAT ZEROFILL NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `dni_UNIQUE` (`dni` ASC) VISIBLE,
  INDEX `FK_id_plan_idx` (`id_plan` ASC) VISIBLE,
  CONSTRAINT `FK_id_plan`
    FOREIGN KEY (`id_plan`)
    REFERENCES `empresa_internet`.`plans` (`id`)
    ON DELETE SET NULL
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;

-- INSERT DB DATA

INSERT INTO `empresa_internet`.`plans`
	(`mbps`, `price`)
VALUES
	(50,1499),
  (100,2999),
  (300,4499),
  (500,5599),
  (1000,6999);

INSERT INTO `empresa_internet`.`clients`
	(`first_name`,`last_name`,`birthdate`,`dni`,`state`,`city`,`id_plan`,`discount`)
VALUES
  ("Pedro","Gómez",NOW(),34934011,"Buenos Aires","Avellaneda",3,0.3),
  ("Lucas","Fernandez",NOW(),27934912,"Mendoza","San Rafael",5,0.0),
  ("Andrea","Guno",NOW(),38493494,"Buenos Aires","Vicente Lopez",1,0.3),
  ("Morena","Pagliano",NOW(),15934955,"Córdoba","Cordoba",4,0.2),
  ("Santiago","Perez",NOW(),49425802,"Buenos Aires","Moreno",1,0.5),
  ("Lucía","Rossi",NOW(),12495053,"Buenos Aires","CABA",2,0.0),
  ("Agustina","Zabaleta",NOW(),54924001,"Santa Fe","Rosario",5,0.7),
  ("Germán","None",NOW(),38593150,"Misiones","Posadas",3,0.1),
  ("Esteban","Álvarez",NOW(),28519240,"Buenos Aires","Bahía Blanca",3,0.5),
  ("Guillermo","Robledo",NOW(),33857185,"La Pampa","Santa Rosa",2,0.0),
  ("Jimena","Suarez",NOW(),27819538,"Buenos Aires","CABA",1,0.2);

-- Ejercicio 4
-- Plantear 10 consultas SQL que se podrían realizar a la base de datos. Expresar las sentencias.

USE empresa_internet;

SELECT COUNT(id) FROM clients;

SELECT clients.state, COUNT(id)
	FROM clients
	GROUP BY clients.state;

SELECT clients.state, clients.city, COUNT(id)
	FROM clients
	GROUP BY clients.state, clients.city
	ORDER BY state;

SELECT plans.mbps, plans.price
	FROM plans
	ORDER BY plans.price DESC;

SELECT COUNT(*)
	FROM clients
    WHERE discount >= 0.5;

SELECT clients.first_name, clients.last_name, plans.mbps FROM clients
	INNER JOIN plans ON clients.id_plan = plans.id
    WHERE plans.mbps >= 500;

SELECT clients.birthdate, YEAR(NOW()) - YEAR(clients.birthdate) as edad
	FROM clients
    WHERE YEAR(NOW()) - YEAR(clients.birthdate) >= 30;

SELECT clients.*
	FROM clients
    WHERE discount = 0.0 AND city = "CABA";