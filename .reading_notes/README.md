# client-go 阅读笔记

## 初步阅读步骤
1. 查看 corev1 发起请求是如何进行的。


## Informer
informer 也被称为 shared informer ，他是可以共享使用的，如果每一个 informer 使用一个 reflector 那么会运行相当多的 listandwatch 会增加 api 的复杂。shared informer 可以使同一类资源 informer 共享一个 reflector 可以节约资源。


## 例子
- indexer/index_test.go: index 的使用方法


## 参考资料
[client-go的使用及源码分析](https://www.bookstack.cn/read/huweihuang-kubernetes-notes/develop-client-go.md)
[client-go 之 DeltaFIFO 实现原理](https://mp.weixin.qq.com/s?__biz=MzU4MjQ0MTU4Ng==&mid=2247485864&idx=1&sn=2011dfed276fe75a767d1e55f7d979ce&chksm=fdb906b5cace8fa3a01b911ac1004f6d6b57d8e0ae0d9a0e0746cdfa988947cc0e37ad2a980c&scene=21#wechat_redirect)