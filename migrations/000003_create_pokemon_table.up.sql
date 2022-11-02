CREATE TABLE IF NOT EXISTS `Pokemon` (
`id` bigint NOT NULL AUTO_INCREMENT,
`type_id` bigint NOT NULL DEFAULT 0,
`name`varchar(100) NOT NULL DEFAULT 0,
`height` double NOT NULL DEFAULT 0,
`weight` double NOT NULL DEFAULT 0,
`stock` double NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_pokemon_id` (`id`),
KEY `idx_type_id` (`type_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;