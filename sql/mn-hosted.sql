-- ----------------------------
-- Table structure for t_account
-- ----------------------------
DROP TABLE IF EXISTS `t_account`;
CREATE TABLE `t_account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account` varchar(255) NOT NULL,
  `passwd` varchar(255) NOT NULL, -- 暂时用简单的md5加密就好，后期再考虑安全加密
  -- `token` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'token',(token存redis)
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_order
-- ----------------------------
DROP TABLE IF EXISTS `t_order`;
CREATE TABLE `t_order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单id',
  `userid` bigint(20) NOT NULL COMMENT '所属用户id',
  `vps` varchar(255),
  `coinname` varchar(32) COMMENT '主节点币名称',
  `status` int(11) DEFAULT 0 COMMENT '订单状态 0 未付费 1 已付费机器未申请 2 已付费机器已申请主节点未完成 3已完成',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;


-- ----------------------------
-- Table structure for t_masternode
-- ----------------------------
DROP TABLE IF EXISTS `t_coinlist`;
CREATE TABLE `t_coinlist` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `coinname` varchar(32) COMMENT '主节点币名称',
  `mnrequired` int(11) DEFAULT 0 COMMENT '主节点质押所需货币',
  `mnprice` int(11) DEFAULT 0 COMMENT '主节点托管所需货币(分)',
  `volume` int(11) DEFAULT 0 COMMENT '收益',
  `roi` int(11) DEFAULT 0 COMMENT '收益率',
  `monthlyincome` int(11) DEFAULT 0 COMMENT '当月收益',
  `mnhosted` int(11) DEFAULT 0 COMMENT '托管数',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`coinname`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for t_masternode
-- ----------------------------
DROP TABLE IF EXISTS `t_masternode`;
CREATE TABLE `t_masternode` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主节点唯一id',
  `coinname` varchar(32) COMMENT '主节点币类型',
  `mnkey` varchar(32) COMMENT '主节点私钥',
  `userid` bigint(20) NOT NULL COMMENT '所属用户id',
  `vps` varchar(255),
  `dockerid` varchar(128) COMMENT 'dockerid',
  `status` int(11) DEFAULT 0 COMMENT '0未发布、1已发布',
  `syncstatus` int(11) DEFAULT 0 COMMENT '0未同步、2已同步',
  `mnstatus` varchar(32) COMMENT '主节点状态',
  `createtime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `expiretime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '到期时间',
  `updatetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`coinname`,`mnkey`),
  UNIQUE KEY (`id`),
  KEY (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4;
