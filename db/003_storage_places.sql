-- migrate:up

CREATE TABLE storage_places (
    storage_place_id INT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    title VARCHAR(255),
    code VARCHAR(60) UNIQUE
);

-- migrate:down
DROP TABLE storage_places;