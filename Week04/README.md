学习笔记

## 工程项目结构

### 常规目录
- /api: API协议定义目录，如proto文件
- /cmd: 可执行文件目录，建立子文件如myapp/main.go
- /internal: 私有应用程序和库代码，外部程序不可用
    - 常见子目录有 model、dao、service、server
    - 如果有多个服务时建议internal下加一层如/internal/app1/biz
- /pkg: 外部程序可用的库代码 
- /configs: 配置文件
- /test: 额外外部测试程序和数据

app服务分类
- interface: 对外的BFF服务，接受来自用户的请求
- service：对内的微服务，仅接受内部服务或网关的请求
- admin：区别于service,面向于运营测的服务，通常数据权限更高，与service共用缓存和DB
- job：流式任务处理，常驻进程
- task: 定时任务，类似cronjob

app目录结构v1
    - model:放对应存储层的结构体，是对如MySQL的表的映射
    - dao: 数据读写层，数据库和缓存在这一层处理
    - service: 组合各种数据访问来构建业务逻辑
    - server: 服务的启动，KIT库可以抽象出该层

app目录v2
    - biz: 业务逻辑的组装层
    - data: 业务数据访问，办好cache、db等封装， 实现了biz的repo接口 
    - service: 实现了api定义的服务层，处理了DTO到biz领域实体的转换(DTO(Data Transfer Object)->DO (Domain Object)
    
### Kit工程
公司统一的一个基础库
- 统一
- 标准库式布局
- 高度抽象
- 支持插件

## API设计

## 配置管理

## 包管理

## 测试

## references

[Layout参考](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)