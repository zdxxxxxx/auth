DROP TABLE IF EXISTS `mfw_auth_operations`;
CREATE TABLE `mfw_auth_operations` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `value` varchar(100) NOT NULL UNIQUE COMMENT '值',
  `name` varchar(100) NOT NULL UNIQUE COMMENT '名称',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='操作表';