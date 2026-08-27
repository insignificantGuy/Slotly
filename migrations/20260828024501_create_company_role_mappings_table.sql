-- +goose Up
CREATE TABLE IF NOT EXISTS company_role_mappings (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    company_id CHAR(36) NOT NULL,
    role_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_company_role (company_id, role_id),
    KEY idx_company_role_mappings_deleted_at (deleted_at),
    CONSTRAINT fk_company_role_mappings_company
        FOREIGN KEY (company_id) REFERENCES companies (company_id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_company_role_mappings_role
        FOREIGN KEY (role_id) REFERENCES roles (id)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS company_role_mappings;
