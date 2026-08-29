-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id CHAR(36) NOT NULL DEFAULT (UUID()),
    full_name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    phone_number VARCHAR(20),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (user_id),
    UNIQUE KEY uni_users_email (email),
    UNIQUE KEY uni_users_phone_number (phone_number)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Seed user referenced by user_role_mappings.
INSERT INTO users (user_id, full_name, email, password, phone_number, created_at, updated_at)
VALUES (
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    'Test User',
    'test@test.com',
    'password',
    '1234567890',
    CURRENT_TIMESTAMP(3),
    CURRENT_TIMESTAMP(3)
);

-- +goose Down
DROP TABLE IF EXISTS users;
