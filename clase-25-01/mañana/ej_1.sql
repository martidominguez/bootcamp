-- tenemos dos tablas: empleado y departamento

-- nombre, puesto y localidad de los departamentos donde trabajan los vendedores
SELECT d.nombre_depto, e.puesto, d.localidad
FROM departamento d 
JOIN empleado e ON d.depto_nro = e.depto_nro;

-- departamentos con más de cinco empleados
SELECT d.nombre_depto
FROM departamento d
JOIN empleado e ON d.depto_nro = e.depto_nro
GROUP BY d.nombre_depto
HAVING COUNT(*) > 5;

-- nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’
SELECT e.nombre, e.salario, d.nombre_depto
FROM departamento d
JOIN empleado e ON d.depto_nro = e.depto_nro
WHERE e.puesto = (SELECT puesto FROM empleados WHERE nombre = 'Mito' AND apellido = 'Barchuck');

-- datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre
SELECT e.nombre, e.apellido, e.puesto, e.fecha_alta, e.salario, e.comision
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Contabilidad'
ORDER BY e.nombre;

-- nombre del empleado que tiene el salario más bajo
SELECT nombre
FROM empleado
ORDER BY salario
LIMIT 1;

-- datos del empleado que tiene el salario más alto en el departamento de ‘Ventas'
SELECT e.nombre, e.apellido, e.puesto, e.fecha_alta, e.salario, e.comision
FROM empleado e
JOIN departamento d ON e.depto_nro = d.depto_nro
WHERE d.nombre_depto = 'Ventas'
ORDER BY e.salario
LIMIT 1;


