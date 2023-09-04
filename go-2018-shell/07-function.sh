#!/bin/sh
# 判断传递进来的文件名是不是一个目录，如果不是，创建目录

isDirectory(){
  echo "$0;$1"
  name=$1
  if [ -d "$name" ];then
    echo "$name 是一个目录"
  else
    # 创建目录
    mkdir "$name"
    result=$?
    if [ 0 -ne $result ];then
      echo "创建目录失败"
      exit
    else
      echo "创建目录成功"
    fi
  fi
}

isDirectory test2
