-- name: CreateUser :one
INSERT INTO users(id,created_At,updated_At,name)
VALUES ($1, $2,$3, $4)
RETURNING *;
