-- migrate:up

ALTER TABLE storage_items
    ADD COLUMN ean    VARCHAR(50);

-- migrate:down
ALTER TABLE storage_items
    DROP COLUMN ean;