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

CREATE TYPE task AS
(
    name text,
    description text
);

CREATE TYPE endpoint AS
(
    tasks task[],
    date_started date
);


CREATE TABLE discipline
(
   id serial not null unique,
   name varchar(255) not null,
   endpoints endpoint[] not null,
   groups varchar(255)[] not null
);

CREATE TABLE users_discipline
(
    id serial not null unique,
    discipline_id int references discipline(id) on delete cascade not null,
    users_id int references users(id) on delete cascade not null
);


