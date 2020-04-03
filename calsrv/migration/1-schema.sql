set timezone = 'Europe/Moscow';

create user calserv with encrypted password :basepass;

create database calserv owner calserv;
grant all privileges on database calserv to calserv;

\connect calserv
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

\connect calserv calserv

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name character varying(50) NOT NULL
  );


CREATE TABLE IF NOT EXISTS events (
  id UUID DEFAULT uuid_generate_v4() PRIMARY KEY  ,
  Owner_id integer NOT NULL,
  Title text NOT NULL,
  Text  text,
  StartTime timestamp without time zone NOT NULL,
  EndTime timestamp without time zone NOT NULL,
  foreign key(Owner_id) references users(id)
  );

CREATE INDEX owner_idx ON events (Owner_id);
CREATE INDEX StartTime_idx ON events (StartTime);


GRANT all privileges on all tables in schema public to calserv;
