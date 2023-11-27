
-- name: GetAPIKeys :many
SELECT * FROM api_keys_schema.api_keys;

-- name: CreateAPIKey :one
INSERT INTO api_keys_schema.api_keys(api_key_id, api_key_value, created_at) 
VALUES ($1, $2, $3)
RETURNING api_key_id; 

