## 游戏actor拉起流程设计

### 使用一致性哈希算法做定位以及负载均衡

    1.所有请求都根据目标的id和type,采用一致性哈希算法，落到某个环上某个node上  
    2.node创建actor(此时不load数据),再使用mongodb的update{Upsert:true}来获取获取所有权(此时不需要nodeId)  
    3.每获取到读取一次此key是否属于自生，属于则获取到，不属于则间隔100ms循环3次获取(如果不做rebalance可以直接获取actor位置信息返回重定下命令)  
    4.失败则返回错误(通常次错误只会在宕机的30s出现,增减node也会有约一次db操作时长的时间出现此情况，一般会被3解决)  
    5.成功则拉起actor,actor间隔15s往db对应的key续约30s(也就是宕机后最多有30s此actor不可用)  

### actor保活

    1.利用mongo的ttl机制
    2.修改mongo ttl检查间隔为1s (db.adminCommand({setParameter: 1, ttlMonitorSleepSecs: 1});)
    3.每15s去续约一次ttl,mongodb的update({id=self,node=selfNode}{Upsert:true})。(注意:需要当前node id。updateCount=0则需要销毁自己)
    4.actor销毁时候需要立即delete key。({id=self,node=selfNode})。(注意:需要是自己的进程id)

### 注意

    1.获取所有权只需actor的id和Upsert:true  
    2.续约不仅需要actor的id和Upsert:true 还需要nodeId也一致(避免给别的node的actor续约了)

### 增加节点,减少节点,宕机等情况处理

    1.减少节点或者节点宕机,后续请求根据一致性哈希算法，此节点的actor触发时候会去环上临近的下一个node拉起,减少有0-2s不可用，宕机因为ttl延迟的原因0-30s不可用  
    2.增加节点，禁锢一致性哈希算法，此节点的前1个节点的部分数据会触发重新平衡，先下线后重新被消息触发再这个新节点拉起，中间有0-2s左右不可用。  

### todo