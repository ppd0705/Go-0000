第十二周：Runtime
---

## Goroutine原理
Goroutine是一个与其他goroutine并行运行在同一地址空间的Go函数和方法

### 特性
- 内存占用：初始分配2KB，运行过程中可以自动扩容，而thread会默认分配一个1-8MB的栈内存，且不可扩容
- 创建/销毁：在用户态进行
- 调度切换：切换成本远比线程小
- 复杂性：通讯简单，创建和退出简单

### GMP
G：goroutine
M: thread
P: local queue 
### M:N模型
Go runtime会创建M个线程，之后的多个N个goroutine都会依附在这M个线程上执行

### goroutine的创建
- G0
- M0

### syscall
Go封装了了syscall

## 内存分配原理
## GC原理
## channel原理