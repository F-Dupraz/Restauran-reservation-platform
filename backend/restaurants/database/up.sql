DROP TABLE IF EXISTS restaurants;

CREATE TABLE restaurants (
  id VARCHAR(50) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  owner VARCHAR(255) NOT NULL,
  city VARCHAR(100) NOT NULL,
  days_open VARCHAR(100)[],
  specialties VARCHAR(255)[],
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);