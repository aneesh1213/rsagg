-- +goose Up

CREATE TABLE feeds(
    id   UUID Primary Key,
    created_at Timestamp not NULL,
    updated_at Timestamp not NULL,
    name Text not NULL,
    url TEXT UNIQUE not null, 
    user_id UUID not null REFERENCES users(id) ON DELETE CASCADE 
);

-- +goose Down
DROP TABLE feeds;