-- migrate:up

CREATE TABLE categories (
    id                INT AUTO_INCREMENT PRIMARY KEY,
    created_at        DATETIME NOT NULL,
    updated_at        DATETIME NOT NULL,
    deleted_at        DATETIME NULL,
    title             VARCHAR(250),
    default_unit      VARCHAR(50),
    path              VARCHAR(250)
);

-- migrate:down

DROP TABLE categories;