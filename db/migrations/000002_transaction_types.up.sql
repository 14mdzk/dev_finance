CREATE TABLE IF NOT EXISTS transaction_types (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  description VARCHAR DEFAULT '',
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);