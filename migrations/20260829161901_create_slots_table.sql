-- +goose Up
CREATE TABLE IF NOT EXISTS slots (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    listing_id CHAR(36) NOT NULL,
    date DATETIME(3) NOT NULL,
    start_time DATETIME(3) NOT NULL,
    end_time DATETIME(3) NOT NULL,
    duration INT NOT NULL,
    is_booked TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (id),
    KEY idx_slots_deleted_at (deleted_at),
    KEY idx_slots_listing_id (listing_id),
    CONSTRAINT fk_slots_listing
        FOREIGN KEY (listing_id) REFERENCES listings (listing_id)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS slots;
