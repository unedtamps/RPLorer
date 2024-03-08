CREATE TABLE "account_detail" (
  "account_id" uuid NOT NULL,
  "give_like" bigint NOT NULL DEFAULT 0,
  "get_like" bigint NOT NULL DEFAULT 0,
  "followers_count" bigint NOT NULL DEFAULT 0,
  "following_count" bigint NOT NULL DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE INDEX account_detail_index ON "account_detail" ("account_id");

ALTER TABLE "account_detail" ADD CONSTRAINT account_detail_user FOREIGN KEY ("account_id") REFERENCES "account" ("id");

CREATE TRIGGER account_detail_on_update
BEFORE UPDATE ON account_detail
FOR EACH ROW
EXECUTE FUNCTION on_update();
