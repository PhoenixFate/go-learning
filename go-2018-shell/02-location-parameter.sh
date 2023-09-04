#! /bin/bash
# 通过$0 $1 $2 ... 来接受传入的参数
echo "第一个参数: $0"
echo "第一个参数: $1"
echo "第一个参数: $2"
echo "第一个参数: $3"
echo "传递的参数的个数：$#"
# shellcheck disable=SC2145
echo "传递的所有的参数 \$@: $@"
echo "传递的所有的参数 \$*: $*"
echo "当前脚本的进程id: $$"


for i in "$@"
do
   echo $i   #会经历$#次循环
done

# shellcheck disable=SC2066
for i in "$*"
do
   echo $i  #只会进行一次循环,如果$*没有加双引号则会进行$#次循环
done