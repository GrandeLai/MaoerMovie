CREATE TABLE `user` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户姓名',
    `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户性别，1表示男，2表示女',
    `mobile` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户电话',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `email` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户邮箱',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '用户状态，0表示用户可用，1表示用户已删除',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_mobile_unique` (`mobile`),
    UNIQUE KEY `idx_email_unique` (`email`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
