/*
 Navicat Premium Data Transfer

 Source Server         : xinzf0520.top
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : xinzf0520.top:20607
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 12/06/2022 21:49:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `video_id` int DEFAULT NULL COMMENT '视频id',
  `status` tinyint DEFAULT NULL COMMENT '评论状态',
  `content` text CHARACTER SET utf8mb3 COLLATE utf8_general_ci COMMENT '内容',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_video_id_status` (`video_id`,`status`) USING BTREE COMMENT '获取某个视频的所有评论'
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of comment
-- ----------------------------
BEGIN;
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (13, 11, 11, 1, '恒定', '2022-06-12 13:36:01', '2022-06-12 13:36:01');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (14, 8, 17, 1, '他', '2022-06-12 15:12:32', '2022-06-12 15:12:32');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (15, 8, 16, 2, '你好', '2022-06-12 15:24:01', '2022-06-12 15:24:01');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (16, 8, 17, 1, '真诚', '2022-06-12 17:26:16', '2022-06-12 17:26:16');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (17, 8, 17, 1, '这马克', '2022-06-12 17:26:48', '2022-06-12 17:26:48');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (18, 8, 19, 1, '你好', '2022-06-12 18:02:54', '2022-06-12 18:02:54');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (19, 8, 19, 1, '帝豪', '2022-06-12 18:14:32', '2022-06-12 18:14:32');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (20, 8, 19, 1, '我是不是', '2022-06-12 18:16:09', '2022-06-12 18:16:09');
INSERT INTO `comment` (`id`, `user_id`, `video_id`, `status`, `content`, `created_at`, `updated_at`) VALUES (21, 12, 19, 2, '现在', '2022-06-12 18:29:44', '2022-06-12 18:29:44');
COMMIT;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `video_id` int DEFAULT NULL COMMENT '视频id',
  `status` tinyint DEFAULT NULL COMMENT '点赞状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id_status_video_id` (`user_id`,`status`,`video_id`) USING BTREE COMMENT '获取某个用户所有点赞视频id'
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of favorite
-- ----------------------------
BEGIN;
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (1, 7, 1, 0, '2022-06-11 16:31:44', '2022-06-11 16:31:46');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (2, 6, 1, 1, '2022-06-11 22:25:53', '2022-06-11 22:25:53');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (3, 15, 1, 2, '2022-06-12 11:47:15', '2022-06-12 11:47:35');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (4, 8, 1, 1, '2022-06-12 12:05:41', '2022-06-12 15:00:39');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (5, 11, 1, 1, '2022-06-12 12:22:29', '2022-06-12 12:22:29');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (7, 11, 7, 1, '2022-06-12 12:44:59', '2022-06-12 12:44:59');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (9, 11, 8, 1, '2022-06-12 12:57:09', '2022-06-12 12:57:54');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (10, 11, 5, 2, '2022-06-12 12:58:14', '2022-06-12 12:58:24');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (11, 11, 11, 1, '2022-06-12 13:36:10', '2022-06-12 13:36:10');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (12, 6, 11, 1, '2022-06-12 14:13:02', '2022-06-12 14:13:03');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (13, 6, 10, 1, '2022-06-12 14:14:05', '2022-06-12 14:14:05');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (14, 8, 17, 1, '2022-06-12 15:12:23', '2022-06-12 15:12:23');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (15, 8, 16, 1, '2022-06-12 15:23:55', '2022-06-12 15:23:55');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (16, 7, 17, 2, '2022-06-12 16:11:18', '2022-06-12 16:11:19');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (17, 8, 19, 1, '2022-06-12 18:16:23', '2022-06-12 18:16:23');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (18, 8, 18, 1, '2022-06-12 18:26:54', '2022-06-12 18:26:54');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (19, 12, 19, 1, '2022-06-12 18:30:01', '2022-06-12 18:30:01');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (20, 0, 18, 1, '2022-06-12 20:16:43', '2022-06-12 20:16:43');
INSERT INTO `favorite` (`id`, `user_id`, `video_id`, `status`, `created_at`, `updated_at`) VALUES (21, 0, 17, 1, '2022-06-12 20:16:55', '2022-06-12 20:16:55');
COMMIT;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `followed_user` int DEFAULT NULL COMMENT '被关注者id',
  `status` tinyint DEFAULT '1' COMMENT '关注状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user_id_status_followed_user` (`user_id`,`status`,`followed_user`) USING BTREE COMMENT '查询关注列表',
  KEY `idx_followed_user_status_user_id` (`followed_user`,`status`,`user_id`) USING BTREE COMMENT '查询粉丝列表'
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of follow
-- ----------------------------
BEGIN;
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (1, 6, 7, 1, '2022-06-11 17:05:58', '2022-06-11 17:06:00');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (13, 15, 1, 1, '2022-06-12 11:47:37', '2022-06-12 11:47:37');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (14, 7, 8, 1, '2022-06-12 16:11:20', '2022-06-12 16:11:20');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (15, 8, 12, 1, '2022-06-12 18:26:32', '2022-06-12 18:26:32');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (16, 13, 12, 1, '2022-06-12 18:27:58', '2022-06-12 18:27:58');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (17, 12, 8, 1, '2022-06-12 18:28:50', '2022-06-12 18:28:50');
INSERT INTO `follow` (`id`, `user_id`, `followed_user`, `status`, `created_at`, `updated_at`) VALUES (18, 12, 12, 1, '2022-06-12 18:29:23', '2022-06-12 18:29:23');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '姓名',
  `password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
  `follow_count` int DEFAULT NULL COMMENT '关注数',
  `follower_count` int DEFAULT NULL COMMENT '粉丝数',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE COMMENT '根据用户名查询信息（如注册）'
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (6, 'liufei', '$2a$04$QZBijSqDj70lYUc3bHPkRO4UPDSTlzEdM4ZADBayJEwau0lAikEe.', 0, 0, '2022-06-10 22:28:36');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (7, 'nyf12', '$2a$04$Tzk6xUVGZ/IU.W11/4X8fupOYDRGTWVMHVENiuYf5ymEc0OakcgYG', 1, 0, '2022-06-10 22:35:31');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (8, 'qjh123', '$2a$04$s2zjwbRIwUyh7mKyof7P2uo2Yb2qQBx7xVkE1WcOzWjeMHDu65m7q', 1, 2, '2022-06-11 21:12:47');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (9, 'abnd', '$2a$04$k8Wm4X5WvICUNrrqeexpRuitSyfRGPzK7YHm8ll0GjmbmtrIahedu', 0, 0, '2022-06-11 22:10:25');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (11, 'baize2', '$2a$04$IvCrpJEkJBQjTw3g6Q8hVuXUJyGh95Nm7.Am374.vHRF0kp0nN0Qi', 0, 0, '2022-06-12 12:21:50');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (12, 'wqy', '$2a$04$850HYaW3Nm/0qgLGcge/0uA.gqHJgHN6trxombJrXbpr2hGAvwO8O', 2, 3, '2022-06-12 15:38:19');
INSERT INTO `user` (`id`, `name`, `password`, `follow_count`, `follower_count`, `created_at`) VALUES (13, 'woshi', '$2a$04$zcQ17Toq7WPEPqe72dZyVOJWwmrSIpuz.usW1yokFtC8W/S8d0fMS', 1, 0, '2022-06-12 18:27:41');
COMMIT;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int DEFAULT NULL COMMENT '作者id',
  `play_url` varchar(1024) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '播放路径',
  `cover_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8_general_ci DEFAULT NULL COMMENT '封面路径',
  `favorite_count` int DEFAULT NULL COMMENT '喜欢数',
  `comment_count` int DEFAULT NULL COMMENT '评论数',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `title` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_author_id` (`author_id`) USING BTREE COMMENT '根据用户id查询所有发布的视频'
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of video
-- ----------------------------
BEGIN;
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (8, 6, 'https://video.liufei.fun/sv/2eaa675f-1815638c4ce/2eaa675f-1815638c4ce.mp4', 'https://video.liufei.fun/553d29da2b3844c3bc930e5888f4c5d4/snapshots/216c32eac6064241ada9704637c88101-00002.jpg', 1, 0, '2022-06-12 12:42:58', 'test1');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (9, 6, 'https://video.liufei.fun/sv/5359e249-1815645a2da/5359e249-1815645a2da.mp4', 'https://video.liufei.fun/f4cbfa7bdcf74a5eb3940d716175e607/snapshots/1146837906c74ff89cfcb7ff0f4d58c6-00001.jpg', 0, 0, '2022-06-12 12:57:00', '小姐姐很可爱');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (10, 6, 'https://video.liufei.fun/sv/5a28a-1815646969d/5a28a-1815646969d.mp4', 'https://video.liufei.fun/b2cae3b102fc4de895c947515bfeac77/snapshots/7877b959ebfa4cddad821a9e99e66e71-00001.jpg', 0, 0, '2022-06-12 12:58:01', 'wow！！');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (11, 6, 'https://video.liufei.fun/sv/35fa8c9-1815646ecea/35fa8c9-1815646ecea.mp4', 'https://video.liufei.fun/f4964eb8a6014b6bb4f5910a52727f2b/snapshots/d2fa3c8315cf4c8ba237fb212cbaabfa-00001.jpg', 1, 1, '2022-06-12 12:58:22', '羽毛球');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (13, 8, 'https://video.liufei.fun/sv/32b668dc-18156b6ce05/32b668dc-18156b6ce05.mp4', 'https://video.liufei.fun/0da1facdb7e44b84bf3bbf50b4507e13/snapshots/c2a49cfa8e7647019ea68b802977d6fa-00005.jpg', 0, 0, '2022-06-12 15:00:31', '害');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (14, 8, 'https://video.liufei.fun/sv/43f75d23-18156bb6367/43f75d23-18156bb6367.mp4', 'https://video.liufei.fun/ebb38a03437f411e9035873185c92e0c/snapshots/96a6da7362ae4299ab4a59b58bc659f9-00002.jpg', 0, 0, '2022-06-12 15:05:31', '跳舞');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (15, 8, 'https://video.liufei.fun/sv/f489b4f-18156bba6bf/f489b4f-18156bba6bf.mp4', 'https://video.liufei.fun/af29520eda3f44589b77eeded1f7a72a/snapshots/b63c6028061f437ebfae9c43366ede28-00002.jpg', 0, 0, '2022-06-12 15:05:48', '哈');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (16, 8, 'https://video.liufei.fun/sv/4fba75-18156bbedf7/4fba75-18156bbedf7.mp4', 'https://video.liufei.fun/00695d834896485d8eb03cf4207c4c60/snapshots/e744791b640d4a9ca2f84cf14deb63bf-00001.jpg', 1, 0, '2022-06-12 15:06:06', '的');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (17, 8, 'https://video.liufei.fun/sv/3b73ac16-18156bc7fb0/3b73ac16-18156bc7fb0.mp4', 'https://video.liufei.fun/011f13290385477b8587fba464bb2dcd/snapshots/80ba4089a6e04c198dd0a01014ae7a40-00001.jpg', 2, 3, '2022-06-12 15:06:43', '本科');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (18, 12, 'https://video.liufei.fun/sv/101ec4e2-18156da596a/101ec4e2-18156da596a.mp4', 'https://video.liufei.fun/77aa62b45d1a48789d0727c5486c3981/snapshots/ad766cbc7a4144bca18a784ef03380bc-00001.jpg', 2, 0, '2022-06-12 15:39:20', '真不错');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (19, 12, 'https://video.liufei.fun/sv/5359e249-1815645a2da/5359e249-1815645a2da.mp4', 'https://video.liufei.fun/8789322bdc174a608fd9810199ed09ca/snapshots/d64929a099e048e291508bfef43dad1d-00001.jpg', 2, 3, '2022-06-12 16:42:02', 'nice');
INSERT INTO `video` (`id`, `author_id`, `play_url`, `cover_url`, `favorite_count`, `comment_count`, `created_at`, `title`) VALUES (20, 6, 'https://video.liufei.fun/sv/4ecc930a-18157b35697/4ecc930a-18157b35697.mp4', 'https://video.liufei.fun/sv/4ecc930a-18157b35697/4ecc930a-18157b35697.mp4?x-oss-process=video/snapshot,t_30000,f_jpg,w_0,h_0,m_fast,ar_auto', 0, 0, '2022-06-12 19:39:25', '小礼物');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
