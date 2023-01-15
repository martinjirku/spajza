-- migrate:up
CREATE TABLE storage_items (
    storage_item_id   INT AUTO_INCREMENT PRIMARY KEY,
    created_at        DATETIME NOT NULL,
    updated_at        DATETIME NOT NULL,
    deleted_at        DATETIME NULL,
    title             VARCHAR(255),
    storage_place_id  INT NULL,
    category_id       INT NULL,
    baseline_amount   DOUBLE NOT NULL,
    current_amount    DOUBLE NOT NULL,
    quantity          ENUM('mass','length','volume','temperature','time','count') NOT NULL,
    unit              VARCHAR(50) NOT NULL,
    expiration_date   DATETIME NULL,


    CONSTRAINT storage_item_storage_place_fk
        FOREIGN KEY (storage_place_id)
        REFERENCES  storage_places(storage_place_id)
        ON DELETE   SET NULL
        ON UPDATE   CASCADE,
    CONSTRAINT storage_item_category_fk
        FOREIGN KEY (category_id)
        REFERENCES  categories(id)
        ON DELETE   SET NULL
        ON UPDATE   CASCADE
);

-- migrate:down
DROP TABLE storage_items;