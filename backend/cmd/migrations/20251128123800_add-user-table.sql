-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    firstName VARCHAR(255) NOT NULL,
    lastName VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (firstName, lastName, email, password) 
VALUES 
    ('user1', 'user1', 'user1@mail.com', '$2a$10$59cO02HhDzowFt.FA8Q3quC3mEHEaugpVUeMZFZ4Ll9r3PpnCqvzK'),
    ('user2', 'user2', 'user2@mail.com', '$2a$10$59cO02HhDzowFt.FA8Q3quC3mEHEaugpVUeMZFZ4Ll9r3PpnCqvzK'),
    ('user3', 'user3', 'user3@mail.com', '$2a$10$59cO02HhDzowFt.FA8Q3quC3mEHEaugpVUeMZFZ4Ll9r3PpnCqvzK');


-- +goose Down
DROP TABLE IF EXISTS users;