-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Switch to using the `snippetbox` database.
USE snippetbox;

-- created from docker compose instead as MYSQL_USER
CREATE USER 'web'@'localhost';
-- 'MYSQL_USER'@'%' if created in docker compose
GRANT SELECT, INSERT ON snippetbox.* TO 'web'@'localhost';
-- created from docker compose as MYSQL_PASSWORD
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
