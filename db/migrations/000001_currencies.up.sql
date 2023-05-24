CREATE TABLE IF NOT EXISTS currencies (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  abbreviation VARCHAR UNIQUE NOT NULL,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);

INSERT INTO currencies(name, abbreviation) VALUES('Indonesian Rupiah', 'IDR');