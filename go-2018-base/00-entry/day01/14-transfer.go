package main

//类型转换
import "fmt"

func main() {

	var flag bool
	flag = true
	fmt.Printf("flag: %T\n", flag)

	//bool 类型不能转int
	//Cannot convert expression of type bool to type int
	//int(flag)

	//0假  非0真
	//整型不能转bool

	var ch byte
	ch = 'a'
	var t int
	//类型转换，能够转换才能转换
	t = int(ch)
	fmt.Printf("ch=%c \n", ch)
	fmt.Printf("t=%d\n", t)

	fmt.Println("14")
}
