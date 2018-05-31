create table users (
  id serial primary key,
  first_name varchar NOT NULL DEFAULT '',
  last_name varchar NOT NULL DEFAULT '',
  email varchar not null NOT NULL DEFAULT '',
  password_hash varchar NOT NULL DEFAULT '',
  phone_number varchar unique NOT NULL DEFAULT '',
  age integer NOT NULL DEFAULT 0,
  address_1 varchar NOT NULL DEFAULT '',
  address_2 varchar NOT NULL DEFAULT '',
  city varchar NOT NULL DEFAULT '',
  region varchar NOT NULL DEFAULT '',
  zipcode integer NOT NULL DEFAULT 0
);

create table tools (
  id serial primary key,
  title varchar NOT NULL DEFAULT '',
  tool_type varchar NOT NULL DEFAULT '',
  price double precision NOT NULL DEFAULT 0,
  tool_owner int references users NOT NULL DEFAULT -1
);