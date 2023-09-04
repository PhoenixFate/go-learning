package main

//. 操作：这个包导入之后在调用的时候，可以省略包名
//点操作
import . "fmt"

// 给包名起别名
import myTime "time"

//有时，用户可能需要导入一个包，但是不需要引用这个包的标识符。在这种情况，可以使用空白标识符_来重命名这个导入：
//_操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
import _ "log"

func main() {

	Println("01")
	myTime.Sleep(1000)
	Println("11 point")

}
