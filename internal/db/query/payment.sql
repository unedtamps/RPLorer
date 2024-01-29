-- name: CreatePayment :one
INSERT INTO "Payment" (id, user_id, premium_type_id, amount) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetPaymentByUserId :one
SELECT * FROM "Payment" WHERE user_id = $1 LIMIT 1;

-- name: GetPaymentById :one
SELECT * FROM "Payment" WHERE id = $1 LIMIT 1;

-- name: UpdatePaymentStatus :one
UPDATE "Payment" SET status = $1 WHERE id = $2 RETURNING *;

-- name: GetPremiumTypeName :one
SELECT * FROM "PremiumType" WHERE type_name = $1 LIMIT 1;
