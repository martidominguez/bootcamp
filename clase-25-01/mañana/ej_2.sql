-- create the tables

CREATE SCHEMA library;

USE library;

CREATE TABLE books (
	book_id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(40),
    editorial VARCHAR(40),
    area VARCHAR(40),
    PRIMARY KEY (book_id)
);

CREATE TABLE authors (
	author_id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(40),
    last_name VARCHAR(40),
    nationality VARCHAR(40),
    PRIMARY KEY (author_id)
);

CREATE TABLE books_authors (
	author_id INT NOT NULL,
    book_id INT NOT NULL,
    PRIMARY KEY (author_id, book_id),
    FOREIGN KEY (author_id) REFERENCES authors(author_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

CREATE TABLE students (
	student_id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(40),
    last_name VARCHAR(40),
    address VARCHAR(100),
    career VARCHAR(40),
    age INT, 
    PRIMARY KEY (student_id)
);

CREATE TABLE loans (
	student_id INT NOT NULL,
    book_id INT NOT NULL,
    loan_date DATE,
    return_date DATE,
    returned BOOL,
    PRIMARY KEY (student_id, book_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);

-- add random information to the books table
INSERT INTO books (title, editorial, area) VALUES
    ('The Great Gatsby', 'Scribner', 'Fiction'),
    ('To Kill a Mockingbird', 'HarperCollins', 'Classics'),
    ('Harry Potter and the Sorcerer''s Stone', 'Scholastic', 'Fantasy'),
    ('1984', 'Penguin Books', 'Dystopian'),
    ('The Hobbit', 'Houghton Mifflin', 'Fantasy');

-- add random information to the authors table
INSERT INTO authors (first_name, last_name, nationality) VALUES
    ('F. Scott', 'Fitzgerald', 'American'),
    ('Harper', 'Lee', 'American'),
    ('J.K.', 'Rowling', 'British'),
    ('George', 'Orwell', 'British'),
    ('J.R.R.', 'Tolkien', 'British');

-- add random information to the books_authors table
INSERT INTO books_authors (author_id, book_id) VALUES
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 4),
    (5, 5);

-- add random information to the students table
INSERT INTO students (first_name, last_name, address, career, age) VALUES
    ('John', 'Doe', '123 Main St', 'Computer Science', 20),
    ('Jane', 'Smith', '456 Oak St', 'Engineering', 22),
    ('Michael', 'Johnson', '789 Elm St', 'Business', 21),
    ('Emily', 'Williams', '101 Pine St', 'Biology', 23),
    ('Daniel', 'Brown', '202 Maple St', 'Psychology', 19);

-- add random information to the loans table
INSERT INTO loans (student_id, book_id, loan_date, return_date, returned) VALUES
    (1, 1, '2024-01-01', '2024-01-15', TRUE),
    (2, 2, '2024-02-01', '2024-02-15', FALSE),
    (3, 3, '2024-03-01', '2024-03-15', TRUE),
    (4, 4, '2024-04-01', '2024-04-15', FALSE),
    (5, 5, '2024-05-01', '2024-05-15', TRUE);

-- datos de los autores
SELECT * FROM authors;

-- nombre y edad de los estudiantes
SELECT first_name, last_name, age FROM students;

-- estudiantes pertenecen a la carrera informática
SELECT * FROM students WHERE career = 'Computer Science';

-- autores son de nacionalidad francesa o italiana
SELECT * FROM authors WHERE nationality = 'French' OR nationality = 'Italian';

-- libros no son del área de internet
SELECT * FROM books WHERE area <> 'Internet';

-- libros de la editorial Penguin Books
SELECT * FROM books WHERE editorial = 'Penguin Books';

-- datos de los estudiantes cuya edad es mayor al promedio
SELECT * FROM students WHERE age > (SELECT AVG(age) FROM students);

-- nombres de los estudiantes cuyo apellido comience con la letra B
SELECT first_name, last_name FROM students WHERE last_name LIKE 'B%';

-- autores del libro “1984”
SELECT first_name, last_name
FROM authors a 
	JOIN books_authors ba ON a.author_id = ba.author_id
	JOIN books b ON ba.book_id = b.book_id
WHERE b.title = '1984';

-- libros que se prestaron al lector “John Doe”
SELECT title
FROM books b 
	JOIN loans l ON b.book_id = l.book_id
	JOIN students s ON l.student_id = s.student_id
WHERE s.first_name = 'John' AND s.last_name = 'Doe';

-- nombre del estudiante de menor edad
SELECT first_name, last_name
FROM students
ORDER BY age
LIMIT 1;

-- nombres de los estudiantes a los que se prestaron libros de Fantasía
SELECT first_name, last_name
FROM students s
	JOIN loans l ON s.student_id = l.student_id
    JOIN books b ON l.book_id = b.book_id
WHERE b.area = 'Fantasy';

-- libros que pertenecen a la autora J.K. Rowling
SELECT title, editorial, area
FROM books b
	JOIN books_authors ba ON b.book_id = ba.book_id
    JOIN authors a ON ba.author_id = a.author_id
WHERE a.first_name = 'J.K.' AND a.last_name = 'Rowling';

-- títulos de los libros que debían devolverse el 15/01/2024.
SELECT title
FROM books b
	JOIN loans l ON b.book_id = l.book_id
WHERE l.return_date = '2024-01-15';
