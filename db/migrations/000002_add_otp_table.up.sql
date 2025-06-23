CREATE TABLE otps (
    id          bigserial PRIMARY KEY,
    code        VARCHAR(20) NOT NULL,
    model       VARCHAR(20) NOT NULL,
    model_id    VARCHAR(255) NOT NULL,
    expires_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW() + INTERVAL '1 hour',
    created_at  TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);
