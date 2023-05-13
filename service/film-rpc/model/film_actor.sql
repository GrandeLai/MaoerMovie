CREATE TABLE `film_actor`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `film_id` bigint NOT NULL COMMENT '影片编号,对应mooc_film_t',
  `actor_id` bigint NOT NULL COMMENT '演员编号,对应mooc_actor_t',
  `role_name` varchar(100) NOT NULL DEFAULT '' COMMENT '角色名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影片与演员主表';
