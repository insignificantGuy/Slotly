-- +goose Up
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS slotly_align_companies;
CREATE PROCEDURE slotly_align_companies()
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'password'
    ) THEN
        ALTER TABLE companies DROP COLUMN password;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'name'
    ) THEN
        ALTER TABLE companies
            ADD COLUMN name VARCHAR(255) NOT NULL DEFAULT '' AFTER company_id,
            ADD COLUMN address VARCHAR(255) NOT NULL DEFAULT '' AFTER name,
            ADD COLUMN phone VARCHAR(255) NOT NULL DEFAULT '' AFTER address;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'logo'
    ) THEN
        ALTER TABLE companies
            ADD COLUMN logo VARCHAR(255) NOT NULL DEFAULT '' AFTER email,
            ADD COLUMN description VARCHAR(255) NOT NULL DEFAULT '' AFTER logo,
            ADD COLUMN is_active TINYINT(1) NOT NULL DEFAULT 1 AFTER description;
    END IF;
END;
CALL slotly_align_companies();
DROP PROCEDURE IF EXISTS slotly_align_companies;
-- +goose StatementEnd

ALTER TABLE companies MODIFY email VARCHAR(255) NOT NULL;

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS slotly_revert_companies;
CREATE PROCEDURE slotly_revert_companies()
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'name'
    ) THEN
        ALTER TABLE companies
            DROP COLUMN name,
            DROP COLUMN address,
            DROP COLUMN phone;
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'logo'
    ) THEN
        ALTER TABLE companies
            DROP COLUMN logo,
            DROP COLUMN description,
            DROP COLUMN is_active;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'companies'
          AND COLUMN_NAME = 'password'
    ) THEN
        ALTER TABLE companies ADD COLUMN password VARCHAR(255) AFTER company_id;
    END IF;
END;
CALL slotly_revert_companies();
DROP PROCEDURE IF EXISTS slotly_revert_companies;
-- +goose StatementEnd

ALTER TABLE companies MODIFY email VARCHAR(255) NULL;
