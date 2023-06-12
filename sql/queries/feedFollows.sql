-- name: CreateFeedFollow :one
INSERT INTO feedFollows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFeedFollow :one
SELECT * FROM feedFollows 
WHERE id = $1;

-- name: GetFeedFollows :many
SELECT * FROM feedFollows
WHERE user_id = $1;

-- name: DeleteFeedFollow :one
DELETE FROM feedFollows WHERE id= $1
RETURNING *;