-- name: CreateOtp :one
INSERT INTO otps (code, model, model_id, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOtpByID :one
SELECT * FROM otps
WHERE id = $1 LIMIT 1;

-- name: GetOtpByCodeAndModel :one
SELECT * FROM otps
WHERE code = $1 AND model = $2 AND model_id = $3
  AND expires_at > NOW()
LIMIT 1;

-- name: GetOtp :one
SELECT * FROM otps
WHERE code = $1
  AND model = $2
  AND expires_at > NOW()
LIMIT 1;


-- name: DeleteOtpByID :exec
DELETE FROM otps
WHERE id = $1;

-- name: DeleteExpiredOtps :exec
DELETE FROM otps
WHERE expires_at <= NOW();

