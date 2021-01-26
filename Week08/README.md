第八周: 分布式缓存和分布式事务
---


## 缓存选型

### memcache
使用多线程模型 

提供简单的kv cache存储，value不超过1mb

通常使用大文本或者简单的kv结构使用

使用slab方式做内存管理

### redis
有丰富的数据类型

没有使用内存池，所以存在一定内存碎片，一般使用jemalloc来优化内存分配

部分系统使用memcache+redis双缓存，hgetall的结果缓存到memcache

### Proxy

早期使用twemproxy, 二次开发使用master-workers结合reuse_port

中期自己写overload(go版本github已开源， rust版本没有开源)，作为sidercar

如今使用redis cluster

## 缓存模式

###  数据一致性
- 同步操作DB
- 同步操作Redis
- 订阅binlog发消息到消息队列，下游Job重新补偿一次缓存操作

注：为了避免并发读写时读操作MISS写回时写回了旧数据，使用SETNX写回

Job可使用SET，可能出现ABA问题，使用DELETE更安全，失败时重试

###  多级缓存
先清理下游，下游的缓存时间一般是上游的2倍

###  热点缓存

- 小表广播：对于小数据，直接缓存到本地内存，定时reload
- 多Cluster支持

### 穿透缓存
- singlefly: 使用互斥锁保证归并回源，但不用于批量查询
- 分布式锁： 设置一个lock key，候选人轮询这个lock key，如果不存在则去读缓存，,缓存不命中就继续抢锁
- 消息队列
- lease租约

## 缓存技巧

### 通用小技巧
- 易读性的前提下，key尽量设置可能小
- 拆分key
- 空缓存设置
- 读缓存timeout失败后转读数据库时不触发回写缓存
- 序列化value使用protobuf，尽可能减少size
- generate工具化生成缓存处理通用代码

### redis小技巧
- BITSET过大时使用sharding,如div, mod = divmod(id, 10000),使用div做key,mod做offset
- List可作为抽奖的奖池，提前设置好
- SortedList:用于翻页、排序，杜绝使用zrange返回过大的集合
- 避免超大Value

## 分布式事务

### 事务消息
场景：支付宝转账到余额宝，两者使用不同的数据库

#### Transactional outbox本地消息表

支付宝在完成扣款的同事，记录消息数据到同一个数据库的msg表
```sql
BEGIN;
UPDATE A SET amount = amount - 1000 WHERE user_id = 1;
INSERT INTO msg(user_id, amount, status) Values  (1, 10000, 1);
COMMIT;
```
上述事务成功后，通知余额宝，余额宝处理成功后发送成功消息给支付宝，支付宝收到回复后删除该条消息

通知余额宝方式：
- polling publisher: 定时轮询msg表，把status=1的消息拿出来消息，按id排序保证顺序消费，使用消息队列或者rpc通知余额宝服务

#### transaction log tailing 

使用canal定义msg表(又交易流水表的话可以直接替代msg表)的bin log，解决轮询资源消耗大且延迟高的问题

#### 2PC二阶段提交
适用于同步场景

#### TCC(TRY, CONFIRM, CANCEL)
适用于同步场景

#### 幂等：消费端需要支持多次消费同一条消息
- 全局唯一ID + 去重
```sql
BEGIN;
SELECT COUNT(*) as cnt from message_apply where msg_id = 1000;
IF cnt == 0 THEN
    UPDATE B  SET amount = amount  + 10000 where user_id = 1;    
    INSERT INTO message_apply(msg_id) values(1000);
COMMIT;
```

- 版本号



