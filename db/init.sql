CREATE DATABASE IF NOT EXISTS test;

USE test;

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) UNIQUE NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL  DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE INDEX `users_index_4` ON `users` (`name`);

CREATE INDEX `users_index_5` ON `users` (`email`);

-- default password is 123456
INSERT INTO `users` (`name`, `email`, `password`, `deleted_at`) VALUES 
('datnd', 'datnd@gmail.com',  '$2a$10$2GtL6jZVXqum9PIppg4sUO/UjuEuNeTyrWJ0JD00691xF1Z7hO9Pm', NULL),
('Chanda Toon', 'ctoon1@nbcnews.com', '+7 432 590 9266', 2, '$2a$10$2GtL6jZVXqum9PIppg4sUO/UjuEuNeTyrWJ0JD00691xF1Z7hO9Pm', NULL),

