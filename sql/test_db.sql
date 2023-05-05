CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- created from docker compose instead as MYSQL_USER
CREATE USER 'test_web'@'localhost';
-- 'MYSQL_USER'@'%' if created in docker compose
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE
ON test_snippetbox.* TO 'test_web'@'localhost';

-- created from docker compose as MYSQL_PASSWORD
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
