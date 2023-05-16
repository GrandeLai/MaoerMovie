CREATE TABLE `comment` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `film_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '电影ID',
    `content` varchar(255)  NOT NULL DEFAULT '' COMMENT '内容',
    `score` float(20, 0) NOT NULL DEFAULT 0.0 COMMENT '影片评分',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_film_id` (`film_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
