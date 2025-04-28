INSERT IGNORE INTO `users` (`id`, `name`, `email`, `phone_number`, `country_code`, `created_at`, `updated_at`) VALUES
(1, 'Alice Admin', 'alice@example.com', '8888999900', '+91', NOW(), NOW()),
(2, 'Bob User', 'bob@example.com', '8888999911', '+91', NOW(), NOW());


-- -- +goose Down
-- DELETE FROM users;
