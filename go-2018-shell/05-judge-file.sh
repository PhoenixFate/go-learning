#!/bin/sh
# -d pathname 当pathname存在并且是一个目录时返回真
echo "./test"
if [ -d "test" ];then
  echo "test is a directory"
# -s filename 当filename存在并且文件大小大于0时返回真
elif [ -s "test" ];then
  echo "test is a file"
else
  echo "test is not existed"
fi

if [ -s "test.txt" ];then
  echo "test.txt is a file"
fi


