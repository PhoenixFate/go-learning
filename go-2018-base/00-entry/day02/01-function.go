package main

import "fmt"

//默认函数名首字母小写：私有函数
//默认函数名首字母大写：公有函数

//不支持默认参数
//返回值可以只有类型没有变量名
//func FuncName(/*参数列表*/) (o1 type1, o2 type2/*返回类型*/) {
//	//函数体
//	return v1, v2 //返回多个值
//}

//无参数、无返回值
func MyFunction1() {
	a := 666
	fmt.Println("a=", a)
}

//有参无返回
func MyFunction2(a int, b int) {

	fmt.Println("function2: a=", a, "; b=", b)

}

//无参数，有返回
func MyFunction3() (int, int) {
	a := 666
	b := 777
	return a, b
}

//有参数，有返回
func MyFunction4(a int) int {
	return a
}

//go推荐，给返回值起一个变量名
func MyFunction5(a int) (result int) {
	return a
}

func MyFunction6(a, b int) (int, int) {
	return a, b
}

//go推荐到多个返回值的写法
func MyFunction9() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}

//!!!!!!!!!!!!!!!!!!!
//不定参数 ...type
func MyFunction7(args ...int) {

	fmt.Println(args)
}

//不定参数传参
func test1(args ...int) {
	fmt.Println("test1: ", args)
	//传递全部参数
	test2(args...)
	//传递部分参数
	test3(args[:2]...) //第0-2个参数（不包含第二个）
	test3(args[2:]...) //第2到最后到参数（包含第二个
}

func test2(args ...int) {
	fmt.Println("test2: ", args)
}

func test3(args ...int) {
	fmt.Println("test3: ", args)
}

//官方建议：最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差
func Test02() (value int) { //方式2, 给返回值命名
	value = 250
	return value
}

func Test03() (value int) { //方式3, 给返回值命名
	value = 250
	return
}

func main() {
	MyFunction1()
	MyFunction2(10, 20)
	a, b := MyFunction3()
	fmt.Println("function3: a=", a, "b=", b)

	e := MyFunction4(99)
	fmt.Println("function4: e=", e)

	f := MyFunction5(55555)
	fmt.Println("function5: f=", f)

	c, d := MyFunction6(88, 99)
	fmt.Println("function4: c=", c, "d=", d)

	k1, k2, k3 := MyFunction9()
	fmt.Printf("function9: k1=%d; k2=%d; k3=%d\n", k1, k2, k3)

	MyFunction7(1, 2, 3, 4, 5)
	test1(11, 12, 33, 44, 34)
	fmt.Println("01 function")
}
