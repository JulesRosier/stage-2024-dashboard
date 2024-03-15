CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    data JSONB,
    topic_name VARCHAR(255)
);