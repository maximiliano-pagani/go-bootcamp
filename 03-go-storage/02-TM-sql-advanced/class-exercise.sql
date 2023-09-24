-- Teniendo las tablas de una base de datos de una empresa:
-- Se requiere obtener las siguientes consultas:

-- Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.

SELECT e.nombre, e.puesto, d.localidad
	FROM empleados AS e
    INNER JOIN departamentos AS d ON e.depto_nro = d.depto_nro
    WHERE d.nombre_depto = "Ventas";

-- Visualizar los departamentos con más de 2 empleados.

SELECT d.depto_nro, d.nombre_depto, d.localidad
	FROM departamentos AS d 
	INNER JOIN empleados AS e ON d.depto_nro = e.depto_nro
    GROUP BY d.depto_nro
    HAVING COUNT(e.depto_nro) > 2;

-- Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.

SELECT e.nombre, e.salario, d.nombre_depto
	FROM empleados AS e
    INNER JOIN departamentos AS d ON e.depto_nro = d.depto_nro
    WHERE e.puesto = (
		SELECT e2.puesto
			FROM empleados as e2
			WHERE CONCAT(e2.nombre, " ", e2.apellido) = "Mito Barchuk"
    );

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.

SELECT e.*
	FROM empleados AS e
    WHERE e.depto_nro = (
		SELECT d.depto_nro
			FROM departamentos AS d
			WHERE d.nombre_depto = "Contabilidad"
	)
    ORDER BY e.nombre;

-- Mostrar el nombre del empleado que tiene el salario más bajo.

SELECT e.nombre
	FROM empleados AS e
    WHERE e.salario = (
		SELECT MIN(e2.salario)
			FROM empleados AS e2
    );

-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.

SELECT e.*
	FROM empleados AS e
    INNER JOIN departamentos AS d
    ON d.depto_nro = e.depto_nro
    WHERE d.nombre_depto = "Ventas"
    ORDER BY e.salario DESC LIMIT 1;
    
SELECT e.*
	FROM empleados AS e
    WHERE e.salario = (
		SELECT MAX(e2.salario)
			FROM empleados AS e2
            INNER JOIN departamentos AS d ON d.depto_nro = e2.depto_nro
            WHERE d.nombre_depto = "Ventas"
    );