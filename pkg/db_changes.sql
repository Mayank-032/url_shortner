CREATE TABLE `url_mapper` (
  `id` int NOT NULL AUTO_INCREMENT,
  `short_url` text NOT NULL,
  `long_url` longtext NOT NULL,
  `is_hash_signed` tinyint(1) DEFAULT '0',
  `hash` varchar(45) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash_idx_key` (`hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
