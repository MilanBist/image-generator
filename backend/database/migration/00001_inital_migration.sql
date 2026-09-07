-- +goose Up

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "userName" varchar(50) UNIQUE,
  "email" varchar(255) UNIQUE NOT NULL,
  "passwordHash" text NOT NULL,
  "createdAt" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "uploadedFiles" (
  "id" uuid PRIMARY KEY,
  "userId" uuid NOT NULL,
  "fileName" varchar(255) NOT NULL,
  "storageKey" text NOT NULL,
  "fileType" varchar(50) NOT NULL,
  "mimeType" varchar(100) NOT NULL,
  "fileSize" bigint,
  "width" integer,
  "height" integer,
  "createdAt" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "images" (
  "id" uuid PRIMARY KEY,
  "userId" uuid NOT NULL,
  "sourceFileId" uuid NOT NULL,
  "storageKey" text NOT NULL,
  "fileName" varchar(255),
  "mimeType" varchar(100),
  "width" integer,
  "height" integer,
  "fileSize" bigint,
  "createdAt" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "history" (
  "id" uuid PRIMARY KEY,
  "userId" uuid NOT NULL,
  "sourceFileId" uuid NOT NULL,
  "outputImageId" uuid NOT NULL,
  "operationType" varchar(50) NOT NULL,
  "prompt" text,
  "parameters" jsonb,
  "status" varchar(30) NOT NULL DEFAULT 'pending',
  "createdAt" timestamp NOT NULL DEFAULT (now()),
  "completedAt" timestamp
);

CREATE INDEX ON "uploadedFiles" ("userId");

CREATE INDEX ON "uploadedFiles" ("createdAt");

CREATE INDEX ON "uploadedFiles" ("userId", "createdAt");

CREATE INDEX ON "images" ("userId");

CREATE INDEX ON "images" ("sourceFileId");

CREATE INDEX ON "images" ("createdAt");

CREATE INDEX ON "images" ("userId", "createdAt");

CREATE INDEX ON "history" ("userId");

CREATE INDEX ON "history" ("sourceFileId");

CREATE INDEX ON "history" ("outputImageId");

CREATE INDEX ON "history" ("operationType");

CREATE INDEX ON "history" ("status");

CREATE INDEX ON "history" ("createdAt");

CREATE INDEX ON "history" ("userId", "createdAt");

CREATE INDEX ON "history" ("userId", "status");

ALTER TABLE "uploadedFiles" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "images" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "images" ADD FOREIGN KEY ("sourceFileId") REFERENCES "uploadedFiles" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "history" ADD FOREIGN KEY ("userId") REFERENCES "users" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "history" ADD FOREIGN KEY ("sourceFileId") REFERENCES "uploadedFiles" ("id") DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "history" ADD FOREIGN KEY ("outputImageId") REFERENCES "images" ("id") DEFERRABLE INITIALLY IMMEDIATE;


-- +goose Down

-- Remove foreign keys first
ALTER TABLE "history"
DROP CONSTRAINT IF EXISTS "history_userId_fkey";

ALTER TABLE "history"
DROP CONSTRAINT IF EXISTS "history_sourceFileId_fkey";

ALTER TABLE "history"
DROP CONSTRAINT IF EXISTS "history_outputImageId_fkey";

ALTER TABLE "images"
DROP CONSTRAINT IF EXISTS "images_userId_fkey";

ALTER TABLE "images"
DROP CONSTRAINT IF EXISTS "images_sourceFileId_fkey";

ALTER TABLE "uploadedFiles"
DROP CONSTRAINT IF EXISTS "uploadedFiles_userId_fkey";


-- Drop indexes
DROP INDEX IF EXISTS "uploadedFiles_userId_idx";
DROP INDEX IF EXISTS "uploadedFiles_createdAt_idx";
DROP INDEX IF EXISTS "uploadedFiles_userId_createdAt_idx";

DROP INDEX IF EXISTS "images_userId_idx";
DROP INDEX IF EXISTS "images_sourceFileId_idx";
DROP INDEX IF EXISTS "images_createdAt_idx";
DROP INDEX IF EXISTS "images_userId_createdAt_idx";

DROP INDEX IF EXISTS "history_userId_idx";
DROP INDEX IF EXISTS "history_sourceFileId_idx";
DROP INDEX IF EXISTS "history_outputImageId_idx";
DROP INDEX IF EXISTS "history_operationType_idx";
DROP INDEX IF EXISTS "history_status_idx";
DROP INDEX IF EXISTS "history_createdAt_idx";
DROP INDEX IF EXISTS "history_userId_createdAt_idx";
DROP INDEX IF EXISTS "history_userId_status_idx";


-- Drop tables in reverse dependency order
DROP TABLE IF EXISTS "history";
DROP TABLE IF EXISTS "images";
DROP TABLE IF EXISTS "uploadedFiles";
DROP TABLE IF EXISTS "users";