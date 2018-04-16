create table users (
  id serial primary key,
  first_name varchar,
  last_name varchar,
  email varchar not null unique,
  password_hash varchar,
  phone_number varchar,
  age integer,
  address_1 varchar,
  address_2 varchar,
  city varchar,
  region varchar,
  zipcode integer
);

create table tools (
  id serial primary key,
  title varchar,
  tool_type double precision,
  price varchar,
  tool_owner varchar references users(email)
);