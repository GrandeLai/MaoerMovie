CREATE TABLE `film` (
     `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
     `film_name` varchar(100) NOT NULL DEFAULT '' COMMENT '影片名称',
     `film_en_name` varchar(50) NOT NULL DEFAULT '' COMMENT '影片英文名称',
     `film_type` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '片源类型: 0-2D,1-3D,2-3DIMAX,4-无',
     `film_cover` varchar(200) NOT NULL DEFAULT '' COMMENT '影片主图地址',
     `film_length` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '播放时长，以分钟为单位，不足取整',
     `category_id` bigint NOT NULL DEFAULT  0 COMMENT '影片分类，参照分类表,多个分类以#分割',
     `film_area` varchar(100) NOT NULL DEFAULT '' COMMENT '上映区域',
     `film_time` timestamp(0) NULL DEFAULT NULL COMMENT '影片上映时间',
     `director_id` bigint NULL DEFAULT 0 NULL COMMENT '导演编号',
     `biography` text NOT NULL COMMENT '影片介绍',
     `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
     `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片主表';
