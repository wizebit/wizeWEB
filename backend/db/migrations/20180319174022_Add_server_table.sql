
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
  CREATE TABLE servers (
    id SERIAL PRIMARY KEY,
    User_id int,
    Name text,
    Url text NOT NULL,
    Role text NOT NULL,
    Created_at TIMESTAMP,
    Updated_at TIMESTAMP,
    FOREIGN KEY (User_id) REFERENCES Users(id)
  );
  INSERT INTO servers (User_id, Url, Role,Created_at) VALUES
      ('1','http://master.wizeprotocol.com:11000','raft',NOW()),
      ('1','http://master.wizeprotocol.com:4000','blockchain',NOW()),
      ('1','http://master.wizeprotocol.com:13000','storage',NOW()),
      ('1','http://sl1.wizeprotocol.com:13000','storage',NOW()),
      ('1','http://sl2.wizeprotocol.com:13000','storage',NOW());


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE servers;