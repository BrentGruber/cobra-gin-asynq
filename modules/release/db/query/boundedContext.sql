-- name: Createbounded_contexts :one
INSERT INTO bounded_contexts (
    name,
    manager
) VALUES (
    $1,
    $2
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM bounded_contexts
WHERE id = $1 LIMIT 1;

-- name: Listbounded_contexts :many
SELECT * FROM bounded_contexts
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: Deletebounded_contexts :exec
DELETE FROM bounded_contexts
WHERE id = $1;