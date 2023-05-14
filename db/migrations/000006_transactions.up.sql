CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  transaction_type_id INT NOT NULL,
  transaction_category_id INT NOT NULL,
  currency_id INT NOT NULL,
  transaction_date DATE NOT NULL,
  note TEXT,
  amount NUMERIC(15),
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ,

  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY(transaction_type_id) REFERENCES transaction_types(id),
  FOREIGN KEY(transaction_category_id) REFERENCES transaction_categories(id),
  FOREIGN KEY(currency_id) REFERENCES currencies(id)
);