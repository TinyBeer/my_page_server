# my_page_server

一个用于个人资源管理的网页后端服务。计划支持 TODO 和个人影院功能。

## 启动项目

- docker-compose 启动 mysql
  > 注意 my.cnf 需要提前准备，由于 docker 文件映射先于 my.cnf 文件的创建，如果没有提前准备好 my.cnf 文件会导致 my.cnf 作为文件夹被创建。
  > 获取配置文件的小技巧：启动一个临时镜像，使用 docker cp 从中拷贝出所需文件。
- make run 启动项目

## 功能规划

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
