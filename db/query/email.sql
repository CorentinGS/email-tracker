-- name: GetEmailByUUID :one
SELECT * FROM email WHERE uuid = $1;

-- name: GetEmails :many
SELECT * FROM email;

-- name: GetTrackerByUUID :one
SELECT * FROM tracker WHERE email_uuid = $1;

-- name: GetTrackers :many
SELECT * FROM tracker;

-- name: InsertEmail :one
INSERT INTO email (recipient, subject) VALUES ($1, $2) RETURNING *;

-- name: InsertTracker :one
INSERT INTO tracker (email_uuid, ip_address) VALUES ($1, $2) RETURNING *;