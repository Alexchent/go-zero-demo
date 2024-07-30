CREATE TABLE `policy` (
                          `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                          `cate` varchar(11) NOT NULL COMMENT '类别',
                          `attr` json DEFAULT NULL COMMENT '对象信息',
                          `rule` json DEFAULT NULL COMMENT '匹配规则',
                          `state` tinyint(4) DEFAULT '0' COMMENT '状态0关闭 1开启',
                          `created` datetime DEFAULT NULL,
                          `updated` datetime DEFAULT NULL,
                          `deleted` datetime DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 986 DEFAULT CHARSET = utf8