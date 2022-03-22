#!/bin/bash
set -e
set -x


CUR_DIR="$( cd "$( dirname "$0"  )" && pwd  )"

if [ $(uname -s) = 'Linux' ]; then
	mono CodeGenerator.exe lua $CUR_DIR/rpc $CUR_DIR/proto $CUR_DIR/lua netWorkDataRules.lua
else
	./CodeGenerator lua $CUR_DIR/rpc $CUR_DIR/proto $CUR_DIR/lua netWorkDataRules.lua
fi


echo ok

#read -p "生成成功 按任意键继续" -n 1 -r