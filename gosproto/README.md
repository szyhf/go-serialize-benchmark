# 简述

原装的xjdrew/gosproto有bug，无法正确处理小于minInt32的负整数。

> 但我还是测了……

我自己有szyhf/go-sproto的分支，但只解决了bug并增加了值类型的支持。

> 讲道理，我的修改应该性能上序列化变慢（没空优化），反序列化可能会有点加快（相对序列化有一点点优化）。

ps: xjdrew接受了我的PR，现在已修复负数的bug。20170305.