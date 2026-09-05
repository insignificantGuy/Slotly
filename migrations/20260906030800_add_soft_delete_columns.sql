-- +goose Up
ALTER TABLE users ADD COLUMN is_deleted TINYINT(1) NOT NULL DEFAULT 0;
ALTER TABLE users ADD COLUMN deleted_at DATETIME(3) NULL;
ALTER TABLE companies ADD COLUMN is_deleted TINYINT(1) NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE users DROP COLUMN is_deleted;
ALTER TABLE users DROP COLUMN deleted_at;
ALTER TABLE companies DROP COLUMN is_deleted;
