-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name character varying(50) NOT NULL
  );

-- +goose Down
DROP TABLE users;
