INSERT IGNORE INTO `products` (`id`, `name`, `description`, `currency`, `price`, `category`, `type`, `created_at`, `updated_at`) VALUES
(1, 'Spaghetti Carbonara', 'Classic Roman pasta', 'INR', 750, 'MAIN_COURSE', 'NON_VEGETARIAN', NOW(), NOW()),
(2, 'Margherita Pizza', 'Tomato, mozzarella, basil', 'INR', 600, 'MAIN_COURSE', 'VEGETARIAN', NOW(), NOW()),
(3, 'Chicken Tikka Masala', 'Creamy tomato curry', 'INR', 650, 'MAIN_COURSE', 'NON_VEGETARIAN', NOW(), NOW()),
(4, 'Vegetable Samosa', 'Fried pastry with spiced potatoes', 'INR', 45, 'APPETIZER', 'VEGETARIAN', NOW(), NOW()),
(5, 'Garlic Bread', 'Toasted bread with garlic butter', 'INR', 110, 'APPETIZER', 'VEGETARIAN', NOW(), NOW()),
(6, 'Caesar Salad', 'Romaine lettuce, croutons, parmesan', 'INR', 400, 'APPETIZER', 'VEGETARIAN', NOW(), NOW()),
(7, 'Tiramisu', 'Coffee-flavored Italian dessert', 'INR', 180, 'DESSERT', 'VEGETARIAN', NOW(), NOW()),
(8, 'Mango Lassi', 'Yogurt-based mango smoothie', 'INR', 60, 'DRINK', 'VEGETARIAN', NOW(), NOW()),
(9, 'Coca-Cola', 'Classic cola drink', 'INR', 40, 'DRINK', 'VEGAN', NOW(), NOW()),
(10, 'Naan Bread', 'Leavened Indian flatbread', 'INR', 50, 'MAIN_COURSE', 'VEGETARIAN', NOW(), NOW());

-- -- +goose Down
-- DELETE FROM products;
