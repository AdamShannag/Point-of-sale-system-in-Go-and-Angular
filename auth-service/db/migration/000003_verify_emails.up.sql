CREATE TABLE "verify_emails"
(
    "uuid"        varchar PRIMARY KEY,
    "user_uuid"   varchar        NOT NULL,
    "email"       varchar     NOT NULL,
    "secret_code" varchar     NOT NULL,
    "is_used"     bool        NOT NULL DEFAULT false,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "expired_at"  timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

ALTER TABLE "verify_emails"
    ADD FOREIGN KEY ("user_uuid") REFERENCES "users" ("uuid");

ALTER TABLE "users"
    ADD COLUMN "is_email_verified" bool NOT NULL DEFAULT false;
