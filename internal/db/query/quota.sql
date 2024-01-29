-- name: DecreaseOneQuota :exec
UPDATE "Quota" SET quota_amount = quota_amount - 1 WHERE id = $1;

-- name: GetQuotaByPremiumTypeId :one
SELECT * FROM "Quota" WHERE premium_type_id = $1 LIMIT 1;

-- name: GetQuotaByPremiumTypeIdForUpdate :one
SELECT * FROM "Quota" WHERE premium_type_id = $1 LIMIT 1 FOR NO KEY UPDATE;
