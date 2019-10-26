-- ----------------------------
-- Table structure for t_account
-- ----------------------------
DROP TABLE IF EXISTS `t_account`;
CREATE TABLE `t_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `passwd` varchar(255) NOT NULL, -- 暂时用简单的md5加密就好，后期再考虑安全加密
  -- `token` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'token',(token存redis)
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_order
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `userid` int(11) NOT NULL COMMENT '所属用户id',
  `vps` varchar(255),
  `cointype` varchar(32) '主节点币类型',
  `status` int(11) DEFAULT 0 COMMENT '订单状态 0 未付费 1 已付费机器未申请 2 已付费机器已申请主节点未完成 3已完成',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_masternode
-- ----------------------------
DROP TABLE IF EXISTS `t_masternode`;
CREATE TABLE `t_masternode` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主节点唯一id',
  `userid` int(11) NOT NULL COMMENT '所属用户id',
  `vps` varchar(255),
  `cointype` varchar(32) '主节点币类型',
  `status` int(11) DEFAULT 0 COMMENT '0未发布、1已发布',
  `txid` varchar(128) 'txid',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
