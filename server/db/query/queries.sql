-- name: CreateUser :one
INSERT INTO users (username, email)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateItem :one
INSERT INTO items (owner_id, image_url, category, color, style)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: ListItemsByUser :many
SELECT * FROM items
WHERE owner_id = $1
ORDER BY created_at DESC;

-- name: CreateFriendship :exec
INSERT INTO friendships (user_id_1, user_id_2, status)
VALUES ($1, $2, $3);