-- +goose Up
CREATE TABLE IF NOT EXISTS user_role_mappings (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    user_id CHAR(36) NOT NULL,
    role_id BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY idx_user_role (user_id, role_id),
    KEY idx_user_role_mappings_deleted_at (deleted_at),
    CONSTRAINT fk_user_role_mappings_user
        FOREIGN KEY (user_id) REFERENCES users (user_id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_user_role_mappings_role
        FOREIGN KEY (role_id) REFERENCES roles (id)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO user_role_mappings (created_at, updated_at, user_id, role_id)
VALUES (
    CURRENT_TIMESTAMP(3),
    CURRENT_TIMESTAMP(3),
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    1
);

-- +goose Down
DROP TABLE IF EXISTS user_role_mappings;
