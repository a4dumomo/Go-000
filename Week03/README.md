#学习笔记
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
答:
server的停止需要接受signal的信号触发，server停止完成之后向其他goroutine发送完成信号,最后一并退出。
需要考虑的问题:
 1.http server什么时候开始退出？什么时候http server退出完成？
    接受到signal信号马上开始退出,拒绝新来的请求，处理完已经接受的请求？并且需要做超时控制，
 2.主程序什么时候退出?
    等待http server 处理完成发送 done信号给另外一个goroutine，最终主程序退出。
     