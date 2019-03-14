/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost:3306
 Source Schema         : spider

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : 65001

 Date: 14/03/2019 21:46:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cartoon_chapter
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_chapter`;
CREATE TABLE `cartoon_chapter` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `unique_id` char(32) NOT NULL COMMENT '章节唯一标识',
  `list_unique_id` char(32) NOT NULL COMMENT '漫画列表 唯一标识',
  `resource_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '章节名称',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `conent` text COMMENT '内容',
  `is_free` tinyint(1) DEFAULT '0' COMMENT '是否收费 1收费 0免费',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1抓取完毕，0未进行抓取，-1重新抓取',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '章节列表url',
  `resource_img_url` varchar(255) DEFAULT NULL COMMENT '图片资源',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍章节抓取时间',
  `resource_no` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=496936 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for cartoon_chapter_content
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_chapter_content`;
CREATE TABLE `cartoon_chapter_content` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `resource_no` int(11) DEFAULT NULL,
  `list_unique_id` char(32) NOT NULL COMMENT '漫画列表 唯一标识',
  `chapter_unique_id` char(32) NOT NULL COMMENT '章节唯一标识',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '外部资源url',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍章节资源抓取时间',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for cartoon_list
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_list`;
CREATE TABLE `cartoon_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `unique_id` char(32) NOT NULL COMMENT '章节唯一标识',
  `resource_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '书籍名称',
  `tags` varchar(100) DEFAULT NULL COMMENT '标签',
  `author` varchar(50) DEFAULT NULL COMMENT '作者',
  `detail` varchar(200) DEFAULT NULL COMMENT '详情',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态 1抓取完毕，0未进行抓取，-1重新抓取',
  `resource_url` varchar(255) DEFAULT NULL COMMENT '网站url资源',
  `resource_name` varchar(50) DEFAULT NULL COMMENT '资源名称',
  `resource_img_url` varchar(255) DEFAULT NULL COMMENT '图片资源',
  `download_img_url` varchar(255) DEFAULT NULL COMMENT '下载 保存的图片地址',
  `cdate` int(11) DEFAULT NULL COMMENT '书籍抓取时间',
  `is_free` tinyint(1) DEFAULT NULL COMMENT '是否免费1免费 0收费',
  `is_end` tinyint(1) DEFAULT NULL COMMENT '是否完结 1完结 0未完结',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6484 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for cartoon_resource
-- ----------------------------
DROP TABLE IF EXISTS `cartoon_resource`;
CREATE TABLE `cartoon_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `resource_url` varchar(255) DEFAULT NULL COMMENT 'url资源',
  `resource_name` varchar(50) DEFAULT NULL COMMENT '漫画资源名称',
  `config_name` varchar(50) DEFAULT NULL COMMENT '配置文件名称',
  `resource_no` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=6484 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
