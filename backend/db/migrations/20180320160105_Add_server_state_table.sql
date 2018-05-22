
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE server_state (
  id SERIAL PRIMARY KEY,
  Server_id int,
  Status boolean DEFAULT false,
  Latency int,
  Free_storage int,
  Uptime int,
  Rate int DEFAULT 0,
  Created_at TIMESTAMP,
  FOREIGN KEY (Server_id) REFERENCES Servers(id)
);

-- INSERT INTO server_state (Server_id, Latency, Free_storage, Uptime, Created_at) VALUES
--     ('1','100','100','100',NOW() - INTERVAL '1 HOUR'),
--     ('2','100','100','100',NOW() - INTERVAL '1 HOUR'),
--     ('3','100','100','100',NOW() - INTERVAL '1 HOUR'),
--     ('4','100','100','100',NOW() - INTERVAL '1 HOUR'),
--     ('5','100','100','100',NOW() - INTERVAL '1 HOUR'),
--     ('4','100','100','10',NOW()),
--     ('5','100','100','1000',NOW());

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE server_state CASCADE;