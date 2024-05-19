-- name: GetEmailByUUID :one
SELECT * FROM email WHERE uuid = $1;

-- name: GetEmails :many
SELECT * FROM email;

-- name: InsertEmail :one
INSERT INTO email (recipient, subject) VALUES ($1, $2) RETURNING *;

-- name: GetEmailsWithPagination :many
SELECT * FROM email ORDER BY send_date DESC LIMIT $1 OFFSET $2;