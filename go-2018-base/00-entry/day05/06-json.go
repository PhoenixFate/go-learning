package main

import "fmt"
import "encoding/json"

//要把结构体生成 json，成员首字母必须大写
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

//编码的时候换个名字
type IT2 struct {
	Company  string   `json:"MyCompany"`
	Subjects []string `json:"MySubjects"`
	IsOk     bool     `json:"-"`       //次字段转json的时候会被忽略
	Price    float64  `json:",string"` //转json的时候，数字转string
}

func main() {
	s := IT{"phoenix", []string{"go", "java", "python"}, true, 23.2}
	fmt.Printf("s.Subjects type: %T\n", s.Subjects)

	//根据内容返回json字符串
	//jsonText,err1:=json.Marshal(s)
	jsonText, err1 := json.MarshalIndent(s, "", " ") // 格式化编码

	if err1 == nil {
		fmt.Println(string(jsonText))
	}

	s2 := IT2{"phoenix", []string{"go", "java", "python"}, true, 23.2}
	json2, err2 := json.MarshalIndent(s2, "", " ")
	if err2 == nil {
		fmt.Println(string(json2))
	}

	//map转json
	m := make(map[string]interface{}, 4)
	m["company"] = "aaa"
	m["subjects"] = []string{"go", "java"}
	m["isOk"] = true
	m["price"] = 9.99
	json3, err3 := json.MarshalIndent(m, "", " ")
	if err3 == nil {
		fmt.Println(string(json3))
	}

	//json字符串转换为struct对象
	var temp IT2
	jsonBuffer := `{
	 "MyCompany": "phoenix",
	 "MySubjects": [
	  "go",
	  "java",
	  "python"
	 ],
	 "Price": "23.2"
	}`
	err4 := json.Unmarshal([]byte(jsonBuffer), &temp)
	if err4 == nil {
		fmt.Println("---- temp -----")
		fmt.Println(temp)
		fmt.Println("---- temp -----")
	} else {
		fmt.Println(err4)
	}

	//json转map对象
	mm := make(map[string]interface{}, 2)
	err5 := json.Unmarshal([]byte(jsonBuffer), &mm)
	if err5 == nil {
		fmt.Println(mm)
		fmt.Printf("mm MyCompany type: %T\n", mm["MyCompany"])
		str := mm["MyCompany"]
		fmt.Println(str)
	} else {
		fmt.Println(err5)
	}

	fmt.Println("06 json")
}
