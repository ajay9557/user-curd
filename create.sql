DROP DATABASE IF EXISTS UserInfo;
CREATE DATABASE UserInfo;
USE UserInfo;
CREATE TABLE User(
id int NOT NULL AUTO_INCREMENT,
name varchar(50),
email varchar(50),
phone varchar(50),
age int,
PRIMARY KEY(id));
INSERT INTO User VALUES(1,'Naira','Naira@gmail.com','9866895296',21);
INSERT INTO User VALUES(2,'Mahi','Mahi@gmail.com','6303844857',20);