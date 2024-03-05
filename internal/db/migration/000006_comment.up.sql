CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY NOT NULL,
  "post_id" uuid NOT NULL,
  "account_id" uuid NOT NULL,
  "parrent_id" uuid,
  "body" text NOT NULL,
  "status" status NOT NULL DEFAULT 'OK',
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

ALTER TABLE "comment" ADD CONSTRAINT comment_account_id  FOREIGN KEY ("account_id") REFERENCES "account" ("id");

ALTER TABLE "comment" ADD CONSTRAINT comment_post_id FOREIGN KEY ("post_id") REFERENCES "post" ("id");

ALTER TABLE "comment" ADD CONSTRAINT comment_parrent_id FOREIGN KEY ("parrent_id") REFERENCES "comment" ("id");

CREATE TRIGGER comment_on_update
BEFORE UPDATE ON comment
FOR EACH ROW
EXECUTE FUNCTION on_update();
