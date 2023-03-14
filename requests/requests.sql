# One-To-Many relationship

CREATE TABLE IF NOT EXISTS orders
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    customer_id INT NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders_products
(
    order_id   INT NOT NULL,
    product_id INT NOT NULL,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

# Get all sellers and amount of products they sell

SELECT sellers.id,
       sellers.name,
       sellers.phone_number,
       (SELECT COUNT(products.id) FROM products WHERE seller_id = sellers.id) as products_count
FROM sellers;

SELECT sellers.id,
       sellers.name,
       sellers.phone_number,
       COUNT(p.id) as count
FROM sellers
         LEFT JOIN products p on sellers.id = p.seller_id
GROUP BY sellers.id
LIMIT 1;


# Get all customers and each customer total expenses

SELECT customers.id,
       customers.name,
       customers.phone_number,
       (SELECT SUM(p.price)
        FROM orders_products
                 LEFT JOIN products p on p.id = orders_products.product_id
        WHERE order_id IN (SELECT id FROM orders WHERE customer_id = customers.id)) as total_expense
FROM customers;

SELECT customers.id,
       customers.name,
       customers.phone_number,
       SUM(p.price) as total
FROM customers
         LEFT JOIN orders o on customers.id = o.customer_id
         LEFT JOIN orders_products op on o.id = op.order_id
         LEFT JOIN products p on op.product_id = p.id
GROUP BY customers.id
HAVING total > 100
ORDER BY total DESC;





