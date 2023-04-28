-- migrate:up
CREATE TABLE product_categories(
    category_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    path VARCHAR(512),
    parent_id INT,
    CONSTRAINT product_categories_parent_id
        FOREIGN KEY (parent_id)
        REFERENCES product_categories (category_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

-- migrate:down
DROP TABLE product_categories
