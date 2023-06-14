-- name: UpdateKeys :one
UPDATE key_pair
SET privet_key  = COALESCE(sqlc.narg(privet_key), privet_key),
    public_key  = COALESCE(sqlc.narg(public_key), public_key),
    expired_at  = (now() + interval '60 minutes'),
    modified_at =now() RETURNING *;

-- name: GetKeys :one
SELECT *
FROM key_pair LIMIT 1;