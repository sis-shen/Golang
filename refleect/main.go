package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func (p Person) String() string {
	return fmt.Sprintf("Name is %v , Age is %d", p.Name, p.Age)
}

func testReflect() {
	p := Person{Name: "Tom", Age: 10}
	pt := reflect.TypeOf(p)
	for i := 0; i < pt.NumField(); i++ {
		sf := pt.Field(i)
		fmt.Println(sf.Name, "  ", sf.Tag.Get("json"))
	}
}
func testModify() {
	p := Person{Name: "Tom", Age: 10}
	fmt.Println(p.Name)
	ppv := reflect.ValueOf(&p)
	ppv.Elem().FieldByName("Name").SetString("张三")
	fmt.Println(p.Name)

}

func testPassBy() {
	p := Person{Name: "Tom", Age: 10}

	pt := reflect.TypeOf(p)
	//遍历字段类型
	for i := 0; i < pt.NumField(); i++ {
		fmt.Println("字段", pt.Field(i).Name)
	}
	//遍历方法
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Println("字段", pt.Method(i).Name)
	}
}

func testJson() {
	fmt.Println("start test Json")
	p := Person{Name: "Tom", Age: 10}

	// struct to Json
	jsonBytes, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonBytes))
	} else {
		fmt.Println(err)
	}
	// json to struct
	respJson := "{\"name\": \"Peter\", \"age\": 30}"
	json.Unmarshal([]byte(respJson), &p)
	fmt.Println(p)
}

func (p Person) Print(prefix string) {
	fmt.Printf("%s %s \n", prefix, p.Name)
}

func testCall() {
	p := Person{Name: "Tom", Age: 10}
	pv := reflect.ValueOf(p)
	myPrint := pv.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("老登 ")}
	myPrint.Call(args)
}

func main() {
	//testModify()
	//testPassBy()

	//testJson()
	//testReflect()
	testCall()
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
