学习笔记

httpandlisten:for循环accept，然后接收到一个请求，就开启一个groutine去处理

log.fatal 不会调用defer，因为内部是用os.EXIT，所以一般在main和init中用。

防止goroutine泄露的三个方法：
它的生命周期，超时控制，并发让调用者决定后台还是前台执行。

channel满的时候，或者丢掉，或者阻塞。

内存重排、原子复制

mysql 64k-sector size，

Week03 作业题目：
1.基于 errgroup 实现一个 http server 的启动和关闭 ，
以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。