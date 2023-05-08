CREATE TABLE IF NOT EXISTS transaction_categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  description VARCHAR DEFAULT '',
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);