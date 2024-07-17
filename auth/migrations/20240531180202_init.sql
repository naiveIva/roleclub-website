-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_role AS ENUM
(
    'player',
    'master',
    'admin'
);

CREATE TABLE IF NOT EXISTS players
(
    uuid            uuid        NOT NULL    DEFAULT uuid_generate_v4(),
    first_name      varchar     NOT NULL,
    last_name       varchar     NOT NULL,
    father_name     varchar     NOT NULL,
    tel_number      varchar     NOT NULL    UNIQUE,
    password_hash   varchar     NOT NULL,
    is_hse_student  boolean     NOT NULL,
    is_banned       boolean     NOT NULL    DEFAULT false,
    played_games    bigint      NOT NULL    DEFAULT 0,
    conducted_games bigint      NOT NULL    DEFAULT 0,
    status          user_role   NOT NULL    DEFAULT 'player',
    PRIMARY KEY (uuid)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players CASCADE;

DROP TYPE IF EXISTS user_role;
-- +goose StatementEnd
