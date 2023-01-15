-- migrate:up

CREATE TABLE storage_consumptions (
    storage_item_consumption_id INT AUTO_INCREMENT PRIMARY KEY,
    created_at                  DATETIME NOT NULL,
    updated_at                  DATETIME NOT NULL,
    deleted_at                  DATETIME NULL,
    normalized_amount           DOUBLE,
    unit                        VARCHAR(50),
    storage_item_id             INT NOT NULL,

    CONSTRAINT storage_consumption_storage_item_fk
        FOREIGN KEY (storage_item_id)
        REFERENCES  storage_items(storage_item_id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
)

-- migrate:down
DROP TABLE storage_consumptions;