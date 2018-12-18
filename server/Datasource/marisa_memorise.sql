/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : marisa_memorise

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 18/12/2018 12:02:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for memorise
-- ----------------------------
DROP TABLE IF EXISTS `memorise`;
CREATE TABLE `memorise`  (
  `memoryId` int(11) NOT NULL AUTO_INCREMENT,
  `ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `keyword` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `answer` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`memoryId`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
