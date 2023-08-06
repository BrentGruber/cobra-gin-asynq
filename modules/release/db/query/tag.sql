-- name: Listtags :many
SELECT * FROM tags
ORDER BY tag
LIMIT $1
OFFSET $2;