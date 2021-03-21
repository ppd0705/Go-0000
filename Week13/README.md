第十二周：Kafaka
---

## 基础概念
使用场景
- 行为事件的数据管道
- 实时计算的管道
- 业务消息系统

设计目标
- O(1)复杂度级别提供消息持久化能力
- 高吞吐率
- 支持Server间的消息分区，同时保持Partitionn内消息顺序传输
- 支持离线数据处理
- 在线水平扩展


## Topic & Partition
- 一个Topic可能有多个Partition
- Partition内的消息保证有序性
- Broker数量应该不小于Partition数量

存储原理
- 消息存于文件系统之上
- 利用顺序IO和PageCache达成高吞吐
- 保留所有发布的message,物理是否被消费过，提供策略删除旧数据

offset: 偏移量，相当于当前分区第一条消息的偏移量

文件：
- .index, 索引文件使用稀疏索引，二分查找到最近的消息，
- .timeindex
- .log

## Producer & Consumer
producer发送消息到broker,会根据Partition分发机制发送到指定partition
consumer: consumer不应该比分区数多，剩余的消费者会空闲
consumer group: 每个group都可以消费到全部消息
consumer offset： （offset会发送到一个内部的topic以持久化）
- 自动提交，默认5秒
## Leader & Follower
ISR节点集合

High Watermark  && Log End Offset

## 数据可靠性

required.acks
- 0: broker收到回ack
- 1: leader收到后回ack
- -1：flower收到后回ack

min.insync.replicas >= 2


## 性能优化

架构层面
- partition级别并行： broker、disk、consumer
- ISR

IO层面
- batch读写
- 磁盘顺序IO
- page cache
- zero copy
- 压缩