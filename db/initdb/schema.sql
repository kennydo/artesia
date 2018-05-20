CREATE TABLE "users" (
    "id" BIGSERIAL NOT NULL,
    "email" VARCHAR(256) NOT NULL,
    "password_hash" VARCHAR(128) NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "uq_email" ON "users" (lower("email"));
