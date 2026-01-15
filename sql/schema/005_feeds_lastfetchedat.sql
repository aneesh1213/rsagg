-- +goose Up

alter TABLE feeds add column last_fetched_at Timestamp;


-- +goose Down 
ALTER TABLE feeds drop column last_fetched_at;