CREATE TABLE "post" (
  "id" uuid PRIMARY KEY NOT NULL,
  "account_id" uuid NOT NULL,
  "caption" text NOT NULL,
  "likes_count" bigint NOT NULL DEFAULT 0,
  "comment_count" bigint NOT NULL DEFAULT 0,
  "status" status NOT NULL DEFAULT 'OK',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "post" ADD CONSTRAINT post_account_id  FOREIGN KEY ("account_id") REFERENCES "account" ("id");

CREATE TRIGGER post_on_update
BEFORE UPDATE ON post
FOR EACH ROW
EXECUTE FUNCTION on_update();
