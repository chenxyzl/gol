#!/bin/bash

set -x
set -e

CURDIR=$(cd $(dirname $0); pwd)

RPCDIR=$CURDIR/../../protos/rpc
RPCOUTDIR=$CURDIR/../../src/logic/service/easyrpc
IMPORTS=protos/in/common,war/protos/out/cl
ROBOT_OUTDIR=$CURDIR/../../src/tools/robot
WEBCLIENT_OUTDIR=$CURDIR/../../src/tools/web_client


#生成rpc
# ? 分隔，分别对应rpc_interface.go rpc_register.go user_sender.go client_proxy.go
$CURDIR/CodeGenerator "rpc" $RPCDIR \
    logic/service/user,war/protos/out/cl?war/protos/out/cl?war/protos/out/cl,github.com/golang/protobuf/proto?war/protos/out/cl?logic/service/constant \
    $RPCOUTDIR \
    $RPCOUTDIR \
    $RPCOUTDIR/../user/user_sender.go \
    $RPCOUTDIR/../player/client_proxy.go

#机器人代码
$CURDIR/CodeGenerator "robot" $RPCDIR \
    war/protos/out/cl?war/protos/out/cl?war/protos/out/cl \
    $ROBOT_OUTDIR/robot/robot_rpcservice_interface.go \
    $ROBOT_OUTDIR/robot/robot_rpcservice_dispatch.go \
    $ROBOT_OUTDIR/robotsession/robot_session_sender.go \
    $ROBOT_OUTDIR/packet/rpc_def.go


#WebClient代码
$CURDIR/CodeGenerator "web_client" $RPCDIR \
    war/protos/out/cl?war/protos/out/cl?war/protos/out/cl \
    $WEBCLIENT_OUTDIR/clientsession/client_session_pack_gen.go

#read -p "按任意键继续" -n 1 -r
