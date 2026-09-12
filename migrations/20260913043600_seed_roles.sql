-- +goose Up
INSERT IGNORE INTO roles (id, created_at, updated_at, name)
VALUES
    (1, CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3), 'super_admin'),
    (2, CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3), 'org_admin'),
    (3, CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3), 'staff'),
    (4, CURRENT_TIMESTAMP(3), CURRENT_TIMESTAMP(3), 'customer');

-- +goose Down
DELETE FROM roles
WHERE id IN (1, 2, 3, 4)
  AND name IN ('super_admin', 'org_admin', 'staff', 'customer');
