#!/bin/bash
set -e
set -x


CUR_DIR="$( cd "$( dirname "$0"  )" && pwd  )"

CMD="CodeGenerator.exe"
if [ $(uname -s) = 'Linux' ]; then
	CMD="CodeGenerator"
elif [ $(uname -s) = 'Darwin' ]; then
    CMD="CodeGenerator_mac"
fi

$CUR_DIR/$CMD lua $CUR_DIR/rpc $CUR_DIR/proto $CUR_DIR/lua netWorkDataRules.lua


echo ok

#read -p "生成成功 按任意键继续" -n 1 -r