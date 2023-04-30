

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

-- name: InsertMultipleProductCategories :exec
-- INSERT INTO product_categories (category_id, name, path, parent_id) VALUES 