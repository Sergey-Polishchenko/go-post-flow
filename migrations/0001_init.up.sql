CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL
);

CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  author_id INTEGER REFERENCES users(id),
  allow_comments BOOLEAN DEFAULT TRUE
);

CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  text TEXT NOT NULL,
  author_id INTEGER REFERENCES users(id),
  post_id INTEGER REFERENCES posts(id),
  parent_id INTEGER REFERENCES comments(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
