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
-- Create "stages" table
CREATE TABLE `stages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext NOT NULL,
  `description` longtext NOT NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `idx_stages_deleted_at` (`deleted_at`)
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Create "applyments" table
CREATE TABLE `applyments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext NOT NULL,
  `description` longtext NULL,
  `link` longtext NULL,
  `company_id` bigint unsigned NULL,
  `platform_id` bigint unsigned NULL,
  `stage_id` bigint unsigned NULL,
  `applied_at` datetime(3) NULL,
  `created_at` datetime(3) NULL,
  `updated_at` datetime(3) NULL,
  `deleted_at` datetime(3) NULL,
  PRIMARY KEY (`id`),
  INDEX `fk_companies_applyments` (`company_id`),
  INDEX `fk_platforms_applyments` (`platform_id`),
  INDEX `fk_stages_applyments` (`stage_id`),
  INDEX `idx_applyments_deleted_at` (`deleted_at`),
  CONSTRAINT `fk_companies_applyments` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_platforms_applyments` FOREIGN KEY (`platform_id`) REFERENCES `platforms` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `fk_stages_applyments` FOREIGN KEY (`stage_id`) REFERENCES `stages` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- Insert data on `stages` table
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Applied', "Just applied. There's no answer yet.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('In progress', "Got a contact from the recruiter. But no practical test has yet been proposed.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Testing/talking to the company', "Taking practical tests or talking to the company.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Proposal received', "Received a proposal from the company.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Proposal accepted', "Candidate accepted the proposal made by the company.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Proposal denied', "Candidate denied the proposal made by the company.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Disqualified', "Candidate was disqualified from the selection process.", now(), now());
INSERT INTO `stages` (`title`, `description`, `created_at`, `updated_at`) values ('Paralyzed position', "The selection process was stopped by the company.", now(), now());