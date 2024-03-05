DROP TRIGGER IF EXISTS comment_on_update ON "comment"; 
ALTER TABLE "comment" DROP CONSTRAINT IF EXISTS comment_post_id;
ALTER TABLE "comment" DROP CONSTRAINT IF EXISTS comment_account_id;
ALTER TABLE "comment" DROP CONSTRAINT IF EXISTS comment_parent_id;
DROP TABLE IF EXISTS "comment";
