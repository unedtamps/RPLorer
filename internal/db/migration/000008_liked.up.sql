CREATE TABLE "liked" (
  "account_id" uuid NOT NULL,
  "post_id" uuid NOT NULL,
  PRIMARY KEY ("account_id", "post_id")
);

ALTER TABLE "liked" ADD CONSTRAINT post_liked FOREIGN KEY ("post_id") REFERENCES "post" ("id");

ALTER TABLE "liked" ADD CONSTRAINT account_liked FOREIGN KEY ("account_id") REFERENCES "account" ("id");
