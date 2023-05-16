CREATE TABLE `hall`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `hall_name` varchar(100) NOT NULL DEFAULT '' COMMENT '显示名称',
  `seat_address` varchar(200) NOT NULL DEFAULT '' COMMENT '座位文件存放地址',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '地域信息表';

INSERT INTO `hall` VALUES (1, 'IMAX厅', '/seats/1.json');
INSERT INTO `hall` VALUES (2, 'CGS中国巨幕厅', '/seats/1.json');
INSERT INTO `hall` VALUES (3, '杜比全景声厅', '/seats/1.json');
INSERT INTO `hall` VALUES (4, 'Dolby Cinema厅', '/seats/1.json');
INSERT INTO `hall` VALUES (5, 'RealD厅', '/seats/1.json');
INSERT INTO `hall` VALUES (6, 'RealD 6FL厅', '/seats/1.json');
INSERT INTO `hall` VALUES (7, 'LUXE巨幕厅', '/seats/1.json');
INSERT INTO `hall` VALUES (8, '4DX厅', '/seats/1.json');
INSERT INTO `hall` VALUES (9, 'DTS:X 临境音厅', '/seats/1.json');
INSERT INTO `hall` VALUES (10, '儿童厅', '/seats/1.json');
INSERT INTO `hall` VALUES (11, '4K厅', '/seats/1.json');
INSERT INTO `hall` VALUES (12, '4D厅', '/seats/1.json');
INSERT INTO `hall` VALUES (13, '巨幕厅', '/seats/1.json');
