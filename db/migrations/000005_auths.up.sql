CREATE TABLE IF NOT EXISTS auths (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL,
    token VARCHAR NOT NULL,
    auth_type VARCHAR NOT NULL,
    issued_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    expired_at TIMESTAMPTZ NOT NULL,

    FOREIGN KEY(user_id) REFERENCES users(id)
)