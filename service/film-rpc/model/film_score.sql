CREATE TABLE `film_score`  (
   `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
   `film_id` bigint NOT NULL COMMENT '电影编号',
   `film_score` float(20, 0) NOT NULL DEFAULT 0.0 COMMENT '影片评分',
   `film_score_num` bigint NOT NULL DEFAULT 0 COMMENT '评分人数,以万为单位',
   PRIMARY KEY (`id`) USING BTREE,
   INDEX `film_id`(`film_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片评分表';