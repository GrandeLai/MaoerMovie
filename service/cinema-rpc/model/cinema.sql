CREATE TABLE `cinema`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `cinema_name` varchar(50) NOT NULL DEFAULT '' COMMENT '影院名称',
  `cinema_phone` varchar(50) NOT NULL DEFAULT '' COMMENT '影院电话',
  `brand_id` bigint NOT NULL DEFAULT 0 COMMENT '品牌编号',
  `district_id` bigint NOT NULL DEFAULT 0 COMMENT '区域编号',
  `hall_ids` varchar(200) NOT NULL DEFAULT '' COMMENT '包含的影厅类型,以/作为分割',
  `cinema_imgs` varchar(500) NOT NULL DEFAULT '' NULL COMMENT '影院图片地址',
  `address` varchar(200) NOT NULL DEFAULT '' COMMENT '影院地址',
  `min_price` bigint NOT NULL DEFAULT 0 COMMENT '最低票价',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '影院信息表';

INSERT INTO `cinema` VALUES (1, '大地影院(顺义店)', '18500003333', 1, 1, '/1/3/5/6', '/cinema/30445282__5675168.jpg', '北京市顺义区华联金街购物中心', 60);
INSERT INTO `cinema` VALUES (2, '大地影院(中关村店)', '010-58391939', 1, 2, '/1/2/3/4', '/cinema/30445282__5675168.jpg', '北京市中关村海龙大厦', 60);
INSERT INTO `cinema` VALUES (3, '万达影院(大屯店)', '010-58391939', 2, 3, '/5/6/7/8', '/cinema/44374823__5777386.jpg', '北京市朝阳区大屯路50号金街商场', 60);
INSERT INTO `cinema` VALUES (4, '万达影院(奥体中心店)', '010-58391231', 2, 4, '/1/3/5/6', '/cinema/44374823__5777386.jpg', '北京市朝阳区奥林匹克公园新奥购物中心', 60);
INSERT INTO `cinema` VALUES (5, '万达影院(中南海店)', '010-58398521', 2, 5, '/1/5/7/8', '/cinema/44374823__5777386.jpg', '北京市东城区中南海52号', 60);
INSERT INTO `cinema` VALUES (6, '万达影院(国贸店)', '010-96385274', 2, 6, '/1/2/3/7', '/cinema/5_0805163047.jpg', '北京市朝阳区国贸CBD核心商场5012', 60);
INSERT INTO `cinema` VALUES (7, '慕课影院(大屯店)', '010-98765432', 3, 7, '/1/5/8/9', '/cinema/5_0805163047.jpg', '北京市朝阳区大屯路50号金街商场', 60);
