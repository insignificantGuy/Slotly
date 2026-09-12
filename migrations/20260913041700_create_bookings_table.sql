-- +goose Up
CREATE TABLE IF NOT EXISTS bookings (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at DATETIME(3),
    updated_at DATETIME(3),
    deleted_at DATETIME(3),
    user_id CHAR(36) NOT NULL,
    slot_id BIGINT UNSIGNED NOT NULL,
    listing_id CHAR(36) NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'confirmed',
    PRIMARY KEY (id),
    UNIQUE KEY uni_bookings_slot_id (slot_id),
    KEY idx_bookings_deleted_at (deleted_at),
    KEY idx_bookings_user_id (user_id),
    KEY idx_bookings_listing_id (listing_id),
    CONSTRAINT fk_bookings_user
        FOREIGN KEY (user_id) REFERENCES users (user_id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_bookings_slot
        FOREIGN KEY (slot_id) REFERENCES slots (id)
        ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT fk_bookings_listing
        FOREIGN KEY (listing_id) REFERENCES listings (listing_id)
        ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +goose Down
DROP TABLE IF EXISTS bookings;
