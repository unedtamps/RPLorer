CREATE TYPE "role" AS ENUM (
  'USER',
  'ADMIN'
);

CREATE TYPE "status" AS ENUM (
  'OK',
  'DELETED'
);

CREATE TYPE "account_status" AS ENUM (
  'ACTIVE',
  'INACTIVE'
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION on_update()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';
