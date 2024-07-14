CREATE TABLE "Account" (
  "ID" bigserial PRIMARY KEY,
  "Owner" varchar NOT NULL,
  "Balance" bigint NOT NULL,
  "Currency" varchar NOT NULL,
  "Createdat" timestamp DEFAULT (now())
);

CREATE TABLE "Entries" (
  "ID" bigserial PRIMARY KEY,
  "Account_id" bigint NOT NULL,
  "Amount" bigint NOT NULL,
  "Createdat" timestamp DEFAULT (now())
);

CREATE TABLE "Transac" (
  "ID" bigserial PRIMARY KEY,
  "From_account_id" bigint NOT NULL,
  "To_account_id" bigint NOT NULL,
  "Amount" bigint NOT NULL,
  "Createdat" timestamp DEFAULT (now())
);

CREATE INDEX ON "Account" ("Owner");

CREATE INDEX ON "Entries" ("Account_id");

CREATE INDEX ON "Transac" ("From_account_id");

CREATE INDEX ON "Transac" ("To_account_id");

CREATE INDEX ON "Transac" ("From_account_id", "To_account_id");

ALTER TABLE "Entries" ADD FOREIGN KEY ("Account_id") REFERENCES "Account" ("ID");

ALTER TABLE "Transac" ADD FOREIGN KEY ("From_account_id") REFERENCES "Account" ("ID");

ALTER TABLE "Transac" ADD FOREIGN KEY ("To_account_id") REFERENCES "Account" ("ID");