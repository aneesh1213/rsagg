-- name: CreateUser :one 

INSERT INTO users (id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256((id::text || now()::text || random()::text)::bytea), 'hex'))
RETURNING *;