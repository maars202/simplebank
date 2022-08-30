-- deadlock can occur in this scenario -- solution: reverse order of 
-- execution -- ensure that all smaller ids are updated first before larger ids:
-- TX1
BEGIN;

UPDATE accounts SET balance = balance - 10 WHERE id = 1 RETURNING *;
UPDATE accounts SET balance = balance + 10 WHERE id = 2 RETURNING *;

ROLLBACK;

-- TX2
BEGIN;

UPDATE accounts SET balance = balance + 10 WHERE id = 1 RETURNING *;
UPDATE accounts SET balance = balance - 10 WHERE id = 2 RETURNING *;
-- UPDATE accounts SET balance = balance + 10 WHERE id = 1 RETURNING *;

ROLLBACK;