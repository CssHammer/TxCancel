-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE "users"
(
    "id"      SERIAL PRIMARY KEY,
    "name"    VARCHAR(50) NOT NULL,
    "balance" NUMERIC     NOT NULL CHECK ( "balance" >= 0 )
);

INSERT INTO "users"(name, balance)
VALUES ('Dave', 0);

CREATE TABLE "transactions"
(
    "id"             SERIAL PRIMARY KEY,
    "state"          VARCHAR(10)  NOT NULL,
    "amount"         NUMERIC      NOT NULL,
    "transaction_id" VARCHAR(100) NOT NULL UNIQUE,
    "canceled"       BOOLEAN      NOT NULL
);

CREATE TABLE "counters"
(
    "name"  VARCHAR(255) PRIMARY KEY NOT NULL,
    "value" INTEGER                  NOT NULL
);

INSERT INTO "counters"(name, value)
VALUES ('last_cancelled_id', 0);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE "users";
DROP TABLE "transactions";
DROP TABLE "counters";