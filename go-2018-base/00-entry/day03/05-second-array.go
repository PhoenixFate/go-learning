package main

//二维数组
import "fmt"

func main() {
	var a [3][4]int
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			a[i][j] = i + j
			fmt.Printf("a[%d][%d]=%d\t", i, j, a[i][j])
		}
		fmt.Println()
	}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))

	//定义的时候全部初始化
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			a[i][j] = i + j
			fmt.Printf("b[%d][%d]=%d\t", i, j, b[i][j])
		}
		fmt.Println()
	}

	//定义的时候部分初始化
	c := [3][4]int{{1, 2, 3}, {5, 8}, {9, 10, 11}}
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			a[i][j] = i + j
			fmt.Printf("c[%d][%d]=%d\t", i, j, c[i][j])
		}
		fmt.Println()
	}

	//两个数组之间的比较
	//相同类型的数组之间可以使用 == 或 != 进行比较，但不可以使用 < 或 >，也可以相互赋值：
	a1 := [5]int{1, 2, 3, 4, 5}
	a2 := [5]int{1, 2, 3, 4, 5}
	a3 := [5]int{1, 2, 3}
	fmt.Println("a1==a2: ", a1 == a2)
	fmt.Println("a2==a3: ", a3 == a2)
	a4 := [5]int{}
	//数组之间的相互赋值
	a4 = a1
	fmt.Println(a4)

	fmt.Println("05 二维数组")
}
