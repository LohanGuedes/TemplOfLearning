-- name: CreateUser :one
INSERT INTO users (username, email, password, role)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;
