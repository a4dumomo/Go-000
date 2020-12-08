#学习笔记
1.基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
答:
    接受到signal信号马上跑出err,致使所有的ctx接收到done信号,server拒绝新来的请求，处理已接收的请求，并且需要做超时控制，

     