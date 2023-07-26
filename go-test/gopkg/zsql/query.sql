-- name: create-users-table
CREATE TABLE users (
                       id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
                       name VARCHAR(255),
                       email VARCHAR(255)
);

-- name: create-user
INSERT INTO users (name, email) VALUES(?, ?);

-- name: find-users-by-email
SELECT id,name,email FROM users WHERE email = ?;

-- :type find-one-user-by-email
SELECT id,name,email FROM users WHERE email = ? LIMIT 1;

-- @tag: drop-users-table
DROP TABLE users;