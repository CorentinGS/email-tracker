// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: email.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getEmailByUUID = `-- name: GetEmailByUUID :one
SELECT recipient, subject, send_date, uuid FROM email WHERE uuid = $1
`

func (q *Queries) GetEmailByUUID(ctx context.Context, argUuid uuid.UUID) (Email, error) {
	row := q.db.QueryRow(ctx, getEmailByUUID, argUuid)
	var i Email
	err := row.Scan(
		&i.Recipient,
		&i.Subject,
		&i.SendDate,
		&i.Uuid,
	)
	return i, err
}

const getEmails = `-- name: GetEmails :many
SELECT recipient, subject, send_date, uuid FROM email
`

func (q *Queries) GetEmails(ctx context.Context) ([]Email, error) {
	rows, err := q.db.Query(ctx, getEmails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Email
	for rows.Next() {
		var i Email
		if err := rows.Scan(
			&i.Recipient,
			&i.Subject,
			&i.SendDate,
			&i.Uuid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTrackerByUUID = `-- name: GetTrackerByUUID :one
SELECT tracker_id, open_date, email_uuid, ip_address FROM tracker WHERE email_uuid = $1
`

func (q *Queries) GetTrackerByUUID(ctx context.Context, emailUuid uuid.UUID) (Tracker, error) {
	row := q.db.QueryRow(ctx, getTrackerByUUID, emailUuid)
	var i Tracker
	err := row.Scan(
		&i.TrackerID,
		&i.OpenDate,
		&i.EmailUuid,
		&i.IpAddress,
	)
	return i, err
}

const getTrackers = `-- name: GetTrackers :many
SELECT tracker_id, open_date, email_uuid, ip_address FROM tracker
`

func (q *Queries) GetTrackers(ctx context.Context) ([]Tracker, error) {
	rows, err := q.db.Query(ctx, getTrackers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Tracker
	for rows.Next() {
		var i Tracker
		if err := rows.Scan(
			&i.TrackerID,
			&i.OpenDate,
			&i.EmailUuid,
			&i.IpAddress,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertEmail = `-- name: InsertEmail :one
INSERT INTO email (recipient, subject) VALUES ($1, $2) RETURNING recipient, subject, send_date, uuid
`

type InsertEmailParams struct {
	Recipient string
	Subject   string
}

func (q *Queries) InsertEmail(ctx context.Context, arg InsertEmailParams) (Email, error) {
	row := q.db.QueryRow(ctx, insertEmail, arg.Recipient, arg.Subject)
	var i Email
	err := row.Scan(
		&i.Recipient,
		&i.Subject,
		&i.SendDate,
		&i.Uuid,
	)
	return i, err
}

const insertTracker = `-- name: InsertTracker :one
INSERT INTO tracker (email_uuid, ip_address) VALUES ($1, $2) RETURNING tracker_id, open_date, email_uuid, ip_address
`

type InsertTrackerParams struct {
	EmailUuid uuid.UUID
	IpAddress *string
}

func (q *Queries) InsertTracker(ctx context.Context, arg InsertTrackerParams) (Tracker, error) {
	row := q.db.QueryRow(ctx, insertTracker, arg.EmailUuid, arg.IpAddress)
	var i Tracker
	err := row.Scan(
		&i.TrackerID,
		&i.OpenDate,
		&i.EmailUuid,
		&i.IpAddress,
	)
	return i, err
}
