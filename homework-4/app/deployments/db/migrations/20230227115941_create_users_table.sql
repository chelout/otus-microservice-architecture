-- +goose Up
-- +goose StatementBegin
CREATE SEQUENCE IF NOT EXISTS users_id_seq;
CREATE TABLE users (
    "id" int8 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "email" text,
    "first_name" text,
    "last_name" text,
    "phone" text,
    "username" text,
    PRIMARY KEY ("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP SEQUENCE IF EXISTS users_id_seq;
-- +goose StatementEnd
