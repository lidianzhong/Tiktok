## 极简版抖音

### Tiktok项目结构
```
├── config
│   ├── app.go              应用配置
│   └── database.go         数据库配置
│
├── controller              Controller层
│   ├── feed.go
│   ├── publish_video.go
│   ├── publish_video_test.go
│   └── user.go
│
├── dao                     Dao层
│   ├── db_init.go
│   ├── user.go
│   └── video.go
│
├── middleware              中间件层
│   └── middle.go
│
├── model                   模型
│   ├── author.go
│   ├── user.go
│   └── video.go
│
├── router                  路由
│   └── baseRouters.go
│
├── service                 Service层
│   ├── feed.go
│   ├── publish_video.go
│   └── user.go
│
├── static                  静态文件
│   ├── 0_1693.mp4
│   └── 0_1693.png
│
├── main.go
├── tiktok_create.sql
