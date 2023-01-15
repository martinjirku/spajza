-- migrate:up
CREATE TABLE users (
    id                INT AUTO_INCREMENT PRIMARY KEY,
    created_at        DATETIME NOT NULL,
    updated_at        DATETIME NOT NULL,
    deleted_at        DATETIME NULL,
    password          VARCHAR(255),
    email             VARCHAR(255),

    UNIQUE KEY idx_user_email (email)
);

-- migrate:down
DROP TABLE users;