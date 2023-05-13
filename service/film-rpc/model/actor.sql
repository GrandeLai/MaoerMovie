CREATE TABLE `actor`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `actor_name` varchar(50) NOT NULL DEFAULT '' COMMENT '演员名称',
  `actor_img` varchar(200) NOT NULL DEFAULT '' COMMENT '演员图片位置',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片类型信息表';
