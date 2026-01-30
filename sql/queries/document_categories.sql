-- name: CreateDocumentCategory :one 
INSERT INTO document_categories(id, name, created_at, updated_at, active)
VALUES (
  gen_random_UUID(),
  $1, 
  NOW(),
  NOW(),
  true
  )
RETURNING *;

-- name: GetDocumentCategory :one
SELECT id, name, active, created_at, updated_at from document_categories where id = $1;

-- name: GetDocumentCategories :many
SELECT id, name, active, created_at, updated_at from document_categories;

-- name: UpdateDocumentCategory :one
UPDATE document_categories SET name = $2, updated_at = NOW(), active = $3 where id = $1 RETURNING *;

-- name: DeleteDocumentCategory :exec
DELETE FROM document_categories WHERE id = $1;
