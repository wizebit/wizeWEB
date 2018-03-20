
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE servers (
  id SERIAL PRIMARY KEY,
  Url text NOT NULL UNIQUE,
  Ip text NOT NULL,
  Status boolean DEFAULT true,
  Role text NOT NULL,
  Rate int DEFAULT 0,
  Created_at TIMESTAMP,
  Updated_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE servers;