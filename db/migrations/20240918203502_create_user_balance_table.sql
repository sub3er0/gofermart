-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE user_balance (
    id SERIAL PRIMARY KEY,
    user_id INT,
    balance INT DEFAULT 0,
    used_balance INT DEFAULT 0
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
