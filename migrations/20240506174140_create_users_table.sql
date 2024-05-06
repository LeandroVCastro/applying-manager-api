-- Create "users" table
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  `name` longtext NULL,
  `email` varchar(191) NOT NULL,
  `password` longtext NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_users_deleted_at` (`deleted_at`),
  UNIQUE INDEX `uni_users_email` (`email`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
