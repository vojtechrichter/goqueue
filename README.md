# GoQueue
Background job queue backed by Postgres, and exposed via a HTTP API, allowing workers to be
spawned using any other technology, capable of performing a HTTP request.

## Design

### Postgres schema

#### `job` table
- `id`
- `kind`
- `args` (jsonb)
- `state` (enum)
- `attempt`
- `max_attempts`
- `scheduled_at`
- `attempted_at`
- `errors`
- `created_at`

### The state machine
`available` -> `running` -> `completed`
`running` -> `retryable` -> `running`
`running` -> `discarded`
