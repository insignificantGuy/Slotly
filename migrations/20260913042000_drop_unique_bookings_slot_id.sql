-- +goose Up
ALTER TABLE bookings DROP INDEX uni_bookings_slot_id;
ALTER TABLE bookings ADD KEY idx_bookings_slot_id (slot_id);

-- +goose Down
ALTER TABLE bookings DROP INDEX idx_bookings_slot_id;
ALTER TABLE bookings ADD UNIQUE KEY uni_bookings_slot_id (slot_id);
