CREATE SCHEMA avitoApp;

CREATE TABLE avitoApp.users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL CHECK(phone_number ~ '^\+\d{8,15}$')
);

CREATE TABLE avitoApp.product(
    id SERIAL PRIMARY KEY,
    title VARCHAR(300) NOT NULL,
    description VARCHAR(300),
    price INTEGER NULL CHECK (price > 0),
    category VARCHAR(100) NOT NULL,

    author_product_id INTEGER NOT NULL REFERENCES avitoApp.users(id)
);