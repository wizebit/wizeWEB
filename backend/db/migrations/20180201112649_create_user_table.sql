
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    Private_key text NOT NULL UNIQUE,
    Public_key text NOT NULL,
    Wallet text NOT NULL,
    Status boolean DEFAULT true,
    Role int DEFAULT 20,
    Rate int DEFAULT 0,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Salt text
  );

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;