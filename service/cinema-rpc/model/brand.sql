CREATE TABLE `brand`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `brand_name` varchar(100) NOT NULL DEFAULT '' COMMENT '品牌名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '品牌信息表';

INSERT INTO `brand` VALUES (1, '大地影院');
INSERT INTO `brand` VALUES (2, '万达影城');
INSERT INTO `brand` VALUES (3, '耀莱成龙国际影城');
INSERT INTO `brand` VALUES (4, '保利国际影城');
INSERT INTO `brand` VALUES (5, '博纳国际影城');
INSERT INTO `brand` VALUES (6, '金逸影城');
INSERT INTO `brand` VALUES (7, '中影国际影城');
INSERT INTO `brand` VALUES (8, 'CGV影城');
INSERT INTO `brand` VALUES (9, '橙天嘉禾影城');
INSERT INTO `brand` VALUES (10, '新华国际影城');
INSERT INTO `brand` VALUES (11, '星美国际影城');
INSERT INTO `brand` VALUES (12, '百老汇影城');
INSERT INTO `brand` VALUES (13, 'UME国际影城');
INSERT INTO `brand` VALUES (14, '幸福蓝海国际影城');
INSERT INTO `brand` VALUES (15, '首都电影院');
INSERT INTO `brand` VALUES (16, '华谊兄弟影院');
INSERT INTO `brand` VALUES (17, '卢米埃影城');
INSERT INTO `brand` VALUES (18, '沃美影城');
INSERT INTO `brand` VALUES (19, '美嘉欢乐影城');
INSERT INTO `brand` VALUES (20, '嘉华国际影城');
INSERT INTO `brand` VALUES (21, '17.5影城');
INSERT INTO `brand` VALUES (22, '太平洋电影城');
INSERT INTO `brand` VALUES (23, 'SFC上影影城');
INSERT INTO `brand` VALUES (24, '嘉美国际影城');
INSERT INTO `brand` VALUES (25, '东都影城');
INSERT INTO `brand` VALUES (26, '鲁信影城');
INSERT INTO `brand` VALUES (27, '华影国际影城');
INSERT INTO `brand` VALUES (28, '搜秀影城');
INSERT INTO `brand` VALUES (29, '横店电影城');
INSERT INTO `brand` VALUES (99, '全部');
