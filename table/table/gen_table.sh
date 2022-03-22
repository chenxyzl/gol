#!/bin/bash

CURDIR=$(cd $(dirname $0); pwd)
TEMPLATE=$CURDIR/table_template.txt
TABLE_LIST=$CURDIR/table_list.txt
OUT=$CURDIR/../../src/table

#go run $CURDIR/gen_table.go --help
go run $CURDIR/gen_table.go -template=$TEMPLATE -tablelist=$TABLE_LIST -out=$OUT
read -p "按任意键继续" -n 1 -r
#get_char()  
#{  
#  SAVEDSTTY=`stty -g`  
#  stty -echo  
#  stty raw  
#  dd if=/dev/tty bs=1 count=1 2> /dev/null  
#  stty -raw  
#  stty echo  
#  stty $SAVEDSTTY  
#} 
#
#get_char
