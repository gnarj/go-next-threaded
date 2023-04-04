-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE IF EXISTS todos.todos
ADD COLUMN duration numeric DEFAULT 0;
ALTER TABLE IF EXISTS todos.todos
ADD COLUMN durationUnit text;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE IF EXISTS todos.todos
DROP COLUMN duration;
ALTER TABLE IF EXISTS todos.todos
DROP COLUMN durationUnit;
-- +goose StatementEnd
