CREATE DATABASE IF NOT EXISTS `tiktok` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE tiktok;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for author
-- ----------------------------
DROP TABLE IF EXISTS `author`;
CREATE TABLE `author`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名称',
  `follow_count` bigint NULL DEFAULT NULL COMMENT '关注总数',
  `follower_count` bigint NULL DEFAULT NULL COMMENT '粉丝总数',
  `is_follow` tinyint(1) NULL DEFAULT NULL COMMENT '是否关注(true-已关注，false-未关注)',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户头像',
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户个人页顶部大图',
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '个人简介',
  `total_favorited` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '获赞数量',
  `work_count` bigint NULL DEFAULT NULL COMMENT '作品数',
  `favorite_count` bigint NULL DEFAULT NULL COMMENT '喜欢数',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of author
-- ----------------------------
INSERT INTO `author` VALUES (1, 'Liion', 100, 1000, 1, 'https://cdnjson.com/images/2023/08/03/avator.jpg', 'https://w.wallhaven.cc/full/ex/wallhaven-ex9gwo.png', '小白小白', '10000', 2023, 1024);

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '视频唯一标识ID',
  `author_id` bigint NULL DEFAULT NULL COMMENT '视频作者ID',
  `play_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频播放地址',
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频封面地址',
  `favorite_count` bigint NULL DEFAULT NULL COMMENT '视频的点赞总数',
  `comment_count` bigint NULL DEFAULT NULL COMMENT '视频的评论总数',
  `is_favorite` tinyint(1) NULL DEFAULT NULL COMMENT '是否点赞（true-已点赞，false-未点赞）',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频标题',
  `post_time` datetime NULL DEFAULT NULL COMMENT '视频投稿时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES (1, 1, 'http://vfx.mtime.cn/Video/2021/07/10/mp4/210710094507540173.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 100, 1000, 1, 'Big Buck', '2023-08-22 15:04:25');
INSERT INTO `video` VALUES (2, 1, 'http://www.w3school.com.cn/example/html5/mov_bbb.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 200, 2000, 0, 'Butterfly', '2023-08-23 15:05:50');
INSERT INTO `video` VALUES (3, 1, 'https://media.w3.org/2010/05/sintel/trailer.mp4', 'https://www.helloimg.com/images/2023/08/29/oiEHNM.jpg', 300, 3000, 0, 'Snow Mountain', '2023-08-08 15:06:51');

SET FOREIGN_KEY_CHECKS = 1;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE users (
    user_id          BIGINT AUTO_INCREMENT PRIMARY KEY,
    userName         VARCHAR(32) UNIQUE NOT NULL,
    password         VARCHAR(32)        NOT NULL,
    token            VARCHAR(32)        NOT NULL,
    follow_count     INT     DEFAULT 0,
    follower_count   INT     DEFAULT 0,
    is_follow        BOOLEAN DEFAULT false,
    avatar           VARCHAR(255),
    background_image VARCHAR(255),
    signature        VARCHAR(255),
    total_favorited  VARCHAR(255),
    work_count       INT     DEFAULT 0,
    favorite_count   INT     DEFAULT 0
);

-- ----------------------------
-- 创建注册表单记录表
-- ----------------------------
CREATE TABLE register_forms (
    id       BIGINT AUTO_INCREMENT PRIMARY KEY,
    userName VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(32)        NOT NULL
);


-- ----------------------------
-- 创建登录表单记录表
-- ----------------------------
CREATE TABLE login_forms (
    id       BIGINT AUTO_INCREMENT PRIMARY KEY,
    userName VARCHAR(32) UNIQUE NOT NULL,
    password VARCHAR(32)        NOT NULL
);


