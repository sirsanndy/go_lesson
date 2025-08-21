-- Create customer table
CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    user_name VARCHAR(64) UNIQUE NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    name VARCHAR(128) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Composite index for user_name and email
CREATE INDEX idx_customer_user_email ON customer(user_name, email);

-- Create order table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    transaction_code VARCHAR(64) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id INTEGER NOT NULL REFERENCES customer(id),
    state VARCHAR(32) NOT NULL
);

-- Index for transaction_code
CREATE INDEX idx_order_transaction_code ON orders(transaction_code);

-- Generate 1 million random customers
INSERT INTO customer (user_name, email, name, created_at, updated_at, is_active)
SELECT 
    'user_' || g, 
    'user_' || g || '@go-lesson.com', 
    'Customer ' || g, 
    NOW() - (g % 365) * INTERVAL '1 day', 
    NOW(), 
    (g % 2 = 0)
FROM generate_series(1, 1000000) AS g;

-- Generate 1 million random orders
INSERT INTO orders (transaction_code, created_at, updated_at, user_id, state)
SELECT 
    'TXN' || g,
    NOW() - (g % 365) * INTERVAL '1 day',
    NOW(),
    (random() * 999999 + 1)::integer,
    CASE WHEN g % 5 = 0 THEN 'pending' WHEN g % 5 = 1 THEN 'completed' WHEN g % 5 = 2 THEN 'cancelled' WHEN g % 5 = 3 THEN 'failed' ELSE 'refunded' END
FROM generate_series(1, 1000000) AS g;
