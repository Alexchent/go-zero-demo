CREATE TABLE `policy` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                          `cate` varchar(11) NOT NULL DEFAULT 'pay_mode' COMMENT '类别',
                          `attr` json DEFAULT '' NOT NULL COMMENT '对象信息',
                          `rule` varchar(10240) NOT NULL DEFAULT ''  COMMENT '匹配规则',
                          `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态0关闭 1开启',
                          `created` datetime DEFAULT NOT NULL,
                          `updated` datetime DEFAULT NOT NULL,
                          `deleted` datetime DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 986 DEFAULT CHARSET = utf8