-- name: CreateAccount :one
INSERT INTO accounts (
  owner, 
  balance, 
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

-- -- name: GetAuthor :one
-- SELECT * FROM entries
-- WHERE id = ? LIMIT 1;

-- -- name: ListAuthors :many
-- SELECT * FROM entries
-- ORDER BY name;

-- -- name: CreateAuthor :execresult
-- INSERT INTO entries (
--   name, bio
-- ) VALUES (
--   ?, ?
-- );

-- -- name: DeleteAuthor :exec
-- DELETE FROM authors
-- WHERE id = ?;