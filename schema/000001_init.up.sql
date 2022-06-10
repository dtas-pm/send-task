CREATE TABLE students
(
  id serial not null unique,
  fullname varchar(255) not null,
  login varchar(255) not null unique,
  email varchar(255)[] not null,
  student_group varchar(255) not null,
  institute varchar(255) not null
);

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    email varchar(255) not null unique,
    role varchar(255) not null,
    password_hash varchar(255) not null
);

CREATE TABLE discipline
(
   id serial not null unique,
   name varchar(255),
   endpoints jsonb,
   groups varchar(255),
   unique (name, groups)
);

CREATE TABLE users_discipline
(
    id serial not null unique,
    discipline_id int references discipline(id) on delete cascade not null,
    users_id int references users(id) on delete cascade not null
);

CREATE TABLE plan_discipline
(
   id serial not null unique,
   name varchar(255) not null,
   date_start date not null,
   endpoints jsonb,
   groups varchar(255) not null,
   unique(name, groups)
);

CREATE TABLE users_plan_discipline
(
    id serial not null unique,
    plan_discipline_id int references plan_discipline(id) on delete cascade not null,
    users_id int references users(id) on delete cascade not null
);

CREATE TABLE groups
(
   id serial not null unique,
   name varchar(255) not null unique
);


CREATE OR REPLACE FUNCTION group_insert() RETURNS trigger AS $$
    BEGIN
        INSERT INTO groups(name) VALUES (NEW.student_group) ON CONFLICT DO NOTHING ;
    RETURN NEW;
    END;
    $$
    LANGUAGE 'plpgsql';

CREATE TRIGGER t_students
    AFTER INSERT
    ON students
    FOR EACH ROW
    EXECUTE PROCEDURE group_insert();

INSERT INTO users (name, username, email, role, password_hash) VALUES ('Admin', 'Admin', 'samarec1812@mail.ru', 'admin',  '6e61666f6461696e64613932333732626e696e616e7132726665e84be33532fb784c48129675f9eff3a682b27168c0ea744b2cf58ee02337c5');