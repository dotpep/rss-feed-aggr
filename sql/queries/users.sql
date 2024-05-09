-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username, api_key)
VALUES ($1, $2, $3, $4,
    -- generate 64 varchar sized api_key
    encode(sha256(random()::text::bytea), 'hex')
)
RETURNING *;

-- name: GetUserByAPIKey :one
SELECT * FROM users
WHERE api_key = $1;
