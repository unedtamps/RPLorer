DROP TRIGGER IF EXISTS post_on_update ON "post";
ALTER TABLE "post" DROP CONSTRAINT IF EXISTS post_account_id;
DROP TABLE IF EXISTS "post";
