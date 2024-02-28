CREATE DATABASE "online_payment_platform";
\c online_payment_platform;

CREATE TABLE "merchants" (
    "id" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "country" CHAR(2) NOT NULL,
    "currency" CHAR(3) NOT NULL,
    "website" VARCHAR(255) NOT NULL,
    "webhook_url" VARCHAR(255) NOT NULL,
    "api_key" VARCHAR(255) NOT NULL,
    "created" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY ("id"),
    CONSTRAINT "unique_email" UNIQUE ("email"),
    CONSTRAINT "unique_api_key" UNIQUE ("api_key")
);

CREATE INDEX "idx_merchants_email" ON "merchants" ("email");
CREATE INDEX "idx_merchants_api_key" ON "merchants" ("api_key");


CREATE TABLE "customers" (
    "id" VARCHAR(255) NOT NULL,
    "merchant_id" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "phone" VARCHAR(255) NOT NULL,
    "address" VARCHAR(255) NOT NULL,
    "country" CHAR(2) NOT NULL,
    "created" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY ("id"),
    CONSTRAINT "fk_customers_merchant_id" FOREIGN KEY ("merchant_id") REFERENCES "merchants" ("id") ON DELETE CASCADE
);

CREATE INDEX "idx_customers_merchant_id" ON "customers" ("merchant_id");
CREATE INDEX "idx_customers_email" ON "customers" ("email");
