CREATE DATABASE IF NOT EXISTS stori_db;

CREATE TABLE IF NOT EXISTS user ( user_id bigint not null AUTO_INCREMENT primary key, email varchar(30) not null unique, balance FLOAT);

create table if not exists transactions ( id bigint not null AUTO_INCREMENT primary key, file_name varchar(30) not null, created_date date, amount FLOAT, user_owner bigint not null,  FOREIGN KEY (user_owner) REFERENCES user(user_id));

insert ignore into user (email, balance) VALUES ('patricioyegros@hotmail.com', 7.70), ('yegrospatricio@gmail.com', 0.0);