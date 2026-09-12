-- +goose Up
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS slotly_add_company_role_user_id;
CREATE PROCEDURE slotly_add_company_role_user_id()
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND COLUMN_NAME = 'user_id'
    ) THEN
        ALTER TABLE company_role_mappings
            ADD COLUMN user_id CHAR(36) NULL AFTER deleted_at;
    END IF;
END;
CALL slotly_add_company_role_user_id();
DROP PROCEDURE IF EXISTS slotly_add_company_role_user_id;
-- +goose StatementEnd

DELETE FROM company_role_mappings
WHERE user_id IS NULL;

-- +goose StatementBegin
DROP PROCEDURE IF EXISTS slotly_index_company_role_user_id;
CREATE PROCEDURE slotly_index_company_role_user_id()
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.STATISTICS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND INDEX_NAME = 'idx_company_role'
    ) THEN
        ALTER TABLE company_role_mappings DROP INDEX idx_company_role;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.STATISTICS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND INDEX_NAME = 'idx_user_company'
    ) THEN
        ALTER TABLE company_role_mappings
            MODIFY user_id CHAR(36) NOT NULL,
            ADD UNIQUE KEY idx_user_company (user_id, company_id);
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.TABLE_CONSTRAINTS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND CONSTRAINT_NAME = 'fk_company_role_mappings_user'
    ) THEN
        ALTER TABLE company_role_mappings
            ADD CONSTRAINT fk_company_role_mappings_user
                FOREIGN KEY (user_id) REFERENCES users (user_id)
                ON UPDATE CASCADE ON DELETE CASCADE;
    END IF;
END;
CALL slotly_index_company_role_user_id();
DROP PROCEDURE IF EXISTS slotly_index_company_role_user_id;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP PROCEDURE IF EXISTS slotly_drop_company_role_user_id;
CREATE PROCEDURE slotly_drop_company_role_user_id()
BEGIN
    IF EXISTS (
        SELECT 1 FROM information_schema.TABLE_CONSTRAINTS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND CONSTRAINT_NAME = 'fk_company_role_mappings_user'
    ) THEN
        ALTER TABLE company_role_mappings DROP FOREIGN KEY fk_company_role_mappings_user;
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.STATISTICS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND INDEX_NAME = 'idx_user_company'
    ) THEN
        ALTER TABLE company_role_mappings DROP INDEX idx_user_company;
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.COLUMNS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND COLUMN_NAME = 'user_id'
    ) THEN
        ALTER TABLE company_role_mappings DROP COLUMN user_id;
    END IF;

    IF NOT EXISTS (
        SELECT 1 FROM information_schema.STATISTICS
        WHERE TABLE_SCHEMA = DATABASE()
          AND TABLE_NAME = 'company_role_mappings'
          AND INDEX_NAME = 'idx_company_role'
    ) THEN
        ALTER TABLE company_role_mappings
            ADD UNIQUE KEY idx_company_role (company_id, role_id);
    END IF;
END;
CALL slotly_drop_company_role_user_id();
DROP PROCEDURE IF EXISTS slotly_drop_company_role_user_id;
-- +goose StatementEnd
