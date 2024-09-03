/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 tb_coin_detail
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_detail`;

CREATE TABLE `tb_coin_detail` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `task_id` int(11) NOT NULL DEFAULT '0' COMMENT '任务id',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT '积分，正数是奖励，负数是惩罚',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# 转储表 tb_coin_task
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_task`;

CREATE TABLE `tb_coin_task` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `task` varchar(255) NOT NULL DEFAULT '' COMMENT '任务名称，必须唯一',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT '积分数，正数是奖励积分，负数是惩罚积分，0需要外部调用传值',
  `limit` int(11) NOT NULL DEFAULT '0' COMMENT '每日限额，默认0不限制',
  `start` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生效开始时间',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `sys_status` int(11) NOT NULL DEFAULT '0' COMMENT '状态，默认0整除，1删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `task` (`task`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# 转储表 tb_coin_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_coin_user`;

CREATE TABLE `tb_coin_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `coins` int(11) NOT NULL DEFAULT '0' COMMENT '总积分',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# 转储表 tb_grade_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_info`;

CREATE TABLE `tb_grade_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '等级名称',
  `description` varchar(3000) NOT NULL COMMENT '等级描述信息',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '等级最高的成长数值',
  `expired` int(11) NOT NULL DEFAULT '0' COMMENT '有效期，单位:天，默认0永不过期',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# 转储表 tb_grade_privilege
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_privilege`;

CREATE TABLE `tb_grade_privilege` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `grade_id` int(11) NOT NULL DEFAULT '0' COMMENT '等级id',
  `product` varchar(255) NOT NULL DEFAULT '' COMMENT '产品',
  `function` varchar(255) NOT NULL DEFAULT '' COMMENT '功能',
  `description` varchar(3000) NOT NULL DEFAULT '' COMMENT '描述信息',
  `expired` int(11) NOT NULL DEFAULT '0' COMMENT '有效期，单位:天，默认0永不过期',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `sys_status` int(11) NOT NULL DEFAULT '0' COMMENT '状态，默认0整除，1删除',
  PRIMARY KEY (`id`),
  KEY `grade_id` (`grade_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# 转储表 tb_grade_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `tb_grade_user`;

CREATE TABLE `tb_grade_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `grade_id` int(11) NOT NULL DEFAULT '0' COMMENT '等级id',
  `expired` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '成长数值',
  `sys_created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `sys_updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
