
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE servers (
  id SERIAL PRIMARY KEY,
  Name text NOT NULL UNIQUE,
  Url text NOT NULL UNIQUE,
  Role text NOT NULL,
  Created_at TIMESTAMP,
  Updated_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE servers;