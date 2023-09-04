package main

import "fmt"

func notEmpty(data []string) []string {
	out := make([]string, 0, 10) //make([]string,0)
	for _, str := range data {
		if str != "" {
			out = append(out, str)
		}
	}
	return out
}

func notEmpty2(data []string) []string {
	i := 0
	for _, str := range data {
		if str != "" {
			data[i] = str
			i++
			fmt.Println(i)
		}
	}
	return data[:i]
}

func removeDuplicate(data []string) []string {
	out := make([]string, 0)
	for _, str := range data {
		i := 0
		for _, str2 := range out {
			if str == str2 {
				break
			}
			i++
		}
		if i == len(out) {
			out = append(out, str)
		}
	}
	return out
}

func main() {
	s1 := []int{1, 2, 4, 5, 6}
	//append(切片对象，待追加对象)
	s2 := append(s1, 66, 77)
	fmt.Println(s2)

	data := []string{"red", "", "black", "", "yellow"}
	fmt.Println("data: ", data)

	afterData := notEmpty(data)
	fmt.Println(afterData)
	fmt.Println("data: ", data)
	afterData2 := notEmpty2(data)
	fmt.Println(afterData2)

	data2 := []string{"red", "black", "red", "pink", "blue", "pink", "red"}
	afterData3 := removeDuplicate(data2)
	fmt.Println(afterData3)
}
