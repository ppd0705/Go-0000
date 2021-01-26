学习笔记

## 功能模块

- 变更功能： 添加记录、删除记录、清空历史
- 读取功能：按timeline返回topN
- 其他功能： 记录是否首次观看

历史记录是一个极高tps写入、极高qps读取的业务服务

## 架构设计
BFF -> history-service -> kafka -> history-job -> HBase

BFF切分
- 重要性
- 是否是垂直业务
- 流量大小
- 迭代频率


使用Write-back的思路，先将数据写入分布式缓存，再写回数据库

使用redis-cluster做缓存，有概率丢数据

redis会保存每个用户一定数据，写入会做截断判断如保存500条

HBase会做ttl和条数截断

批量聚合消息再发kafka堆积消息，再由下游服务写入HBase

为了节约传输成本，kafka消息只有uid和avid，具体数据需要从redis读取
## 存储设计
## 可用性设计
## reference