CREATE DATABASE employees;
CREATE USER 'user' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON employees.* TO 'user';
FLUSH PRIVILEGES;
USE `employees`;

CREATE TABLE `users`
(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    lastname VARCHAR(255),
    datetime DATETIME,
    address VARCHAR(255),
    gender VARCHAR(10),
    age INT
);

