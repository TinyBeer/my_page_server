# my_page_server

一个用于个人资源管理的网页后端服务。计划支持备忘录和个人影院功能。
目前仅提供简易前端页面访问服务，前端后序会逐步迁移到 [my_page_view](https://github.com/TinyBeer/my_page_view) 项目中。

## 启动项目

- docker-compose 启动 mysql
  > 注意 my.cnf 需要提前准备，由于 docker 文件映射先于 my.cnf 文件的创建，如果没有提前准备好 my.cnf 文件会导致 my.cnf 作为文件夹被创建。
  > 获取配置文件的小技巧：启动一个临时镜像，使用 docker cp 从中拷贝出所需文件。

<!-- todo 项目启动脚本准备 -->

## 功能规划

- [ ] 项目启动脚本
- [ ] 个人影院
- [x] 个人备忘录 初步实现
- [x] 备忘录添加完成确认功能

## 依赖

- gin web 服务
- 基于 jwt 的双 token 认证
- crypto 密码脱敏
- gorm 数据访问
- wire 依赖注入
- swagger api 接口文档
- viper 配置管理
- mysql:5.7 数据存储

## 项目结构

```sh
my_page_server
│  go.mod
│  go.sum
│  main.go
│  readme.md
│  wire.go
│  wire_gen.go
│
├─config  # 配置相关
│      config.go
│      config.yaml
│
├─database # 数据库相连接
│      mysql.go
│
├─delivery # 对外数据接口
│  │  inter.go
│  │  router.go
│  │  router_page.go
│  │  web.go
│  │
│  ├─handler  # gin handler
│  │      memo.go
│  │      user.go
│  │      video.go
│  │
│  └─middleware # 中间件
│          jwt.go
│
├─docs  # swagger 文档
│      docs.go
│      swagger.json
│      swagger.yaml
│
├─html  # 前端页面
│      load.html
│      memo.html
│      movie.html
│      nav.html
│      video.html
│
├─model # 各层级数据模型
│      delivery.go
│      repository.go
│      usecase.go
│
├─repository # 数据访问层
│      inter.go
│      memo.go
│      user.go
│
├─static # 静态文件
│  ├─css
│  │      load.css
│  │      login.css
│  │      memo.css
│  │      movie.css
│  │      nav.css
│  │      reset.css
│  │      video.css
│  │
│  ├─icon
│  │      favicon16.ico
│  │      favicon48.ico
│  │
│  ├─img
│  │  │  fastx.jpg
│  │  │  god_father.jpg
│  │  │  legend.jpg
│  │  │  lion.jpg
│  │  │  old_fox.jpg
│  │  │  pianist.jpg
│  │  │
│  │  └─memo
│  │          0.png
│  │          1.png
│  │          2.png
│  │          3.png
│  │          4.png
│  │          5.png
│  │          6.png
│  │          7.png
│  │
│  ├─js
│  │      auth.js
│  │      memo.js
│  │      nav.js
│  │      video.js
│  │
│  └─video
│      └─movie
│              fastx.mp4
│              god_father.mp4
│              legend.mp4
│              lion.mp4
│              old_fox.mp4
│              pianist.mp4
│
└─usecase # 业务层
    │  inter.go
    │  memo.go
    │  user.go
    │
    └─auth
            auth.go
```
