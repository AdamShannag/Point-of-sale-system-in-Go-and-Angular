CREATE TABLE "users"
(
    "uuid"            varchar PRIMARY KEY,
    "username"        varchar UNIQUE NOT NULL,
    "email"           varchar UNIQUE NOT NULL,
    "phone"           varchar UNIQUE NOT NULL,
    "hashed_password" varchar        NOT NULL,
    "address"         varchar        NOT NULL,
    "user_type"       varchar        NOT NULL,
    "added_by"        varchar        NOT NULL,
    "created_at"      timestamp      NOT NULL DEFAULT (now()),
    "modified_at"     timestamp      NOT NULL DEFAULT (now())
);