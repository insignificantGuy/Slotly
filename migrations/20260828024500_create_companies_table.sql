-- +goose Up
CREATE TABLE IF NOT EXISTS companies (
    company_id CHAR(36) NOT NULL DEFAULT (UUID()),
    password VARCHAR(255),
    email VARCHAR(255),
    created_at DATETIME(3),
    updated_at DATETIME(3),
    PRIMARY KEY (company_id),
    UNIQUE KEY uni_companies_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Seed company referenced by company_role_mappings and listings.
INSERT INTO companies (company_id, password, email, created_at, updated_at)
VALUES (
    'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
    'password',
    'company@test.com',
    CURRENT_TIMESTAMP(3),
    CURRENT_TIMESTAMP(3)
);

-- +goose Down
DROP TABLE IF EXISTS companies;
