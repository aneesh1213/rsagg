-- +goose Up

CREATE TABLE feed_follows(
    id   UUID Primary Key,
    created_at Timestamp not NULL,
    updated_at Timestamp not NULL,
    user_id UUID not null REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID not null REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
); 

-- +goose Down
DROP TABLE feed_follows;