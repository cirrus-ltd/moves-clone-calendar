CREATE TABLE rate_calendar (
    id DATE PRIMARY KEY,
    version INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    rate INT
);
