CREATE TABLE "account" (
  "id" uuid PRIMARY KEY NOT NULL,
  "first_name" varchar(512) NOT NULL,
  "last_name" varchar(512) NOT NULL,
  "username" varchar(512) NOT NULL,
  "email" varchar(512) NOT NULL,
  "role" role NOT NULL DEFAULT 'USER',
  "status" status NOT NULL DEFAULT 'OK',
  "password" varchar(512) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TRIGGER account_on_update
BEFORE UPDATE ON account
FOR EACH ROW
EXECUTE FUNCTION on_update();

