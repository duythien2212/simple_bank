-- name: CreateUser :one
INSERT INTO users (
  username, 
  hash_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users 
SET 
  hash_password = coalesce(sqlc.narg(hash_password), hash_password),
  password_change_at = coalesce(sqlc.narg(password_change_at), password_change_at),
  full_name = coalesce(sqlc.narg(full_name), full_name),
  email = coalesce(sqlc.narg(email), email)
WHERE
  username = sqlc.arg(username)
RETURNING *;
