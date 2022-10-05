-- level
--   0 = NONE
--   1 = super admin
--   2 = admin
CREATE TABLE admins (
    id BIGINT PRIMARY KEY NOT NULL UNIQUE,
    level SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    modified_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE bilibili_users (
    uid BIGINT PRIMARY KEY NOT NULL UNIQUE,
    counter BIGINT NOT NULL DEFAULT 0,
    is_whitelist BOOLEAN NOT NULL DEFAULT FALSE,
    ban_until TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    modified_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE records (
    record_id SERIAL PRIMARY KEY,
    uid BIGINT NOT NULL,
    description TEXT NOT NULL,
    chat_id BIGINT,
    message_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    modified_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE reports (
    report_id SERIAL PRIMARY KEY,
    uid BIGINT NOT NULL,
    description TEXT NOT NULL,
    file_type SMALLINT NOT NULL,
    file_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc'),
    modified_at TIMESTAMP NOT NULL DEFAULT (now() at time zone 'utc')
);