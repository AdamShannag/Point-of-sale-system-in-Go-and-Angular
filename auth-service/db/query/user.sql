-- name: CreateUser :one
INSERT INTO users (uuid,
                   username,
                   email,
                   phone,
                   hashed_password,
                   address,
                   user_type,
                   added_by,
                   created_at,
                   modified_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE uuid = $1
   or username = $1
   or email = $1
   or phone = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET hashed_password   = COALESCE(sqlc.narg(hashed_password), hashed_password),
    username          = COALESCE(sqlc.narg(username), username),
    email             = COALESCE(sqlc.narg(email), email),
    phone             = COALESCE(sqlc.narg(phone), phone),
    address           = COALESCE(sqlc.narg(address), address),
    user_type         = COALESCE(sqlc.narg(user_type), user_type),
    modified_at     = sqlc.arg(modified_at)::timestamptz,
    created_at      = created_at
WHERE uuid = sqlc.arg(uuid)
RETURNING *;

-- name: UpdatePassword :execrows
UPDATE users
SET hashed_password = sqlc.arg(hashed_password),
    modified_at     = sqlc.arg(modified_at)::timestamptz
WHERE username = sqlc.arg(username);

-- name: GetUsername :one
SELECT username
FROM users
WHERE uuid = $1 LIMIT 1;