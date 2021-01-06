# 学习笔记

## 1.滑动窗口统计
参考 Hystrix 实现一个滑动窗口计数器。

## 2.解决问题

+ 统计窗口的大小

    每秒独占一个bucket,保存到map,不存在则以当前时间戳来创建Key

+ 过期bucket处理

    a.每次调用incrment处理调用清理函数 removeOldBuckets
       优点:处理简单
       缺点:每次触发计数都需要遍历map，可能会存在效率问题
    b.另开goroutine定时处理,removeAsyOldBuckets
       优点:效率更高,异步处理
       
       
    
