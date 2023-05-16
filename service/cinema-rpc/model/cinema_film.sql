CREATE TABLE `cinema_film`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `film_id` bigint NOT NULL DEFAULT 0 COMMENT '电影编号',
  `cinema_id` bigint NOT NULL DEFAULT 0 COMMENT '电影院编号',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影厅电影信息表';

INSERT INTO `cinema_film` VALUES (1, 35349973142863883,2);
INSERT INTO `cinema_film` VALUES (2, 36053463016144911,2);