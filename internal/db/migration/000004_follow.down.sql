ALTER TABLE "follow" DROP CONSTRAINT account_followed_id;
ALTER TABLE "follow" DROP CONSTRAINT account_following_id;
DROP TABLE IF EXISTS "follow";
