#!/bin/bash
set -x
set -e

DIR="$( cd "$( dirname "$0"  )" && pwd  )"
echo pwd=$DIR
PojDir=$DIR/..

OS_NAME=$(uname -s)

if [ "$OS_NAME" = 'Linux' ]; then
  echo "未配置linux下的protoc和protoc-gen-go"
	exit 1
elif [ "$OS_NAME" = 'Darwin' ]; then
  PROTOC=$DIR/protoc/protoc
	GenGoPath=$DIR/protoc/protoc-gen-go
	export PATH=$DIR/protoc:$PATH
else
  echo "未配置windows下的protoc和protoc-gen-go"
  exit 1
fi

OutPath=$PojDir/foundation

$PROTOC -I=$PojDir/protos/in -I=$PojDir/protos/out/ --go_out=$OutPath $PojDir/protos/in/*.proto
$PROTOC -I=$PojDir/protos/in -I=$PojDir/protos/out/ --go_out=$OutPath $PojDir/protos/out/*.proto
#
#for i in ${in_protos[@]}
#do
#	if [ ! -d $path ]; then
#		mkdir -p $path
#	fi
#	$PROTOC -I=$PojDir/protos/in --proto_path=$PojDir/protos/out/ --go_out=$OutPath $PojDir/protos/in/$i.proto
#done
#
#
#for i in ${out_protos[@]}
#do
#	if [ ! -d $path ]; then
#		mkdir -p $path
#	fi
#	$PROTOC -I=$PojDir/protos/out -I=$PojDir/protos/in --go_out=$OutPath $PojDir/protos/out/$i.proto
#done