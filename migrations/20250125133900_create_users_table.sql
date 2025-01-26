-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES (
        '60dcee7a-fb13-4971-9d97-9b829bec0231',
        'admin',
        'test@test.ee',
        '2025-01-25 13:39:00',
        '2025-01-25 13:39:00'
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
