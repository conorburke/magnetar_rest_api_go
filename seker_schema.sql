create table users (
  id serial primary key,
  first_name varchar NOT NULL DEFAULT '',
  last_name varchar NOT NULL DEFAULT '',
  email varchar not null NOT NULL DEFAULT '',
  password_hash varchar NOT NULL DEFAULT '',
  phone_number varchar unique NOT NULL DEFAULT '',
  birthdate date,
  address_1 varchar NOT NULL DEFAULT '',
  address_2 varchar NOT NULL DEFAULT '',
  city varchar NOT NULL DEFAULT '',
  region varchar NOT NULL DEFAULT '',
  zipcode integer NOT NULL DEFAULT 0
);

create table tools (
  id serial primary key,
  title varchar NOT NULL DEFAULT '',
  category varchar NOT NULL DEFAULT '',
  price double precision NOT NULL DEFAULT 0,
  tool_owner int references users NOT NULL DEFAULT -1,
  description varchar not null default ''
);

create table rented_tools (
  id serial primary key,
  tool_id int references tools NOT NULL DEFAULT -1,
  renter_id int references users NOT NULL DEFAULT -1,
  start_date timestamp with time zone,
  end_date timestamp with time zone 
);

create table tool_pictures (
  id serial primary key,
  tool_id int references tools NOT NULL DEFAULT-1,
  image varchar not null default ''
);