-- name: InsertTracker :one
INSERT INTO tracker (email_uuid, ip_address) VALUES ($1, $2) RETURNING *;

-- name: GetTrackerByUUID :one
SELECT * FROM tracker WHERE email_uuid = $1;

-- name: GetTrackers :many
SELECT * FROM tracker;

-- name: GetTrackersWithPagination :many
SELECT * FROM tracker ORDER BY open_date DESC LIMIT $1 OFFSET $2;

