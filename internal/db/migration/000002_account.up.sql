CREATE TABLE "account" (
  "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
  "first_name" varchar(512) NOT NULL,
  "last_name" varchar(512) NOT NULL,
  "username" varchar(512) NOT NULL UNIQUE,
  "email" varchar(512) NOT NULL UNIQUE,
  "role" role NOT NULL DEFAULT 'USER',
  "status" account_status NOT NULL DEFAULT 'INACTIVE',
  "password" varchar(512) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT 'now()',
  "updated_at" timestamp NOT NULL DEFAULT 'now()'
);

CREATE TRIGGER account_on_update
BEFORE UPDATE ON account
FOR EACH ROW
EXECUTE FUNCTION on_update();

