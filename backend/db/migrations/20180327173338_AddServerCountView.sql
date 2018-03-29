
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE VIEW serverCountView AS
  SELECT
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'blockchain') AS TotalBlockchainCount,
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'raft') AS TotalRaftCount,
    (SELECT COUNT(*) FROM servers WHERE servers.role = 'storage') AS TotalStorageCount,
    (SELECT COUNT(*)
      FROM servers AS s
          LEFT JOIN ( SELECT s1.*
              FROM server_state as s1
              LEFT JOIN server_state AS s2
                   ON s1.server_id = s2.server_id AND s1.created_at < s2.created_at
              WHERE s2.server_id IS NULL AND st.status = TRUE) AS st
          ON (s.id = st.server_id AND st.status = TRUE )) AS TotalSuspiciosCount;
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP VIEW IF EXISTS serverCountView
