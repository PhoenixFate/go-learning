#!/bin/bash
# 对当前目录下的所有文件进行遍历
list=$(ls)
for name in $list;do
  echo  "文件名或者目录名：$name"
done