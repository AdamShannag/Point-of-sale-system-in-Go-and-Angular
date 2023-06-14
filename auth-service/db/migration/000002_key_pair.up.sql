CREATE TABLE key_pair
(
    uuid        varchar PRIMARY KEY,
    privet_key  BYTEA   NOT NULL,
    public_key  BYTEA   NOT NULL,
    expired_at  timestamp NOT NULL DEFAULT (now() + interval '60 minutes'),
    created_at  timestamp NOT NULL DEFAULT (now()),
    modified_at timestamp NOT NULL DEFAULT (now())
);

insert into key_pair (uuid, privet_key, public_key, expired_at)
values ('484e0085-2c7c-474b-b00d-25e3ee23815e', '', '', now())