CREATE TABLE IF NOT EXISTS user ( user_id bigint not null primary key, email varchar(30) not null unique, balance FLOAT);

create table if not exists transactions ( id bigint not null AUTO_INCREMENT primary key , file_name varchar(30) not null, created_date date, amount FLOAT, user_owner bigint not null,  FOREIGN KEY (user_owner) REFERENCES user(user_id));

INSERT INTO user (user_id, email, balance) SELECT * FROM (SELECT (SELECT MAX(user_id)+1 AS id FROM user),'patricioyegros@hotmail.com', 100.0) AS tmp WHERE NOT EXISTS (SELECT email from user where email = 'patricioyegros@hotmail.com') LIMIT 1;

INSERT INTO user (user_id, email, balance) SELECT * FROM (SELECT (SELECT MAX(user_id)+1 AS id FROM user), 'yegrospatricio@gmail.com', 100.0) AS tmp WHERE NOT EXISTS (SELECT email from user where email = 'yegrospatricio@gmail.com') LIMIT 1;