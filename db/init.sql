CREATE DATABASE IF NOT EXISTS test;

USE test;

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

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


-- default password is 123456
INSERT INTO `users` (`name`, `email`, `password`, `deleted_at`) VALUES 
('datnd', 'datnd@gmail.com', '$2a$10$2GtL6jZVXqum9PIppg4sUO/UjuEuNeTyrWJ0JD00691xF1Z7hO9Pm', NULL);