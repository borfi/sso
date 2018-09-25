CREATE TABLE `users` (
  `user_id` bigint(30) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `user_account` varchar(60) NOT NULL DEFAULT '' COMMENT '用户账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `type` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '类型 0普通用户 1管理员',
  `country` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '国家或地区 0大陆 1香港 2澳门 3台湾',
  `phone` varchar(255) NOT NULL DEFAULT '' COMMENT '手机号码',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `real_name` varchar(255) NOT NULL DEFAULT '' COMMENT '真实姓名',
  `net_name` varchar(255) NOT NULL DEFAULT '' COMMENT '网名',
  `head_img` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  `sex` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '性别 0女 1男',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0未激活 1已激活 2锁定 3删除',
  `add_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `last_login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次登陆时间',
  `last_login_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '最后一次登陆IP',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_account` (`user_account`) USING BTREE,
  KEY `phone` (`phone`) USING BTREE,
  KEY `add_time` (`add_time`) USING BTREE,
  KEY `password` (`password`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

