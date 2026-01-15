-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL, 
    Description Text,
    published_at Timestamp not null,
    url text not null UNIQUE,
    feed_id UUID not null REFERENCES feeds(id) on DELETE CASCADE

);

-- +goose Down

DROP TABLE posts;