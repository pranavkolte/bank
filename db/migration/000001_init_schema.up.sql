CREATE TABLE "accounts" (
  "id" BIGSERIAL PRIMARY KEY,
  "owner" VARCHAR NOT NULL,
  "avail_balance" BIGINT NOT NULL,
  "currency" VARCHAR NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "entries" (
  "id" BIGSERIAL PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE TABLE "transfers" (
  "transaction_id" BIGSERIAL PRIMARY KEY,
  "from_acc_id" bigint NOT NULL,
  "to_acc_id" bigint NOT NULL,
  "amount" bigint,
  "created_at" timestamptz DEFAULT 'now()'
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "entries" ("account_id");

CREATE INDEX ON "transfers" ("from_acc_id");

CREATE INDEX ON "transfers" ("from_acc_id");

CREATE INDEX ON "transfers" ("from_acc_id", "to_acc_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("from_acc_id") REFERENCES "accounts" ("id");

ALTER TABLE "transfers" ADD FOREIGN KEY ("to_acc_id") REFERENCES "accounts" ("id");
