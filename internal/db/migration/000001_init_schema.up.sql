
CREATE TYPE payment_status AS ENUM ('pending', 'success', 'failed');

CREATE TABLE "User" (
  "id" varchar(36) PRIMARY KEY NOT NULL,
  "name" varchar(512) NOT NULL,
  "email" varchar(512) UNIQUE NOT NULL,
  "password" varchar(512) NOT NULL,
  "account_status" boolean NOT NULL DEFAULT false,
  "acount_type" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE "Todo" (
  "id" varchar(36) PRIMARY KEY,
  "user_id" varchar(36) NOT NULL,
  "todo_name" varchar(512) NOT NULL,
  "status" boolean DEFAULT false NOT NULL,
  "description" text NOT NULL,
  "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
  "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE "Payment" (
  "id" varchar(36) PRIMARY KEY,
  "user_id" varchar(36) NOT NULL,
  "status" payment_status NOT NULL DEFAULT 'pending',
  "premium_type_id" varchar(36) NOT NULL,
  "amount" bigint NOT NULL
);

CREATE INDEX ON "User" ("email");

CREATE INDEX ON "Todo" ("user_id");

ALTER TABLE "Payment" ADD CONSTRAINT "user_payment" FOREIGN KEY ("user_id") REFERENCES "User" ("id");

ALTER TABLE "Todo" ADD CONSTRAINT "user_todo" FOREIGN KEY ("user_id") REFERENCES "User" ("id");

CREATE TABLE "PremiumType" (
  "id" varchar(36) PRIMARY KEY,
  "type_name" varchar(512) NOT NULL
);

CREATE TABLE "Quota" (
  "id" varchar(36) PRIMARY KEY,
  "premium_type_id" varchar(36),
  "quota_amount" bigint NOT NULL CHECK (quota_amount >= 0)
);

ALTER TABLE "Quota" ADD CONSTRAINT "premium_type_quota" 
FOREIGN KEY ("premium_type_id") 
REFERENCES "PremiumType" ("id") ON DELETE CASCADE;

ALTER TABLE "Payment" ADD CONSTRAINT "premium_type_payment" FOREIGN KEY ("premium_type_id") REFERENCES "PremiumType" ("id");

INSERT INTO "PremiumType" (id, type_name) VALUES ('dcaa02fb-e960-40db-b38b-75dc104fb017', 'premium_rare');
INSERT INTO "PremiumType" (id, type_name) VALUES ('59318f8a-3e06-49e7-86f7-006039fb2112', 'premium_epic');

INSERT INTO "Quota" (id, premium_type_id, quota_amount) VALUES ('ce5e767d-441c-4736-95f5-02267d8e7129', 'dcaa02fb-e960-40db-b38b-75dc104fb017', 10);
INSERT INTO "Quota" (id, premium_type_id, quota_amount) VALUES ('6719c9b4-b5f0-4289-8f9b-ed85d4610078', '59318f8a-3e06-49e7-86f7-006039fb2112', 5);
