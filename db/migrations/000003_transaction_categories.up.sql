CREATE TABLE IF NOT EXISTS transaction_categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR NOT NULL,
  description VARCHAR DEFAULT '',
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ
);

INSERT INTO transaction_categories (name, description)
VALUES('Food & Beverage', 'Food & Beverage description'),
  ('Shopping', 'Shopping description'),
  ('Transportation', 'Transportation description'),
  ('Bill & Utility', 'Bill & Utility description'),
  ('Travel', 'Travel description'),
  ('Family', 'Family description'),
  ('Education', 'Education description'),
  ('Investation', 'Investation description'),
  ('Business', 'Business description'),
  ('Gift', 'Gift description'),
  ('Freelance', 'Freelance description'),
  ('Fashion', 'Fashion description'),
  ('Tool', 'Tool description'),
  ('Book', 'Book description'),
  ('Health', 'Health description');