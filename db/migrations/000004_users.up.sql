CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(16) UNIQUE NOT NULL,
  password VARCHAR NOT NULL,
  full_name VARCHAR NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ
);

-- password: satudua3
INSERT INTO users(username, password, full_name) VALUES 
('admin', '$2a$10$p1.jrAFTR/YUGy3EiG3PcONQCmLwC0.u2WAvK4N9IG.2QMVVTjy1C', 'Admin'), 
('user_1', '$2a$10$p1.jrAFTR/YUGy3EiG3PcONQCmLwC0.u2WAvK4N9IG.2QMVVTjy1C', 'User Satu');
