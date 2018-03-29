
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE VIEW serverListView AS
  SELECT s.user_id AS user_id,
      s.url AS url,
      s.role AS role,
      s.id AS s_id,
      st.*

  FROM servers AS s
      LEFT JOIN ( SELECT s1.*
          FROM server_state as s1
          LEFT JOIN server_state AS s2
               ON s1.server_id = s2.server_id AND s1.created_at < s2.created_at
          WHERE s2.server_id IS NULL ) AS st
      ON (s.id = st.server_id)

  ORDER BY s.id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP VIEW IF EXISTS serverListView
