#!/bin/sh

# get files
git status -uno| grep -e modified -e "new file"  > check.txt
# must not use ｜
# cat check.txt|while read line;

#get add/modify files sorted by disk use size
cmd="du -sk"
changed=0
while read line
do
changed=1
line=${line#*:}
cmd="$cmd $line"
done<check.txt

#no file change exit im
if [ $changed -eq 0 ]; then
exit 0
fi

# check size
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
done<check.txt
