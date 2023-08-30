-- 创建用户表
-- mysql -h 172.16.32.106 -P 53022 -u root -p
-- kRqqJTjJ  
CREATE TABLE users (

    user_id BIGINT AUTO_INCREMENT PRIMARY KEY,

    userName VARCHAR(32) UNIQUE NOT NULL,

    password VARCHAR(32) NOT NULL,

    token VARCHAR(32) NOT NULL,

    follow_count INT DEFAULT 0,

    follower_count INT DEFAULT 0,

    is_follow BOOLEAN DEFAULT false,

    avatar VARCHAR(255),

    background_image VARCHAR(255),

    signature VARCHAR(255),

    total_favorited VARCHAR(255),

    work_count INT DEFAULT 0,

    favorite_count INT DEFAULT 0

);



-- 创建注册表单记录表

CREATE TABLE register_forms (

    id BIGINT AUTO_INCREMENT PRIMARY KEY,

    userName VARCHAR(32) UNIQUE NOT NULL,

    password VARCHAR(32) NOT NULL

);



-- 创建登录表单记录表

CREATE TABLE login_forms (

    id BIGINT AUTO_INCREMENT PRIMARY KEY,

    userName VARCHAR(32) UNIQUE NOT NULL,

    password VARCHAR(32) NOT NULL

);

