CREATE DATABASE IF NOT EXISTS taskmanagmentsystem;


CREATE TABLE IF NOT EXISTS users(
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role  ENUM('user','manager','admin') DEFAULT 'user'
);


CREATE TABLE IF NOT EXISTS tasks (
    user_id INT ,
    id INT AUTO_INCREMENT PRIMARY KEY,             
    title VARCHAR(255) NOT NULL,                   
    description VARCHAR(255),                             
    priority ENUM('low', 'middle', 'high') DEFAULT 'low', 
    status ENUM('pending', 'done', 'outstanding') DEFAULT 'pending', 
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,   
    to_done TIMESTAMP NULL                     
);


CREATE TABLE IF NOT EXISTS task_dependencies (
    user_id INT,
    task_id INT,                                   
    dependent_task_id INT
);

CREATE TABLE IF NOT EXISTS logs (
    user_id INT,
    task_id INT,  
    date TIMESTAMP,
    action VARCHAR(255)
)