-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `uuid` VARCHAR(255) NOT NULL,
    `user_id` BIGINT NOT NULL,
    `restaurant_id` BIGINT NOT NULL,
    `status` enum('PENDING','CONFIRMED','PREPARING','DELIVERED','CANCELLED') NOT NULL DEFAULT 'PENDING',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uuid` (`uuid`),
    KEY `restaurant_id` (`restaurant_id`),
    CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`restaurant_id`) REFERENCES `restaurants` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    KEY `user_id` (`user_id`),
    CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd 