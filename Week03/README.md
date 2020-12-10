学习笔记

## Goroutine
- 启动goroutine时清楚其生命周期
- 建议提供方不主动启动goroutine,让调用者决定 
- 使用context做超时控制或者channel通信控制

## Memory model
- memory reordering内存重排
- memory barrie内存屏障
  - 读屏障
  - 写屏障
    - Write through
    - Write Back
- Memory model
  - Happens Before
    - if event e1 happens before event e2, then we say that e2 happens after e1
## Package sync
- 锁使用：最晚加锁，最早释放
- sync.atomic
- Mutex的实现
  - Barging
  - Handsoff
  - Spinning
- errgroup: 多任务并行执行并等待
- sync.Pool  

## chan
### 种类
- unbuffered channels
- buffered channels
### Design Philosophy

## Package context
- 作为函数的首个参数显式传递
- Value不建议更改

## references
- [The Go Memory Mode](https://golang.org/ref/mem)
- [内存屏障](https://zhuanlan.zhihu.com/p/43526907)