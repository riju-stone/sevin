-- extension for generating uuids
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tasks(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    command TEXT NOT NULL,
    scheduled_at TIMESTAMP NOT NULL,
    picked_at TIMESTAMP,
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    failed_at TIMESTAMP
);

CREATE INDEX idx_task_scheduled_at ON tasks(scheduled_at);
