CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    eventhub_timestamp TIMESTAMP WITH TIME ZONE NOT NULL,
    event_timestamp TIMESTAMP WITH TIME ZONE,
    schema_id INTEGER,
    schema_format TEXT NOT NULL,
    topic_name VARCHAR(255) NOT NULL,
    topic_offset BIGINT NOT NULL,
    topic_partition INTEGER NOT NULL,
    event_type TEXT NOT NULL,
    event_headers JSONB,
    event_key JSONB,
    event_value JSONB,
    last_indexed_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS event_index_configs (
    id SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    event_type VARCHAR(255) NOT NULL,
    key_selector VARCHAR(255)[] NOT NULL,
    index_column VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS timestamp_configs (
    id SERIAL PRIMARY KEY,
    inserted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    event_type VARCHAR(255) NOT NULL,
    key_selector VARCHAR(255)[] NOT NULL
);