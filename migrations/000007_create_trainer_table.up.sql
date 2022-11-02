CREATE TABLE IF NOT EXISTS `Trainers` (
`id` bigint NOT NULL AUTO_INCREMENT,
`name`varchar(100) NOT NULL DEFAULT 0,
`gender` varchar(10) NOT NULL DEFAULT 0,
`country_code` varchar(5) NOT NULL DEFAULT 0,
`is_active` tinyint NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_trainer_id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;