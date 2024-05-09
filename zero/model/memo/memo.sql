CREATE TABLE `gorm_memo` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `content` varchar(255) NOT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;