-- Create "companies" table
CREATE TABLE `companies` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `description` longtext NULL,
  `website` longtext NULL,
  `linkedin` longtext NULL,
  `glassdoor` longtext NULL,
  `instagram` longtext NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_companies_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "platforms" table
CREATE TABLE `platforms` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` longtext NOT NULL,
  `website` longtext NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_platforms_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "applyments" table
CREATE TABLE `applyments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext NOT NULL,
  `description` longtext NULL,
  `link` longtext NULL,
  `company_id` bigint unsigned NULL,
  `platform_id` bigint unsigned NULL,
  `applied_at` datetime(3) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_companies_applyments` (`company_id`),
  INDEX `fk_platforms_applyments` (`platform_id`),
  INDEX `idx_applyments_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_companies_applyments` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_platforms_applyments` FOREIGN KEY (`platform_id`) REFERENCES `platforms` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
