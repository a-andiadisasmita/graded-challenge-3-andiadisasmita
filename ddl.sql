-- ddl.sql: Create tables and seed data

-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    age INT NOT NULL
);

-- User activity logs table
CREATE TABLE user_activity_logs (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    description VARCHAR(255) NOT NULL
);

-- Posts table
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id)
);

-- Comments table
CREATE TABLE comments (
    id SERIAL PRIMARY KEY,
    content VARCHAR(255) NOT NULL,
    user_id INT REFERENCES users(id),
    post_id INT REFERENCES posts(id)
);

-- Seed data for Users
INSERT INTO users (full_name, email, username, password, age) VALUES
('Alice Johnson', 'alice@gmail.com', 'alicej', 'password123', 25),
('Bob Smith', 'bob@yahoo.com', 'bobsmith', 'securepass456', 30),
('Charlie Brown', 'charlie@hotmail.com', 'charlieb', 'password789', 22),
('Daisy Miller', 'daisy@gmail.com', 'daisym', 'mypassword321', 28),
('Eve Wilson', 'eve@outlook.com', 'evewilson', '1234secure', 35),
('Frank Adams', 'frank@gmail.com', 'franka', '7890password', 27),
('Grace Lee', 'gracelee@company.com', 'gracelee', 'leeSecure123', 29);

-- Seed data for User Activity Logs
INSERT INTO user_activity_logs (user_id, description) VALUES
(1, 'Registered new account'),
(2, 'Registered new account'),
(3, 'Registered new account'),
(4, 'Registered new account'),
(5, 'Registered new account'),
(6, 'Registered new account'),
(7, 'Registered new account');

-- Seed data for Posts
INSERT INTO posts (content, image_url, user_id) VALUES
('Hello World!', 'http://example.com/image1.jpg', 1),
('My first post!', 'http://example.com/image2.jpg', 2),
('Check out this cool pic', 'http://example.com/image3.jpg', 3),
('Sunny day!', 'http://example.com/image4.jpg', 4),
('Feeling great!', 'http://example.com/image5.jpg', 5),
('Nature is beautiful', 'http://example.com/image6.jpg', 6),
('Family time', 'http://example.com/image7.jpg', 7);

-- Seed data for Comments
INSERT INTO comments (content, user_id, post_id) VALUES
('Nice post!', 2, 1),
('Great pic!', 3, 2),
('Amazing!', 4, 3),
('Lovely view', 5, 4),
('So true!', 6, 5),
('Awesome!', 7, 6),
('Cool!', 1, 7);
