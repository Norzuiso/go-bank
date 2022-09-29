-- name: CreateTransfers :one
INSERT into transfers(
    from_account_id,
    to_account_id,
    amount
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 limit 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfers :exec
UPDATE transfers
set amount =$2
WHERE id = $1;

-- name: DeleteTransfers :exec
DELETE FROM transfers WHERE id = $1;