-- Create table products with detailed comments
CREATE TABLE products (
    id SERIAL PRIMARY KEY,                           -- Auto-increment primary key
    name VARCHAR(255) NOT NULL,                      -- Product name, not null
    category VARCHAR(255) NOT NULL,                  -- Product category, not null
    price NUMERIC(10, 2) NOT NULL,                   -- Product price with 2 decimal precision, not null
    description TEXT,                                -- Product description, nullable
    stock_quantity INT NOT NULL DEFAULT 0,           -- Stock quantity, default value is 0
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),   -- Creation timestamp, default value is current time
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),   -- Update timestamp, default value is current time
    deleted_at TIMESTAMPTZ                           -- Deletion timestamp, nullable, for soft delete
);

-- Add comments to the table and columns
COMMENT ON TABLE products IS 'Table storing product information';

COMMENT ON COLUMN products.id IS 'Product ID, auto-increment primary key';
COMMENT ON COLUMN products.name IS 'Product name, not null';
COMMENT ON COLUMN products.category IS 'Product category, not null';
COMMENT ON COLUMN products.price IS 'Product price with 2 decimal precision, not null';
COMMENT ON COLUMN products.description IS 'Product description, nullable';
COMMENT ON COLUMN products.stock_quantity IS 'Stock quantity, default value is 0';
COMMENT ON COLUMN products.created_at IS 'Creation timestamp, default value is current time';
COMMENT ON COLUMN products.updated_at IS 'Update timestamp, default value is current time';
COMMENT ON COLUMN products.deleted_at IS 'Deletion timestamp, nullable, for soft delete';

-- Create indexes to optimize queries, with detailed comments
CREATE INDEX idx_products_name ON products(name);
COMMENT ON INDEX idx_products_name IS 'Index created on product name';

CREATE INDEX idx_products_category ON products(category);
COMMENT ON INDEX idx_products_category IS 'Index created on product category';

CREATE INDEX idx_products_price ON products USING BRIN(price);
COMMENT ON INDEX idx_products_price IS 'BRIN index created on product price';

CREATE INDEX idx_products_created_at ON products(created_at);
COMMENT ON INDEX idx_products_created_at IS 'Index created on creation timestamp';

-- Insert sample data into products table
INSERT INTO products (name, category, price, description, stock_quantity, created_at, updated_at)
VALUES
('Wireless Mouse', 'Electronics', 25.99, 'Ergonomic wireless mouse with long battery life', 150, NOW(), NOW()),
('Bluetooth Speaker', 'Electronics', 49.99, 'Portable Bluetooth speaker with high-quality sound', 85, NOW(), NOW()),
('Yoga Mat', 'Sports & Outdoors', 19.99, 'Non-slip yoga mat for all types of yoga and exercise', 200, NOW(), NOW()),
('Stainless Steel Water Bottle', 'Kitchen & Dining', 15.50, 'Double-walled stainless steel water bottle', 300, NOW(), NOW()),
('Running Shoes', 'Sports & Outdoors', 79.99, 'Lightweight running shoes for all terrains', 120, NOW(), NOW()),
('LED Desk Lamp', 'Home & Office', 29.99, 'Adjustable LED desk lamp with multiple brightness levels', 75, NOW(), NOW());
