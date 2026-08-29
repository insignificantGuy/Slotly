-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    name VARCHAR(255),
    actions VARCHAR(255),
    PRIMARY KEY (id),
    UNIQUE KEY uni_roles_name (name),
    KEY idx_roles_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Seed role id=1 referenced by user_role_mappings and company_role_mappings.
INSERT INTO roles (id, created_at, updated_at, name, actions)
VALUES (1, CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3), 'admin', 'all');

-- +goose Down
DROP TABLE IF EXISTS roles;
