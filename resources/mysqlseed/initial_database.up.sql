CREATE DATABASE IF NOT EXISTS gotemplate
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

use gotemplate;





CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `company_code` NVARCHAR(250) NULL DEFAULT '',
  `user_id` NVARCHAR(250) NULL DEFAULT '',
  `user_name` NVARCHAR(250) NULL DEFAULT '',
  `user_role_id` BIGINT NULL DEFAULT 0,
  `user_group_ids` NVARCHAR(250) NULL DEFAULT '',
  `user_info` BLOB NULL,
  `user_state` INT NOT NULL DEFAULT 0,
  `thumbnail_image_url`  NVARCHAR(2500) NULL DEFAULT '',
  `last_modified` BIGINT NULL DEFAULT 0,
  `issued_date` BIGINT NULL DEFAULT 0,
  `activation_date` BIGINT NULL DEFAULT 0,
  `expiry_date` BIGINT NULL DEFAULT 0,
  `reference_id` NVARCHAR(250) NULL DEFAULT '',
  PRIMARY KEY (`id`));

CREATE INDEX  idx_user_id ON users(user_id); 