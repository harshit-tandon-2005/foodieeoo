INSERT INTO `restaurant_products` (`restaurant_id`, `product_id`, `is_available`, `created_at`, `updated_at`) VALUES
(1, 1, TRUE, NOW(), NOW()),
(1, 2, TRUE, NOW(), NOW()),
(1, 3, TRUE, NOW(), NOW()),
(1, 4, TRUE, NOW(), NOW()),
(1, 5, TRUE, NOW(), NOW()),
(2, 6, TRUE, NOW(), NOW()),
(2, 7, TRUE, NOW(), NOW()),
(2, 8, TRUE, NOW(), NOW()),
(2, 9, TRUE, NOW(), NOW()),
(2, 10, TRUE, NOW(), NOW());

-- -- +goose Down
-- DELETE FROM restaurant_products;