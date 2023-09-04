#!/bin/bash
num=10
if [ $num == "10" ];then
  echo "if num==10"
fi

if [ $num == "1" ];then
  echo "elif num==1"
elif [ $num == "10" ]; then
  echo "elif num==10"
else
  echo "else num"
fi
      
  