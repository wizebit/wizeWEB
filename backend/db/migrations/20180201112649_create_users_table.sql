
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

-- seed
  INSERT INTO users (Private_key, Public_key, Wallet, Role) VALUES (
    'a71cfc28e2f53a33646ff62bd096e35d5633d4888257e7e1b58e94bc7d9e129e
', 'b068fb2a52ef46889e4326ba3ad32f79ae26bee44d12ff4f849dbb9fbef6464a
', '4d5a774817197631c1fc61fbda8ae4b5267ac869d5db4349efb84407d99fc020
', '0'
  );
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;