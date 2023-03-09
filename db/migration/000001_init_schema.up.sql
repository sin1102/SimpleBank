CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "balance" bigint NOT NULL,
  "currency" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entry" (
  "id" bigserial PRIMARY KEY,
  "accounts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "transfer" (
  "id" bigserial PRIMARY KEY,
  "from_accounts_id" bigint NOT NULL,
  "to_accounts_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entry" ("accounts_id");

CREATE INDEX ON "transfer" ("to_accounts_id");

CREATE INDEX ON "transfer" ("from_accounts_id");

CREATE INDEX ON "transfer" ("to_accounts_id", "from_accounts_id");

COMMENT ON COLUMN "entry"."amount" IS 'can be positive or negative';

COMMENT ON COLUMN "transfer"."amount" IS 'must be positive';

ALTER TABLE "entry" ADD FOREIGN KEY ("accounts_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("from_accounts_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfer" ADD FOREIGN KEY ("to_accounts_id") REFERENCES "accounts" ("id");
