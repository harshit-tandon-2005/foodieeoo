INSERT IGNORE INTO `restaurants` (`id`, `name`, `address`, `phone_number`, `country_code`, `created_at`, `updated_at`) VALUES
(1, 'Pasta Paradise', '123 Pasta Lane, New Delhi', '8888999922', '+91', NOW(), NOW()),
(2, 'Curry Corner', '456 Spice Street, New Delhi', '8888999933', '+91', NOW(), NOW());

-- -- +goose Down
-- DELETE FROM restaurants;