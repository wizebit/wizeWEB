
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE VIEW serverCountView AS
  SELECT
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'blockchain') AS total_blockchain_count,
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'raft') AS total_raft_count,
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'storage') AS total_storage_count,
    (SELECT COUNT(*)
   FROM servers AS s
     LEFT JOIN ( SELECT s1.*
                 FROM server_state as s1
                   LEFT JOIN server_state AS s2
                     ON s1.server_id = s2.server_id AND s1.created_at < s2.created_at
                 WHERE s2.server_id IS NULL) AS st
       ON (s.id = st.server_id)
  WHERE st.status = FALSE) AS total_suspicios_count;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP VIEW serverCountView;
