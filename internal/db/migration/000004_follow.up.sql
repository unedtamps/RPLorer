CREATE TABLE "follow" (
  "account_followed" uuid NOT NULL,
  "account_following" uuid NOT NULL,
  PRIMARY KEY ("account_followed", "account_following")
);

ALTER TABLE "follow" ADD CONSTRAINT account_followed_id FOREIGN KEY ("account_followed") REFERENCES "account" ("id");

ALTER TABLE "follow" ADD CONSTRAINT account_following_id FOREIGN KEY ("account_following") REFERENCES "account" ("id");
