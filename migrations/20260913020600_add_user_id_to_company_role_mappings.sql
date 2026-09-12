-- +goose Up
ALTER TABLE company_role_mappings
    ADD COLUMN user_id CHAR(36) NULL AFTER deleted_at;

UPDATE company_role_mappings
SET user_id = 'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa'
WHERE user_id IS NULL;

ALTER TABLE company_role_mappings
    DROP INDEX idx_company_role,
    MODIFY user_id CHAR(36) NOT NULL,
    ADD UNIQUE KEY idx_user_company (user_id, company_id),
    ADD CONSTRAINT fk_company_role_mappings_user
        FOREIGN KEY (user_id) REFERENCES users (user_id)
        ON UPDATE CASCADE ON DELETE CASCADE;

-- +goose Down
ALTER TABLE company_role_mappings
    DROP FOREIGN KEY fk_company_role_mappings_user,
    DROP INDEX idx_user_company,
    DROP COLUMN user_id,
    ADD UNIQUE KEY idx_company_role (company_id, role_id);
