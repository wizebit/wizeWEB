
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
  CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    User_id int,
    Url text NOT NULL,
    Role text NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP
  );
  INSERT INTO servers (User_id, Url, Role,Created_at) VALUES
      ('1','127.0.0.1:8080/raft','raft',NOW()),
      ('1','127.0.0.1','blockchain',NOW()),
      ('1','127.0.0.1','storage',NOW()),
      ('2','127.0.0.1','raft',NOW()),
      ('2','127.0.0.1','blockchain',NOW()),
      ('2','127.0.0.1','storage',NOW()),
      ('3','127.0.0.1','raft',NOW()),
      ('3','127.0.0.1','blockchain',NOW()),
      ('3','127.0.0.1','storage',NOW());


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE servers;