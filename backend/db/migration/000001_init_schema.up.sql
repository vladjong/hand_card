CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    login VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS cards (
    id SERIAL PRIMARY KEY,
    organization VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL UNIQUE,
    category_id BIGINT REFERENCES categories (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_cards (
    user_id BIGINT REFERENCES users (id) NOT NULL,
    card_id BIGINT REFERENCES cards (id) NOT NULL UNIQUE
);
