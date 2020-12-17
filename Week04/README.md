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

### GRPC
#### 特性和优点
- 支持多种语言
- 轻量级，高性能
- IDL 基于文件定义服务
- 基于HTTP2设计，支持双向流，消息压缩，单TCP多路复用，服务端推送等
### 建立一个专门的仓库存放proto文件，加git hook自动从相关仓库同步过来
- 方便跨部门合作
- 版本管理
- 规范化检查

### Compatibility
- 非破坏性向后兼容
  - 加接口
  - 加字段
- 破坏性向后兼容：升级major版本号

### Naming Conventions
- url: package_name.version.service.name.method
- package: package_name.version

### Primitive Fields
区分默认值还是就是默认值：可以将自动包在额外的Message里面，为nil则为默认值

### Errors
- 不使用公司自定义全局错误码
- 使用标准的GRPC错误码
- 具体的错误放在Details里面
- service error(服务端) --> grpc error(传输时转换成grpc标准错误) --> service error(客户端翻译)


### Update接口设置
- 所有字段共用一个接口，而非一个字段一个接口
- google.protobuf.FieldMask可表示那个字段需要更新，默认更新所有字段

## 配置管理

### 四种配置
- 环境变量配置：Region、Cluster、Env、Color、Discovery、AppID等信息由平台在运行时打入容器环境变量
- 静态配置：http/gRPC server、redis、mysql等
- 动态配置：在线的开关来控制简单的策略，会频繁调整和使用
- 全局配置：使用全局配置模板来定制化常用组件

### 实践建议
- 配置初始化和对象初始化解耦
- 区分可选和必选

## 包管理
mod： 使用https://github.com/gomods/athens加速

## 测试
- 测试类型
  - 小型测试： Unit test
  - 中型测试: 集成测试
  - 大型测试：端到端测试
- 测试标准
  - 不同类型的项目有不同的需求
  - kit库需要大量单元测试
  - 业务型项目更多大中型测试
- 单元测试基本要求
 - 快速
 - 环境一致，不污染坏境
 - 可并行，不要求顺序
 
- 实践：使用go官方的Subtest + Gomock完成整个单元测试
  - api: 适合集成测试，直接测试api,使用yape等API测试框架，维护大量业务测试case
  - data: 使用docker compose组织数据库/中间件来模拟真实环境，teardown时清空数据即可
  - biz: 依赖repo、rpc client，利用gomock模拟interface实现，进行业务单元测试
  - service: 依赖biz的实现，构建biz实现类传入，进行单元测试
  
  
## references

- [Layout参考](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)

- [google api design guide](https://cloud.google.com/apis/design/?hl=zh-cn)