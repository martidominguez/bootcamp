CREATE DATABASE internet_company;

USE internet_company;

CREATE TABLE customers (
	customer_id INT NOT NULL AUTO_INCREMENT,
    dni INT,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    birthdate DATE,
    province VARCHAR(255),
    city VARCHAR(255),
    PRIMARY KEY (customer_id)
);

CREATE TABLE internet_plans (
	internet_plan_id INT NOT NULL AUTO_INCREMENT,
    velocity INT,
    price DECIMAL(8,2),
    discount DECIMAL(8,2),
    PRIMARY KEY (internet_plan_id)
);

CREATE TABLE client_plan (
	booking_id INT NOT NULL AUTO_INCREMENT,
    customer_id INT NOT NULL,
    internet_plan_id INT NOT NULL,
    date DATE,
    PRIMARY KEY (booking_id),
    FOREIGN KEY (customer_id) REFERENCES customers(customer_id),
    FOREIGN KEY (internet_plan_id) REFERENCES internet_plans(internet_plan_id)
);

INSERT INTO customers (dni, first_name, last_name, birthdate, province, city)
VALUES 
	(1, 'Martina', 'Dominguez', '2003-05-14', 'Buenos Aires', 'Buenos Aires'),
    (2, 'Juan', 'Dominguez', '1958-11-20', 'Buenos Aires', 'Buenos Aires'),
    (3, 'Analia', 'Tcholakian', '1961-11-12', 'Buenos Aires', 'Buenos Aires'),
    (4, 'Juan', 'Perez', '2005-05-14', 'Buenos Aires', 'Buenos Aires');
    
SELECT * FROM customers; -- to ckeck the insert

INSERT INTO internet_plans (velocity, price, discount)
VALUES 
	(50, 105.8, 0),
    (100, 200.0, 0),
    (10, 10.5, 0);

SELECT * FROM internet_plans; -- to check the insert

INSERT INTO client_plan (customer_id, internet_plan_id, date)
VALUES
	(1, 1, '2024-01-12'),
    (2, 2, '2023-04-21'),
    (3, 3, '2021-12-22');

SELECT * FROM client_plan -- to check the insert