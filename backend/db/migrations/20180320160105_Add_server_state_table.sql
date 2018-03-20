
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

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE server_state;