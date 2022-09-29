-- name: CreateEntries :one
INSERT into entries(
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 limit 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEntries :exec
UPDATE entries
set amount =$2
WHERE id = $1;

-- name: DeleteEntries :exec
DELETE FROM entries WHERE id = $1;