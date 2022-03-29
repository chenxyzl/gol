#!/bin/sh
git status -uno| grep -e modified -e "new file"  > check.txt
cmd="du -sk"
# cat check.txt|while read line;
while read line
do
line=${line#*:}
# ls -lht $line
cmd="$cmd $line"
done<check.txt
cmd="$cmd | sort -hr"
eval $cmd > check.txt
while read line
do
size=`echo $line | awk -F" " '{print $1}'`
name=`echo $line | awk -F" " '{print $2}'`
if [ $size -gt 10240 ]; then
echo "file:$name too big"
exit 1
fi
# line=${line%%"\t"*}
# line=eavl "echo $line | cut -d$'\t' -f2"
done<check.txt
