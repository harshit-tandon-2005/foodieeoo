-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `code` VARCHAR(255),
    `name` VARCHAR(255) NOT NULL,
    `description` TEXT,
    `price` DECIMAL(10, 2) NOT NULL,
    `category` enum('APPETIZER','MAIN_COURSE','DESSERT','DRINK') NOT NULL DEFAULT 'APPETIZER',
    `type` enum('VEGETARIAN','NON_VEGETARIAN','VEGAN') NOT NULL DEFAULT 'VEGETARIAN',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd 