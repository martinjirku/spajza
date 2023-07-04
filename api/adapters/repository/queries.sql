

-- CATEGORIES:

-- name: ListCategories :many
SELECT * FROM categories WHERE deleted_at IS null;

-- name: CreateCategory :execlastid
INSERT INTO categories(created_at, updated_at, title, path, default_unit) VALUES (NOW(), NOW(),?,?,?);

-- name: UpdateCategory :exec
UPDATE categories SET updated_at=NOW(),title=?,path=?,default_unit=? WHERE id=?;

-- name: DeleteCategory :exec
UPDATE categories SET updated_at=NOW(), deleted_at=NOW() WHERE id=?;


-- PRODUCT CATEGORIES:

-- name: CreateProductCategory :execlastid
INSERT INTO product_categories (category_id, name, path, parent_id) VALUES (?,?,?,?);

-- STORAGE ITEMS:

-- name: CreateStorageItem :execlastid
INSERT INTO storage_items (created_at, updated_at, title,
    storage_place_id, category_id, baseline_amount, current_amount,
    quantity, unit, expiration_date, ean)
VALUES (NOW(),NOW(),?,?,?,?,?,?,?,?,?);

-- name: UpdateStorageItem :exec
UPDATE storage_items SET updated_at=NOW(), title=?, storage_place_id=?,
    category_id=?, baseline_amount=?, unit=?,
    expiration_date=?, ean=?
WHERE storage_item_id=?;

-- name: ListStorageItems :many
SELECT * FROM storage_items WHERE deleted_at IS NULL LIMIT ? OFFSET ?;

-- name: GetStorageItemById :one
SELECT * FROM storage_items WHERE storage_item_id=?;

-- STORAGE CONSUMPTIONS:

-- name: CreateStorageConsumption :execlastid
INSERT INTO storage_consumptions (created_at, updated_at, normalized_amount,
    unit, storage_item_id)
VALUES (NOW(),NOW(),?,?,?);

-- name: ListStorageConsumptions :many
SELECT * FROM storage_consumptions
WHERE storage_item_id IN (SELECT storage_item_id FROM storage_items);

-- name: GetStorageConsumptionById :many
SELECT * FROM storage_consumptions WHERE storage_item_id=?;

-- STORAGE PLACE:

-- name: CreateStoragePlace :execlastid
INSERT INTO storage_places(created_at, updated_at, title, code)
VALUES (NOW(),NOW(),?,?);

-- name: GetStoragePlaceById :one
SELECT * FROM storage_places
WHERE deleted_at IS NULL && storage_place_id=?;

-- name: ListStoragePlaces :many
SELECT * FROM storage_places
WHERE deleted_at IS NULL;

-- name: UpdateStoragePlace :exec
UPDATE storage_places
SET updated_at=NOW(),title=?,code=?
WHERE storage_place_id=?;

-- name: DeleteStoragePlace :exec
UPDATE storage_places SET deleted_at=NOW()
WHERE storage_place_id=?