CREATE TABLE IF NOT EXISTS `Users` (
`id` bigint NOT NULL AUTO_INCREMENT,
`role_id` bigint NOT NULL,
`name`varchar(100) NOT NULL DEFAULT 0,
`email` varchar(100) NOT NULL,
`password` varchar(100) NULL,
`status` tinyint NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_user_id` (`id`),
KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;