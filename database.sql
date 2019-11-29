##### 用户表
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(20) DEFAULT NULL,
  `nick_name` varchar(20) DEFAULT NULL,
  `name` varchar(10) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `trade_pwd` varchar(255) DEFAULT NULL,
  `login_pwd` varchar(255) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `email` varchar(30) DEFAULT NULL,
  `level` char(1) DEFAULT NULL,
  `kyc_level` char(1) DEFAULT NULL,
  `identity_card` varchar(18) DEFAULT NULL,
  `card_type` int(1) DEFAULT NULL,
  `last_login_at` datetime DEFAULT NULL,
  `last_login_ip` varchar(30) DEFAULT NULL,
  `create_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  `country` int(5) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;



##### 订单表
CREATE TABLE `order` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_no` varchar(30) DEFAULT NULL COMMENT '订单号',
  `user_id` int(11) DEFAULT NULL COMMENT '用户ID',
  `symbol` varchar(255) DEFAULT NULL COMMENT '交易对',
  `type` varchar(2) DEFAULT NULL COMMENT '订单类型',
  `state` int(1) DEFAULT NULL COMMENT '订单状态',
  `deposit` decimal(24,12) DEFAULT NULL COMMENT '保证金',
  `price` decimal(24,12) DEFAULT NULL COMMENT '委托价',
  `amount` decimal(24,12) DEFAULT NULL COMMENT '委托量',
  `volume` decimal(24,12) DEFAULT NULL COMMENT '成交量',
  `avg_price` decimal(24,12) DEFAULT NULL COMMENT '成交均价',
  `hold_amount` decimal(24,12) DEFAULT NULL COMMENT '持仓量',
  `received` decimal(24,12) DEFAULT NULL COMMENT '已收到',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '更新时间',
  `sn` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '订单签名',
  `deal_at` datetime DEFAULT NULL COMMENT '成交时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8;


##### 账户表
CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `currency_id` int(11) DEFAULT NULL COMMENT '币种ID',
  `currency` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '币种名称',
  `balance` decimal(24,12) DEFAULT NULL COMMENT '当前余额',
  `locked` decimal(24,12) DEFAULT NULL COMMENT '冻结',
  `type` int(1) DEFAULT NULL COMMENT '类型 0-真实资金 1-虚拟资金',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '修改时间',
  `sn` varchar(255) DEFAULT NULL COMMENT '加密属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;


##### 授权登录表
CREATE TABLE `auth_login_address` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `address` varchar(30) DEFAULT NULL COMMENT '地址名称',
  `ip_address` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT 'ip',
  `state` int(1) DEFAULT NULL COMMENT '状态 0-未确认 1-已确认',
  `login_type` varchar(10) DEFAULT NULL COMMENT '设备类型',
  `login_at` datetime DEFAULT NULL COMMENT '登录时间',
  `create_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;