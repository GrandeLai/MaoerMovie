CREATE TABLE `pay` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `pay_sn` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '支付流水号',
    `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
    `order_id` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '订单ID',
    `buyer_account` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '下单用户账号',
    `price` double NOT NULL DEFAULT 0.0 COMMENT '支付总金额',
    `subject` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '支付标题',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '支付状态，0表示未支付，1表示已支付，2表示支付失效，默认为0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_pay_sn` (`pay_sn`),
    KEY `idx_order_id` (`order_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
