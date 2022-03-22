## 游戏rpc设计

    1.rpc生成器 生成rpc接口和map
    2.实现rpc接口，并且注册到rpc组件内

## 网络连接

    1.对外kcp(可代理tcp) + websocket
    2.对内nats(100万qps的消息队列)

## gate转发流程

    1.gate维护链接，保存seesion和uid。
    2.对外消息包lengh|Request。
        a.转发时候设置Request.extend:uid(一致性hash的对象),并通过nats转发到home
        b.home push需要设置Reply.extend:session，gate通过session找到对应的gate