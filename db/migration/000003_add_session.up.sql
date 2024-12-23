CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar UNIQUE NOT NULL,
  "is_block" boolean NOT NULL DEFAULT false,
  "expired_at" timestamp NOT NULL,
  "create_at" timestamp NOT NULL DEFAULT 'now()'
);

-- CREATE UNIQUE INDEX ON "account" ("owner", "currency");

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");
