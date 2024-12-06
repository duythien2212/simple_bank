CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hash_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_change_at" timestamp NOT NULL DEFAULT '0001-01-01: 00:00:00Z',
  "create_at" timestamp NOT NULL DEFAULT 'now()'
);

-- CREATE UNIQUE INDEX ON "account" ("owner", "currency");

ALTER TABLE "account" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "account" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner", "currency");
