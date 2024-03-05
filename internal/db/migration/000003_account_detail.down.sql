DROP TRIGGER IF EXISTS account_detail_on_update ON "account_detail";
DROP INDEX IF EXISTS account_detail_index;
ALTER TABLE "account_detail" DROP CONSTRAINT IF EXISTS account_detail_user;
DROP TABLE IF EXISTS "account_detail";
