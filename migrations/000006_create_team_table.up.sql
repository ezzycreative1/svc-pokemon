CREATE TABLE IF NOT EXISTS `Teams` (
`id` bigint NOT NULL AUTO_INCREMENT,
`trainer_id` bigint NOT NULL DEFAULT 0,
`pokemon_id` bigint NOT NULL DEFAULT 0,
`battle_id` bigint NOT NULL DEFAULT 0,
`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_team_id` (`id`),
KEY `idx_trainer_id` (`trainer_id`),
KEY `idx_pokemon_id` (`pokemon_id`),
KEY `idx_battle_id` (`battle_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;