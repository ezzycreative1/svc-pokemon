CREATE TABLE IF NOT EXISTS `Battle` (
`id` bigint NOT NULL AUTO_INCREMENT,
`team_1` bigint NOT NULL DEFAULT 0,
`team_2` bigint NOT NULL DEFAULT 0,
`location` text NOT NULL,
`winner_id` bigint NOT NULL DEFAULT 0,
`score` double NOT NULL DEFAULT 0,
`start_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
`end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
KEY `idx_battle_id` (`id`),
KEY `idx_team_1_id` (`team_1`),
KEY `idx_team_2_id` (`team_2`),
KEY `idx_winner_id` (`winner_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;