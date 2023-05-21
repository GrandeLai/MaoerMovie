CREATE TABLE `show`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `cinema_id` bigint NOT NULL DEFAULT 0 COMMENT '影院编号',
  `film_id` bigint NOT NULL DEFAULT 0 COMMENT '电影编号',
  `begin_time` varchar(50) NOT NULL DEFAULT '' COMMENT '开始时间',
  `end_time` varchar(50) NOT NULL DEFAULT '' COMMENT '结束时间',
  `hall_id` bigint NOT NULL DEFAULT 0 COMMENT '放映厅类型编号',
  `price` double NOT NULL DEFAULT 0.0 COMMENT '票价',
  `date` timestamp NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '放映日期',
  `film_language` varchar(50) NOT NULL DEFAULT '' COMMENT '电影语言',
  `surplus_num` bigint NOT NULL DEFAULT 0 COMMENT '剩余票数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '放映场次表';

INSERT INTO `show` VALUES (1, 1, 2, '09:50', '11:20', 1,  19.9,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (2, 1, 3, '11:50', '13:20', 2,  30,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (3, 1, 3, '13:50', '15:20', 3,  40,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (4, 1, 2, '15:50', '17:20', 3,  60,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (5, 2, 2, '11:50', '13:20', 4,  22,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (6, 2, 2, '11:50', '13:20', 5,  44,'2023-05-13','国语2D',24);
INSERT INTO `show` VALUES (7, 2, 2, '11:50', '13:20', 6,  66,'2023-05-13','国语2D',24);