-- +goose Up
-- +goose StatementBegin
CREATE TABLE coupons (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `filename` VARCHAR(255) NOT NULL,
    `code` VARCHAR(64) NOT NULL,
    `is_active` BOOLEAN NOT NULL DEFAULT TRUE,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_coupons_code` (`code`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS coupons;
-- +goose StatementEnd 