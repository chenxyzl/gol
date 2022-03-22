## 协议文件格式

### CallDir ReturnType FuncName(Param in) = FuncId

* CallDir：可填CL或者LC，CL表示Client->Logic，LC表示LogicServer->Client
* ReturnType：Client->Logic发送带返回值，如果不带就填void.Logic->Client的必须是void
* FuncName：名字不能重复
* (Param in)：Param参数类型，in固定占位符
* FuncId：不能重复，按照功能分段，并且绝对不能侵占下面的id段  
  <font color=#FF000>**如果客户端请求的消息不能立即返回，返回消息的FuncId同请求消息FuncId相同**</font>

```c++
    ID_MSG_NONE          ID = 0
    ID_MSG_LOGIC_MIN     ID = 10000
    ID_MSG_LOGIC_MAX     ID = 30000
    ID_MSG_BEGIN         ID = 10001
    ID_MSG_END           ID = 10999
    ID_MSG_C2G_Login     ID = 10002
    ID_MSG_G2C_Login     ID = 10003
    ID_MSG_C2G_Create    ID = 10004
    ID_MSG_G2C_Create    ID = 10005
    ID_MSG_C2G_Offline   ID = 10006
    ID_MSG_C2G_KeepAlive ID = 10007
    ID_MSG_G2C_KeepAlive ID = 10008
    ID_MSG_C2G_SayHi     ID = 10009
    ID_MSG_G2C_SayHi     ID = 10010
    ID_MSG_G2C_Broadcast ID = 10011
    ID_MSG_G2C_Offline   ID = 10012
```

### 范例

```c++
//登录
CL cl.SC_Login RequestUserData(cl.CS_Login in) = 20000 RequestLogin

//同步数据(比如奖励之类)
LC void SyncSomething(cl.SC_CommonRet in) = 20001
```

## 工具说明

[../../tools/codegen/README.md](../../tools/codegen/README.md)

## 消息分段

### 每个大功能模块一个.h文件

* player 20000~20099