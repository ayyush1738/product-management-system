-- Create users table
ALTER USER postgres PASSWORD '1738';

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    product_name VARCHAR(255) NOT NULL,
    product_description TEXT,
    product_images TEXT[] NOT NULL,
    compressed_product_images TEXT[],
    product_price DECIMAL NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Optional: Add indexes for performance
CREATE INDEX IF NOT EXISTS idx_user_id ON products(user_id);
CREATE INDEX IF NOT EXISTS idx_product_name ON products(product_name);
