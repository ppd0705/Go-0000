
## Question

1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

作为error，应该Wrap,但应该使用自定义的NotFoundError，解耦数据库强依赖。