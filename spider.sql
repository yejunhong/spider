/*
Navicat MySQL Data Transfer

Source Server         : 202.43.91.26
Source Server Version : 50553
Source Host           : 202.43.91.26:3306
Source Database       : caiji

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-04-01 17:42:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for cartoon_chapter
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_chapter`;
CREATE TABLE `cartoon_chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_no` char(4) DEFAULT NULL COMMENT '漫画编号',
  `unique_id` char(32) NOT NULL COMMENT '章节唯一标识',
  `list_unique_id` char(32) NOT NULL COMMENT '漫画列表 唯一标识',
  `conent` text COMMENT '内容',
  `is_free` tinyint(1) DEFAULT '0' COMMENT '是否收费 1收费 0免费',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1抓取完毕，0未进行抓取，-1重新抓取',
  `resource_name` varchar(100) DEFAULT NULL COMMENT '章节名称',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '章节列表url',
  `resource_img_url` varchar(255) DEFAULT NULL COMMENT '图片资源',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍章节抓取时间',
  `book_type` tinyint(1) DEFAULT '1' COMMENT '数据类型 1漫画 2小说',
  PRIMARY KEY (`id`,`unique_id`),
  UNIQUE KEY `unique_id` (`unique_id`)
) ENGINE=MyISAM AUTO_INCREMENT=668948 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for cartoon_chapter_content
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_chapter_content`;
CREATE TABLE `cartoon_chapter_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `resource_no` char(4) DEFAULT NULL COMMENT '所属漫画编号',
  `list_unique_id` char(32) NOT NULL COMMENT '漫画列表 唯一标识',
  `chapter_unique_id` char(32) NOT NULL COMMENT '章节唯一标识',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '外部资源url',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍章节资源抓取时间',
  PRIMARY KEY (`id`),
  KEY `list_unique_id` (`list_unique_id`,`chapter_unique_id`)
) ENGINE=MyISAM AUTO_INCREMENT=968083 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for cartoon_list
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_list`;
CREATE TABLE `cartoon_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cartoon_id` int(11) DEFAULT NULL COMMENT '漫画id',
  `resource_no` char(4) DEFAULT NULL COMMENT '资源编号',
  `unique_id` char(32) NOT NULL COMMENT '章节唯一标识 md5(书籍名称+站点名称)',
  `tags` varchar(100) DEFAULT NULL COMMENT '标签',
  `author` varchar(100) DEFAULT NULL COMMENT '作者',
  `detail` varchar(255) DEFAULT NULL COMMENT '详情',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1抓取完毕，0未进行抓取，-1重新抓取',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '网站url资源',
  `resource_name` varchar(100) DEFAULT NULL COMMENT '资源名称',
  `resource_img_url` varchar(255) DEFAULT NULL COMMENT '图片资源',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍抓取时间',
  `is_free` tinyint(4) DEFAULT '0' COMMENT '是否付费 1付费 0免费',
  `is_end` tinyint(4) DEFAULT '0' COMMENT '是否完结 1完结 0未完结',
  `chapter_count` int(10) DEFAULT '0' COMMENT '章节数量',
  `book_type` tinyint(1) DEFAULT '1' COMMENT '数据类型 1漫画 2小说',
  PRIMARY KEY (`id`,`unique_id`),
  UNIQUE KEY `unique_id` (`unique_id`)
) ENGINE=MyISAM AUTO_INCREMENT=11257 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for cartoon_resource
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_resource`;
CREATE TABLE `cartoon_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_url` varchar(255) DEFAULT NULL COMMENT 'url资源',
  `resource_name` varchar(50) DEFAULT NULL COMMENT '漫画资源名称',
  `config_name` varchar(50) DEFAULT NULL COMMENT '配置文件名称',
  `resource_no` char(4) DEFAULT NULL COMMENT '资源编号',
  `book_type` tinyint(1) DEFAULT '1' COMMENT '数据类型 1漫画 2小说',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
