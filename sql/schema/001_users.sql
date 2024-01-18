-- + goose Up
CREATE TABEL users (
    id UUID PRIMARY KEY,
    created_at TIMESPTAMP NOT NULL,
    updated_at TIMESPTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;