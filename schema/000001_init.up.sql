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
    password_hash varchar(255) not null
);

CREATE TABLE discipline
(
   id serial not null unique,
   name varchar(255),
   endpoints jsonb,
   groups varchar(255)
);

CREATE TABLE users_discipline
(
    id serial not null unique,
    discipline_id int references discipline(id) on delete cascade not null,
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
