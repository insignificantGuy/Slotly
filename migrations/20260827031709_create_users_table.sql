-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id CHAR(36) NOT NULL DEFAULT (UUID()),
    full_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    phone_number VARCHAR(20),
    is_deleted TINYINT(1) NOT NULL DEFAULT 0,
    deleted_at DATETIME(3),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (user_id),
    UNIQUE KEY uni_users_email (email),
    UNIQUE KEY uni_users_phone_number (phone_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS users;
