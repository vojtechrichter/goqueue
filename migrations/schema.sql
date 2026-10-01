CREATE TYPE job_state AS ENUM ('available', 'running', 'retryable', 'completed', 'discarded');

CREATE TABLE job (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    kind text,
    args jsonb NOT NULL DEFAULT '{}',
    state job_state NOT NULL,
    attempt int NOT NULL,
    max_attempts int NOT NULL,
    scheduled_at timestamptz NOT NULL,
    attempted_at timestamptz NOT NULL,
    errors jsonb NOT NULL DEFAULT '[]',
    created_at timestamptz NOT NULL DEFAULT now()
);
