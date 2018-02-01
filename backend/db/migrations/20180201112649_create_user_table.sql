
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    Name  text,
    First_name text,
    Last_name text,
    Email text NOT NULL UNIQUE,
    Password text NOT NULL,
    Role int DEFAULT 20,
    Rate int DEFAULT 0,
    Status boolean DEFAULT true,
    Created_at TIMESTAMP,
    Salt text
  );

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;