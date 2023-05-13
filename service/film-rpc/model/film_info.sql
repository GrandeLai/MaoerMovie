CREATE TABLE `film_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `film_id` bigint NOT NULL COMMENT '电影编号',
  `film_preSaleNum` bigint NOT NULL DEFAULT '0' COMMENT  '影片预售数量',
  `film_box_office` int(10) NOT NULL DEFAULT '0' COMMENT '影片票房：每日更新，以万为单位',
  `film_imgs` text NOT NULL COMMENT '影片图片集地址,多个图片以逗号分隔',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `film_id`(`film_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片详情表';


