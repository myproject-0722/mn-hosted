-- ----------------------------
-- Table structure for t_account
-- ----------------------------
DROP TABLE IF EXISTS `t_account`;
CREATE TABLE `t_account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `passwd` varchar(255) NOT NULL, -- 暂时用简单的md5加密就好，后期再考虑安全加密
  -- `token` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'token',(token存redis)
  `walletaddress` varchar(128) DEFAULT "" COMMENT '钱包地址(暂时不用)',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_order
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `userid` bigint(20) NOT NULL COMMENT '所属用户id',
  `coinname` varchar(32) COMMENT '主节点币名称',
  `mnkey` varchar(32) COMMENT '主节点私钥',
  `timetype` TINYINT COMMENT '支付时间(1-天,2-月,3-年)',
  `price` int(11) COMMENT '付费金额',
  `txid` varchar(128) COMMENT 'txid,用于数据货币支付，暂时不用',
  `isrenew` int(11) DEFAULT 0 COMMENT 'isrenew 0-非 1-是',
  `status` int(11) DEFAULT 0 COMMENT '订单状态 0 未完成 1 已完成',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_masternode
-- ----------------------------
DROP TABLE IF EXISTS `t_coinlist`;
CREATE TABLE `t_coinlist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `coinname` varchar(32) COMMENT '主节点币名称',
  `mnrequired` int(11) DEFAULT 0 COMMENT '主节点质押所需货币',
  `dprice` int(11) DEFAULT 0 COMMENT '托管价格(人民币/按天)',
  `mprice` int(11) DEFAULT 0 COMMENT '托管价格(人民币/按月)',
  `yprice` int(11) DEFAULT 0 COMMENT '托管价格(人民币/按年)',
  `volume` int(11) DEFAULT 0 COMMENT '收益',
  `roi` int(11) DEFAULT 0 COMMENT '收益率',
  `monthlyincome` int(11) DEFAULT 0 COMMENT '当月收益',
  `mnhosted` int(11) DEFAULT 0 COMMENT '托管数',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`coinname`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

INSERT INTO `t_coinlist` VALUES (1,'dash',1000,1,2,3,0,0,0,0,'2019-11-27 11:25:32','2019-11-27 11:25:32');
INSERT INTO `t_coinlist` VALUES (2,'vds',10000,1,2,3,0,0,0,0,'2019-11-27 11:25:32','2019-11-27 11:25:32');
INSERT INTO `t_coinlist` VALUES (3,'snowgem',10000,1,2,3,0,0,0,0,'2019-11-27 11:25:32','2019-11-27 11:25:32');


-- ----------------------------
-- Table structure for t_coinprice
-- ----------------------------
DROP TABLE IF EXISTS `t_coinprice`;
CREATE TABLE `t_coinprice` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `coinname` varchar(32) COMMENT '支付币名称',
  `price` bigint(20) DEFAULT 0 COMMENT '与美元汇率*1000000',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`coinname`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_masternode
-- ----------------------------
DROP TABLE IF EXISTS `t_masternode`;
CREATE TABLE `t_masternode` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主节点唯一id',
  `coinname` varchar(32) NOT NULL COMMENT '主节点币类型',
  `mnkey` varchar(32) NOT NULL COMMENT '主节点私钥',
  `userid` bigint(20) NOT NULL COMMENT '所属用户id',
  `orderid` bigint(20) NOT NULL COMMENT '定单id',
  `vps` varchar(255) COMMENT 'vps ip及端口',
  `dockerid` varchar(128) COMMENT 'dockerid',
  `status` int(11) DEFAULT 0 COMMENT '0未发布、1已发布、2已过期',
  `syncstatus` int(11) DEFAULT 0 COMMENT '0未同步、2已同步',
  `mnstatus` int(11) DEFAULT 0 COMMENT '主节点状态',
  `earn` bigint(20) DEFAULT 0 COMMENT '收益',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `expiretime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '到期时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`coinname`,`mnkey`),
  UNIQUE KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;
