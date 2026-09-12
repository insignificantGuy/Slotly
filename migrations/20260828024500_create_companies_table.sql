-- +goose Up
CREATE TABLE IF NOT EXISTS companies (
    company_id CHAR(36) NOT NULL DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    logo VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    is_active TINYINT(1) NOT NULL DEFAULT 1,
    is_deleted TINYINT(1) NOT NULL DEFAULT 0,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (company_id),
    UNIQUE KEY uni_companies_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS companies;
