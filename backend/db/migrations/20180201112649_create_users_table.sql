
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
  CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    Private_key text,
    Public_key text NOT NULL UNIQUE,
    Address text NOT NULL UNIQUE,
    Password text NOT NULL,
    Status boolean DEFAULT true,
    Role int DEFAULT 20,
    Rate int DEFAULT 0,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    Session_key text
  );

  INSERT INTO users (Public_key, Address, Password, Created_at) VALUES
    ('1111111','111111111','100',NOW()),
    ('2222222','222222222','100',NOW()),
    ('3333333','333333333','100',NOW());


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;