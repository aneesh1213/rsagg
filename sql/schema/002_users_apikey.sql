-- +goose Up
ALTER TABLE users ADD COLUMN api_key varchar(64);

UPDATE users SET api_key = encode(sha256((id::text || now()::text || random()::text)::bytea), 'hex');

ALTER TABLE users ADD CONSTRAINT users_api_key_unique UNIQUE(api_key);

ALTER TABLE users ALTER COLUMN api_key SET NOT NULL;

-- +goose Down
ALTER TABLE users DROP COLUMN api_key;