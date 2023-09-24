-- Escenario 
-- Luego de un análisis realizado en un sistema de facturación, se ha detectado un mal diseño en la base de datos. La misma, cuenta con una tabla facturas que almacena datos de diferente naturaleza. 
-- Como se puede observar, la tabla cuenta con datos que podrían ser normalizados y separados en diferentes entidades.

-- Ejercicio 
-- Se solicita para el escenario anterior: 
-- Aplicar reglas de normalización y elaborar un modelo de DER que alcance la tercera forma normal (3FN).
-- Describir con sus palabras cada paso de la descomposición y aplicación de las reglas para visualizar el planteo realizado.

-- 0NF -> 1NF: Descomponer atributos compuestos en otra tabla (atomicidad)

-- 1NF -> 2NF: Migrar atributos que no sean dependientes de la clave principal (id factura) a sus propias tablas.

-- 2NF -> 3NF: Columnas que no son llave primaria, y que son dependientes entre si, se migran a otra tabla, dejando en la tabla original una referencia a la PK de la nueva tabla.