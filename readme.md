# my_page_server

一个用于个人资源管理的网页后端服务。计划支持备忘录和个人影院功能。
目前仅提供简易前端页面访问服务，前端后序会逐步迁移到 [my_page_view](https://github.com/TinyBeer/my_page_view) 项目中。

## 启动项目

- docker-compose 启动 mysql
  > 注意 my.cnf 需要提前准备，由于 docker 文件映射先于 my.cnf 文件的创建，如果没有提前准备好 my.cnf 文件会导致 my.cnf 作为文件夹被创建。
  > 获取配置文件的小技巧：启动一个临时镜像，使用 docker cp 从中拷贝出所需文件。

<!-- todo 项目启动脚本准备 -->

## 功能规划

- [x] 项目启动脚本
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
- mongo:4.2 数据存储
