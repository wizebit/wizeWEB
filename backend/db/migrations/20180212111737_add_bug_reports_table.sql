
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE bug_reports (
  id SERIAL PRIMARY KEY,
  User_id int NOT NULL,
  Message text,
  Picture text,
  Created_at TIMESTAMP,
  Updated_at TIMESTAMP
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE bug_reports;
