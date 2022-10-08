-- level
--   0 = NONE
--   1 = super admin
--   2 = admin
CREATE TABLE admins (
    id BIGINT PRIMARY KEY NOT NULL UNIQUE,
    level SMALLINT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE bilibili_users (
    uid BIGINT PRIMARY KEY NOT NULL UNIQUE,
    counter BIGINT NOT NULL DEFAULT 0,
    is_whitelist BOOLEAN NOT NULL DEFAULT FALSE,
    ban_until TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE records (
    record_id SERIAL PRIMARY KEY,
    uid BIGINT NOT NULL,
    description TEXT NOT NULL,
    chat_id BIGINT,
    message_id BIGINT,
    approved_by BIGINT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);
CREATE TABLE reports (
    report_id SERIAL PRIMARY KEY,
    uid BIGINT NOT NULL,
    description TEXT NOT NULL,
    file_type SMALLINT NOT NULL,
    file_id TEXT NOT NULL,
    submit_by TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc'),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT (now() at time zone 'utc')
);