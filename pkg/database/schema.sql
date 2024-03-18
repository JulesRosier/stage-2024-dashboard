CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    event_timestamp TIMESTAMP NOT NULL,
    topic_name VARCHAR(255) NOT NULL,
    topic_offset BIGINT NOT NULL,
    topic_partition INTEGER NOT NULL,
    event_headers JSONB,
    event_key JSONB,
    event_value JSONB
);