DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(100) NOT NULL,
  owner_of VARCHAR(50)[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS restaurants;

CREATE TABLE restaurants (
  id VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  owner VARCHAR(50) NOT NULL REFERENCES users(id),
  address VARCHAR(255) NOT NULL,
  description TEXT,
  city VARCHAR(100) NOT NULL,
  days_open VARCHAR(100)[],
  capacity INT[],
  specialties VARCHAR(255)[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS reservations;

CREATE TABLE reservations (
  id VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
  user_id VARCHAR(50) NOT NULL REFERENCES users(id),
  restaurant_id VARCHAR(50) NOT NULL REFERENCES restaurants(id),
  day VARCHAR(100) NOT NULL,
  num_guests INT NOT NULL,
  is_done BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
