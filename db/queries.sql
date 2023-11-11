-- name: CreateEvent :one
INSERT INTO events (
        name,
        description,
        start_date,
        end_date,
        one_time
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetEvent :one
SELECT *
FROM events
WHERE id = $1;
-- name: GetEvents :many
SELECT *
FROM events OFFSET $1
LIMIT $2;