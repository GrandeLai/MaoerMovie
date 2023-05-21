CREATE TABLE `order`  (
    `uuid` VARCHAR(100) NOT NULL  COMMENT '主键编号',
    `cinema_id` bigint NOT NULL DEFAULT 0 COMMENT '影院编号',
    `show_id` bigint NOT NULL DEFAULT 0 COMMENT '放映场次编号',
    `film_id` bigint NOT NULL DEFAULT 0 COMMENT '电影编号',
    `seats_ids` varchar(50) NOT NULL DEFAULT '' COMMENT '已售座位编号',
    `seats_position` varchar(200) NOT NULL DEFAULT '' COMMENT '已售座位位置，例如：1/2表示1排2列，多个座位用","隔开',
    `price` double NOT NULL DEFAULT 0.0 COMMENT '订单总金额',
    `order_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) COMMENT '下单时间',
    `user_id` bigint NOT NULL DEFAULT 0 COMMENT '下单人',
    `status` int NULL DEFAULT 0 COMMENT '0-待支付,1-已支付,2-已关闭',
    PRIMARY KEY (`uuid`) USING BTREE,
    KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT = '订单信息表';


