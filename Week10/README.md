第十周: 日志、指标和链路追踪
---

## 日志
### 日志级别
- warning: 一般不会有人处理，所以应该将warning当成error，或者不用warning级别
- fatal： 记录消息后直接调用os.Exit(1)，其他goroutine的defer不会被调用，各种buffer不会被flush，不应该用fatal
- error: 处理方式
  - wrap后往上抛
  - 服务降级，打印warning级别日志，调用planB
- debug: 面向开发者

### 日志选型

#### 需求
- 收集
- 传输
- 分析
- 警告

#### 设计目标
- 接入方式收敛
- 日志格式规范
- 系统高吞吐低延迟
- 高可用、可扩容

#### 格式规范
使用json作为日主输出格式
必要字段
- time: 日志产生时间，ISO8601格式
-  level: 日志等级，如INFO
- app_id: 应用ID，用于表示日志来源
- instance_id: 示例ID, 用于区分同一应用不同实例，如可用hostname

#### B站实现
##### 采集
- logstash: 监听tcp/udp，适用于网络上报日志
- filebeat: 采集日志文件

#####  传输
flink --> kafaka --> logstash

##### 切分
logstash做filer 

后期会改为flink

##### 存储和检索
elasticsearch多集群架构

单集群内
- master node
- data node(分冷热数据，热数据存在SSD)
- client node
每天固定时间进行热数据迁移到冷数据
index提前一天创建



#### 开源组件ELK：
- Logstash：收集app通过网络发送过来的日志（因为是server角色，可能出现热点）
- Elasticsearch：存储和搜索
- Kibana: UI展示

优化
1. logstash agent先将数据写入kafka,数据可以站在在kafka, 避免logstash server故障导致数据丢失
2. 将logstash替换为filebeat，更灵活和资源消耗更小

## 链路追踪

### 设计目标
- 无处不在的部署
- 持续的监控
- 低消耗
- 应用级透明
- 延展性
- 低延迟

### Goole Dapper论文核心思想
每个请求都有一个全局的traceid，一路透传，每一层都生成一个spanid,并记录上游spanid，即可展现调用链关系

traceid: 使用city hash + uuid

- Tree
- Span
- Annotation: 备注服务名和调用函数

### 跟踪采样
- 固定采用：1/1024
- 应对积极采样：如定个目标1s采集5条，使用滑动窗口

### 经验和优化
- 性能优化： 串行调用该并行、批量查询/写入


## 监控

### 4个黄金指标
- 延迟 (耗时，需要区分正常还是异常)
- 流量
- 错误 （覆盖错误码或者HTTP Status Code）
- 饱和度 (服务容量有多满)

### 系统层面
- CPU、Memory、IO、NetWork、Kernel ContextSwitch
- Runtime: 各类GC

### 统计
Prometheus(拉模式)+ Granfana

### 运行时分析
- 线上打开Profiling端口，或者进一步提供web界面查看Profiling信息
- watchdog:使用滑动窗口方式，当内存或CPU等信号量自动触发profiling采集并存储

## reference
