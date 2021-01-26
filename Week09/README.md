第九周: 网络编程 
---

## 网络通信协议
### TCP
### UDP
### HTTP 1.1/2/3

## Go网络编程基础
## Goim长连接编程
### 组件介绍
- comet: 长连接管理层，主要监控网络端口，并且通过设备ID进行channel绑定，以及实现了Room
- logic: 逻辑层，监控连接connect/disconnect事件，鉴权，消息路由
- job: 通过消息队列进行消息推送削峰处理
### 边缘节点
comet作为边缘节点部署在各地，和logic通过云vpc专线通信

### 负载均衡
同一片区有多个comet节点，comet节点会上报当前负载等信息，终端会先请求拿到comet列表，然后逐个尝试建立长连接

### 心跳保活
- 自适应心跳时间
  - 区间[min=60s，max=300s]
  - 步长step=30s
  - 周期探测：success = current -step, fail = current - step
  
### room 
使用bucket打散，提高并发度

### 消息推送-推拉结合
- 推模式： 有新消息是服务器主动推给客户端
- 拉模式： 客户端主动发起拉取消息的请求
- 推拉结合模式： 有新消息实时通知，客户端在进行新消息的同步

### 读写扩散
- 读扩散： 只用写自己的信箱，而需要读所有人的信箱
- 写扩散： 只用读自己的信箱，而需要写所有人的信箱

### 唯一ID设计

- UUID
- Snowflake: 时间戳(41bit)+机器ID(10bit)+序列化(12bit)
- 基于步长向DB申请
- 基于redis或DB自增方式申请

## IM私信系统

## 其他

### 常用诊断命令
- nload: 看网卡实时流入流出流量
- tcpflow：抓包
- ss
- netstat
- nmon
- top
- vmstat
- iostat
- iotop
- strace -p $pid：追踪系统调用
- perf top: 看cpu看内存


## reference

- [微信递增消息ID](http://www.52im.net/thread-1998-1-1.html)