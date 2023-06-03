CREATE TABLE "sessions"
(
    "uuid"          varchar PRIMARY KEY,
    "user_uuid"     varchar        NOT NULL,
    "refresh_token" varchar     NOT NULL,
    "user_agent"    varchar     NOT NULL,
    "client_ip"     varchar     NOT NULL,
    "is_blocked"    boolean     NOT NULL DEFAULT false,
    "expires_at"    timestamptz NOT NULL,
    "created_at"    timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions"
    ADD FOREIGN KEY ("user_uuid") REFERENCES "users" ("uuid");
