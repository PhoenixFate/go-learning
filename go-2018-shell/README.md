# shell脚本的学习

## 1.1 shell常识
### shell首行
```
shell脚本的首行是固定写法 不写使用系统默认的shell解释器
指定shell脚本的命令解释器
#! /bin/bash
#! /bin/sh
```

### 注释
- \#开头 代表shell中的注释

### 查看系统默认shell
- echo $SHELL

### 执行脚本
- sh xxx.sh
- ./xxx.sh  (chmod u+x xxx.sh 给当前sh添加执行权限)

## 1.2 shell中的变量
- 变量的定义
    - 普通的变量（本地变量）
    ```
    # 定义变量 定义完成之后必须要赋值 =前后不能有空格
    # = 前后有空格 等价于==（比较）
    # 普通变量只能在当前进程中使用
    temp=666
    ```
    - 环境变量 （环境变量一般都是大写的）
    ```
    # 可以把环境变量理解为全局变量，在当前操作系统中都可以访问环境变量
    # 查看当前系统的环境变量 env
    # 分类
    - 系统自带都环境变量 
        - SHELL
        - USER
        - HOME
    - 用户自定义的环境变量
        - 将普通变量提升为系统级别的环境变量
        GOPATH=/Users/phoenix/go 普通环境变量
        set GOPATH=/Users/phoenix/go 系统环境变量 
        export GOPATH=/Users/phoenix/go 系统环境变量
        ~/.bashrc
    ```
  
- 位置变量
    > 执行脚本的时候，可以给脚本传递参数，在脚本内部接受这些参数，需要使用位置变量
    ```
    # ./xxx.sh aa bb cc
    - $0 执行的脚本的文件名字
    - $1 第一个参数
    - $2 第二个参数
    ```      
- 特殊变量
    - $# 获取传递参数的个数
    - $@ 获得传递的所有的参数
    - $? shell脚本执行之后的状态 （成功=0or失败>0）(只能在脚本调用结束之后调用)
    - $$ 脚本进程执行之后对应的进程id
    
- 普通变量取值
```
# 定义变量
    value=123 # 默认以字符串来处理
    value2="123"
    echo $value # 变量的使用
# 使用变量
    - $变量名
    - ${变量名}
# 取命令之后的结果值
    value=$(shell命令)  #推荐
    value=`shell命令`
# 引号的使用
    双引号 "name: $name" 会替换变量
    单引号 'name: $name' 会把里面所有内容原样输出
```
  
## 1.3条件判断和循环
- shell脚本中的if条件判断
```
    # 注意事项
        - if后面要加空格
        - [ 条件 ] 条件前后都要加空格
    if [ 条件判断 ];then
        逻辑处理->shell命令
    fi

    # =========
    if [ 条件判断 ]
    then
        逻辑处理->shell命令
    fi
    
    # =========
    # if ... elif ..fi
    if [ 条件判断 ];then
        shell命令
    elif [ 条件判断 ];then
        shell命令
    elif [ 条件判断 ];then
        shell命令
    else
        shell命令
    fi

```
## 1.4 shell条件测试参数
- -d pathname 当pathname存在并且是一个目录时返回真
- -s filename 当filename存在并且文件大小大于0时返回真

## 1.5 for
- shell中的循环 for/while
```
# 语法 for 变量 in 集合; do; done
```

## 1.6 shell中的函数
```
# 没有函数修饰符，没有参数，没有返回值
# 格式
funcName(){
    函数体
    # 获得传入函数的参数
    arg1=$0
    arg2=$1
    
}
# 没有参数列表，但是可以传参
# 函数调用 的时候传参数
funcName aa bb cc dd
# 函数调用之后的状态:
0 -> 调用成功
非0 -> 失败
```

    