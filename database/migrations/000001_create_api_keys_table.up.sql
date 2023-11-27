CREATE SCHEMA api_keys_schema;

CREATE TABLE api_keys_schema.api_keys (
    api_key_id SERIAL PRIMARY KEY,
    api_key_value TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);


