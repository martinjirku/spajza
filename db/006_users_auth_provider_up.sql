ALTER TABLE users
    ADD COLUMN name             VARCHAR(255),
    ADD COLUMN given_name       VARCHAR(255) NULL,
    ADD COLUMN family_name      VARCHAR(255) NULL,
    ADD COLUMN picture          VARCHAR(255) NULL,
    ADD COLUMN email_verified   TINYINT(1)   NOT NULL DEFAULT 0,
    ADD COLUMN auth_provider    INT          NOT NULL DEFAULT 0;