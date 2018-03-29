
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE server_state (
  id SERIAL PRIMARY KEY,
  Server_id int,
  Ip text NOT NULL,
  Status boolean DEFAULT false,
  Latency int,
  Free_storage int,
  Uptime int,
  Type_active boolean DEFAULT false,
  Rate int DEFAULT 0,
  Created_at TIMESTAMP
);

INSERT INTO server_state (Server_id, Ip, Latency, Free_storage, Uptime, Created_at) VALUES
    ('1','127.0.0.1','100','100','100',NOW()),
    ('2','127.0.0.1','100','100','100',NOW()),
    ('3','127.0.0.1','100','100','100',NOW()),
    ('4','127.0.0.1','100','100','100',NOW()),
    ('5','127.0.0.1','100','100','100',NOW()),
    ('6','127.0.0.1','100','100','100',NOW()),
    ('7','127.0.0.1','100','100','100',NOW()),
    ('8','127.0.0.1','100','100','100',NOW()),
    ('4','127.0.0.1','100','100','10',NOW()),
    ('5','127.0.0.1','100','100','1000',NOW());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE server_state;