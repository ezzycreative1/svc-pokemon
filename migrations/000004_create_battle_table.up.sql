CREATE TABLE IF NOT EXISTS `Battle` (
`id` bigint NOT NULL AUTO_INCREMENT,
`pokemon_id` bigint NOT NULL,
`score` double NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_battle_id` (`id`),
KEY `idx_pokemon_id` (`pokemon_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;