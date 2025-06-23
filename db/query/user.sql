-- name: CreateUser :one
INSERT INTO users (first_name,last_name, email, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET
  first_name = COALESCE(sqlc.narg('first_name'), first_name),
  last_name = COALESCE(sqlc.narg('last_name'), last_name),
  phone = COALESCE(sqlc.narg('phone'), phone),
  country = COALESCE(sqlc.narg('country'), country),
  city = COALESCE(sqlc.narg('city'), city),
  avatar_url = COALESCE(sqlc.narg('avatar_url'), avatar_url),
  is_email_verified = COALESCE(sqlc.narg('is_email_verified'), is_email_verified),
  updated_at = NOW()
WHERE id = sqlc.arg('id')
RETURNING *;
