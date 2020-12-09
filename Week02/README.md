学习笔记

## Error vs Exception
error 只是一个interface,只有一个方法`Error() string`, New对象返回的一个pointer，方便等值比较

其他语言报错机制
- C语言只有单返回值，通常是int类型，通过判断返回值来确定是否成功
- C++引入了exception，但不知道具体是什么异常
- Java引入了checked exception

panic意味着fatal error，表示哪些不可恢复的错误

error的优点：
- 简单
- 也是一个返回值
- 没有隐藏的控制流

## Error Type
- Sentinel Error: 预定义错误，errors.New("xxx")
- Error types: 自定义错误，用switch type断言
- Opaque errors: 透明错误，提供公共方法断言，而不是直接暴露Error
```go
// Assert errors for behaviour, not type
type temporary interface {
    Temorary() bool
}

func IsTemporary(err error) bool {
    te, ok := err.(temporary)
    return ok && te.Temporay()
}
```

## Handling Error

- Indented flow is for errors， 无错误的代码将成成为一条直线
- eliminate-error-handling-by-eliminating-errors
- only handle errors once，使用pkg/errors, 初次报错时使用errors.Wrap保存堆栈信息，处理时使用errors.Cause拿到root error 

## Go1.13 errors
- 引入Unwrap拿到底层错误
- 引入Is和As方法检查错误
- fmt.Errorf引入`%w`可保留错误信息

## Go2 Error inspection

## reference
- [effective-go](https://golang.org/doc/effective_go.html)

- [eliminate-error-handling-by-eliminating-errors](https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors)

- [Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)
- [Proposal: Go 2 Error Inspection](https://go.googlesource.com/proposal/+/master/design/29934-error-values.md)