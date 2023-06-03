-- name: CreateUser :one
INSERT INTO users (uuid,
                   username,
                   email,
                   phone,
                   hashed_password,
                   address,
                   user_type,
                   added_by)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;

-- name: GetUser :one
SELECT *
FROM users
WHERE uuid = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET hashed_password   = COALESCE(sqlc.narg(hashed_password), hashed_password),
    username          = COALESCE(sqlc.narg(username), username),
    email             = COALESCE(sqlc.narg(email), email),
    phone             = COALESCE(sqlc.narg(phone), phone),
    address           = COALESCE(sqlc.narg(address), address),
    user_type         = COALESCE(sqlc.narg(user_type), user_type),
    modified_at       = COALESCE(sqlc.narg(modified_at), modified_at),
    is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE uuid = sqlc.arg(uuid) RETURNING *;

-- name: GetUsername :one
SELECT username
FROM users
WHERE uuid = $1 LIMIT 1;