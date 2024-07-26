-- Create table
CREATE TABLE users (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

CREATE TABLE orders (
    id INT PRIMARY KEY,
    user_id INT,
    amount DECIMAL(10,2),
    created_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE order_items (
    id INT PRIMARY KEY,
    order_id INT,
    product_name VARCHAR(255),
    price DECIMAL(10,2),
    quantity INT,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);


-- Insert data
INSERT INTO users (id, name, email)
VALUES
    (1, 'John Doe', 'johndoe@example.com'), 
    (2, 'Jane Smith', 'janesmith@example.com'), 
    (3, 'Bob Johnson', 'bobjohnson@example.com');

INSERT INTO orders (id, user_id, amount, created_at) VALUES 
(1, 1, 100.00, '2022-01-02 10:30:00'), 
(2, 2, 50.00, '2022-01-03 09:00:00'), 
(3, 1, 150.00, '2022-01-04 14:15:00'), 
(4, 3, 200.00, '2022-01-05 17:45:00'), 
(5, 2, 75.00, '2022-01-06 11:20:00');

--testing with value >= 1000
INSERT INTO orders (id, user_id, amount, created_at) VALUES (6, 2, 1000, '2022-01-06 11:20:00')

INSERT INTO order_items (id, order_id, product_name, price, quantity) VALUES 
(1, 1, 'T-Shirt', 25.00, 2), 
(2, 1, 'Jeans', 50.00, 1), 
(3, 2, 'Socks', 10.00, 5), 
(4, 3, 'Shoes', 75.00, 2), 
(5, 4, 'Jacket', 100.00, 1), 
(6, 5, 'Sweater', 25.00, 3);

-- answer = Jane Smith 1125.00
SELECT u.name,
       SUM(o.amount) AS total_spent
FROM users u
JOIN orders o ON u.id = o.user_id
WHERE o.created_at >= '2022-01-01'
GROUP BY u.name
HAVING SUM(o.amount) >= 1000;
