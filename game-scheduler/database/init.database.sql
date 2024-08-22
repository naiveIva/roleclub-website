CREATE TYPE ggaame_state AS ENUM
(
    'playable',
    'unplayable'
)

CREATE TABLE IF NOT EXISTS games
(
    game_name           varchar     NOT NULL,
    authors             varchar     NOT NULL,
    description         text        NOT NULL,
    complexity          int         NOT NULL,
    players_per_session int         NOT NULL,
    masters_per_session int         NOT NULL,
    roles               varchar[]   NOT NULL,
    game_state          game_state  NOT NULL,
    PRIMARY KEY (game_name),
    CONSTRAINT valid_complexity CHECK (complexity >= 1 AND complexity <= 5),
    CONSTRAINT valid_players_per_session CHECK (players_per_session >= 1),
    CONSTRAINT valid_masters_per_session CHECK (masters_per_session >= 1)
);

CREATE TABLE IF NOT EXISTS scheduled_events
(
    id                   uuid           NOT NULL    DEFAULT uuid_generate_v4(),
    game_name            varchar        NOT NULL,
    event_date           timestamptz    NOT NULL,
    num_of_sessions      int            NOT NULL,
    is_subscription_open boolean        NOT NULL    DEFAULT FALSE,
    PRIMARY KEY (id),
    FOREIGN KEY (game_name) REFERENCES games (game_name) ON DELETE CASCADE
);