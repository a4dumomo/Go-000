# 学习笔记

## 1.滑动窗口统计
参考 Hystrix 实现一个滑动窗口计数器。

## 2.解决问题
1.统计窗口的大小:每秒独占一个bucket,保存到map,不存在则以当前时间戳来创建Key

2.过期bucket处理：
    a.每次调用incrment处理调用清理函数 removeOldBuckets
       优点:处理简单
       缺点:每次触发计数都需要遍历map，可能会存在效率问题
    b.另开goroutine定时处理
       优点:效率更高,异步处理
3.边界节点处理
    map中的key是一直浮动，高并发中，边界的秒key在统计sum中被计算到，但是计算average时，map中的key可能被清除，导致计算不准
   个数不准确处理：在计算sum时，同时返回计算了多少个bucket
   key边界：清除函数中removeOldBuckets 多保留1秒的bucket
    
