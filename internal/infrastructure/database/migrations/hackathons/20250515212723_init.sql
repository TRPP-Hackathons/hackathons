-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    image_url  VARCHAR      NOT NULL
);
CREATE INDEX idx_user_id ON users(id);

CREATE TABLE hackathons
(
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(200) NOT NULL,
    money           INTEGER      NOT NULL CHECK (money > 0),
    participant_min SMALLINT
        CONSTRAINT check_participant_min CHECK (participant_min > 0),
    participant_max SMALLINT
        CONSTRAINT check_participant_max CHECK (participant_max < 20),
    image_url       VARCHAR      NOT NULL
);
CREATE INDEX idx_hackathons_id ON hackathons(id);
CREATE INDEX idx_hackathons_name ON hackathons(name);

CREATE TYPE role_type AS ENUM ('Backend-разработчик', 'Frontend-разработчик', 'Дизайнер');

CREATE TABLE participants
(
    user_id      INTEGER   NOT NULL,
    hackathon_id INTEGER   NOT NULL,
    role         role_type NOT NULL,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id)
        REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_hackathon_id FOREIGN KEY (hackathon_id)
        REFERENCES hackathons (id) ON DELETE CASCADE
);
CREATE INDEX idx_participants_user_id ON participants(user_id);
CREATE INDEX idx_participants_hackathon_id ON participants(hackathon_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS participants;
DROP TABLE IF EXISTS hackathons;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_user_id;
DROP INDEX IF EXISTS idx_hackathons_id;
DROP INDEX IF EXISTS idx_hackathons_name;
DROP INDEX IF EXISTS idx_participants_user_id;
DROP INDEX IF EXISTS idx_participants_hackathon_id;

DROP TYPE IF EXISTS role_type;
-- +goose StatementEnd
