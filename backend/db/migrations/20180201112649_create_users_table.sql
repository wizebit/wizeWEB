
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

  INSERT INTO users (Private_key, Public_key, Address, Password, Role, Created_at) VALUES
    ('xWbemAt0z9oUX_xJZlC41ShWEztwwfcT3ijanM0tBVZMsa35pLyZQYr7SvOmhUDtGkAADlRJ5C1Ae_EXaB4CI0olRBHu_qoTsZwaa5C1th6mUBlDMxit0Ol55u5Xtiy8',
    '14868046059215250896028577020428367825674010679948899014672678699988209453275',
    '1JoPHwPYKLEhwaApUfrPLSPhEzzZR8cZB7',
    'c67068413da28e9f36e5ee4d68962b361bb2b3123361959744e417e005343ca0', 0,NOW());



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;