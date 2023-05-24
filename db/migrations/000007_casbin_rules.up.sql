CREATE TABLE IF NOT EXISTS casbin_rules(
    id SERIAL PRIMARY KEY,
    p_type VARCHAR NOT NULL,
    v0 VARCHAR NOT NULL DEFAULT '',
    v1 VARCHAR NOT NULL DEFAULT '',
    v2 VARCHAR NOT NULL DEFAULT '',
    v3 VARCHAR NOT NULL DEFAULT '',
    v4 VARCHAR NOT NULL DEFAULT '',
    v5 VARCHAR NOT NULL DEFAULT ''
);

INSERT INTO casbin_rules (p_type,v0,v1,v2) VALUES
     ('p', 'transaction_categories', 'transaction_categories', 'create'),
     ('p', 'transaction_categories', 'transaction_categories', 'read'),
     ('p', 'transaction_categories', 'transaction_categories', 'update'),
     ('p', 'transaction_categories', 'transaction_categories', 'delete'),

     ('p', 'transaction_types', 'transaction_types', 'create'),
     ('p', 'transaction_types', 'transaction_types', 'read'),
     ('p', 'transaction_types', 'transaction_types', 'update'),
     ('p', 'transaction_types', 'transaction_types', 'delete'),

     ('p', 'currencies', 'currencies', 'create'),
     ('p', 'currencies', 'currencies', 'read'),
     ('p', 'currencies', 'currencies', 'update'),
     ('p', 'currencies', 'currencies', 'delete'),

     ('p', 'users', 'users', 'read'),
     ('p', 'users', 'users', 'change_password'),
     ('p', 'users', 'users', 'delete'),

     ('p', 'transactions', 'transactions', 'create'),
     ('p', 'transactions', 'transactions', 'read'),
     ('p', 'transactions', 'transactions', 'update'),
     ('p', 'transactions', 'transactions', 'delete'),

     ('p', 'profile', 'profile', 'read'),
     ('p', 'profile', 'profile', 'change_password'),

     ('g', 'admin', 'transaction_categories', 'create'),
     ('g', 'admin', 'transaction_categories', 'read'),
     ('g', 'admin', 'transaction_categories', 'update'),
     ('g', 'admin', 'transaction_categories', 'delete'),
     ('g', 'admin', 'transaction_types', 'create'),
     ('g', 'admin', 'transaction_types', 'read'),
     ('g', 'admin', 'transaction_types', 'update'),
     ('g', 'admin', 'transaction_types', 'delete'),
     ('g', 'admin', 'currencies', 'create'),
     ('g', 'admin', 'currencies', 'read'),
     ('g', 'admin', 'currencies', 'update'),
     ('g', 'admin', 'currencies', 'delete'),
     ('g', 'admin', 'users', 'read'),
     ('g', 'admin', 'users', 'change_password'),
     ('g', 'admin', 'users', 'delete'),

     ('g', '1', 'admin', ''),
     
     ('p', 'transactions', 'transactions', 'create'),
     ('p', 'transactions', 'transactions', 'read'),
     ('p', 'transactions', 'transactions', 'update'),
     ('p', 'transactions', 'transactions', 'delete'),

     ('p', 'profile', 'profile', 'read'),
     ('p', 'profile', 'profile', 'change_password'),

     ('g', 'user', 'transaction_types', 'read'),
     ('g', 'user', 'transaction_categories', 'read'),
     ('g', 'user', 'currencies', 'read'),
     ('g', 'user', 'transactions', 'create'),
     ('g', 'user', 'transactions', 'read'),
     ('g', 'user', 'transactions', 'update'),
     ('g', 'user', 'transactions', 'delete'),

     ('g', 'user', 'profile', 'read'),
     ('g', 'user', 'profile', 'change_password'),

     ('g', 'usr', 'user', '');