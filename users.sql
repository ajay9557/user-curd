drop database if exists users;
create database users;
use users;

create table user(
                     id int NOT NULL PRIMARY KEY,
                     name VARCHAR(40) NOT NULL,
                     email VARCHAR(40) NOT NULL UNIQUE,
                     phone VARCHAR(10),
                     age int
);

INSERT INTO user VALUES (1, 'John', 'john21@example.com', '9728810299', 21);
INSERT INTO user VALUES (2, 'Jess', 'jessH90@example.com', '9844537168', 19);
