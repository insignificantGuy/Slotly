-- +goose Up
CREATE TABLE IF NOT EXISTS listings (
    listing_id CHAR(36) NOT NULL DEFAULT (UUID()),
    type VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    company_id CHAR(36) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    created_at DATETIME(3) NOT NULL,
    updated_at DATETIME(3) NOT NULL,
    deleted_at DATETIME(3),
    PRIMARY KEY (listing_id),
    KEY idx_listings_deleted_at (deleted_at),
    KEY idx_listings_company_id (company_id),
    CONSTRAINT fk_listings_company
        FOREIGN KEY (company_id) REFERENCES companies (company_id)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS listings;
