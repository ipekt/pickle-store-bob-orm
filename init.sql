CREATE TABLE products (
    product_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INTEGER NOT NULL DEFAULT 0
);
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE orders (
    order_id SERIAL PRIMARY KEY,
    customer_id INTEGER REFERENCES customers(customer_id),
    product_id INTEGER REFERENCES products(product_id),
    quantity INTEGER NOT NULL,
    order_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO products (name, description, price, stock_quantity)
VALUES 
('Dill Pickles', 'Crispy dill pickles', 4.99, 100),
('Sweet Pickles', 'Sweet and tangy pickles', 5.49, 50),
('Bread and Butter Pickles', 'Classic bread and butter pickles', 5.99, 75),
('Spicy Pickles', 'Pickles with a spicy kick', 6.49, 30),
('Garlic Pickles', 'Pickles infused with garlic', 4.99, 60);

INSERT INTO customers (first_name, last_name, email)
VALUES 
('John', 'Doe', 'john.doe@example.com'),
('Jane', 'Smith', 'jane.smith@example.com'),
('Alice', 'Johnson', 'alice.johnson@example.com'),
('Bob', 'Brown', 'bob.brown@example.com'),
('Charlie', 'Davis', 'charlie.davis@example.com');

INSERT INTO orders (customer_id, product_id, quantity)
VALUES 
(1, 1, 2),  -- John Doe orders 2 Dill Pickles
(2, 2, 1),  -- Jane Smith orders 1 Sweet Pickles
(3, 3, 4),  -- Alice Johnson orders 4 Bread and Butter Pickles
(4, 4, 1),  -- Bob Brown orders 1 Spicy Pickles
(5, 5, 3),  -- Charlie Davis orders 3 Garlic Pickles
(1, 3, 1),  -- John Doe orders 1 Bread and Butter Pickles
(2, 4, 2),  -- Jane Smith orders 2 Spicy Pickles
(3, 1, 5);  -- Alice Johnson orders 5 Dill Pickles
