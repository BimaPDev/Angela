CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE items (
  id SERIAL PRIMARY KEY,
  owner_id INT NOT NULL REFERENCES users(id),
  image_url TEXT NOT NULL,
  category VARCHAR(50),
  color VARCHAR(50),
  style VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE friendships (
  user_id_1 INT NOT NULL REFERENCES users(id),
  user_id_2 INT NOT NULL REFERENCES users(id),
  status VARCHAR(20) DEFAULT 'accepted',
  PRIMARY KEY (user_id_1, user_id_2)
);