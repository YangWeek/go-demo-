package Day3

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

// Go语言中的 map 在并发情况下，只读是线程安全的，同时读写是线程不安全的
func Test1() {
	//var myMap map[string]int  这个是声明
	//myMap = make(map[string]int) 这个是初始化

	map1 := make(map[int]string, 200) //
	map1[1] = "jasjdsja"
	map1[2] = "kaskdad"
	for key, v := range map1 {
		fmt.Printf("key : %v, vaule : %v\n", key, v)
	}
}

// Go语言封装
// 工厂模式是用来创建对象的一种最常用的设计模式，
// 我们不暴露创建对象的具体逻辑，而是提供一个用于创建对象的接口
type person struct {
	name string
	age  int
}

func NewPerson() *person {
	return &person{}
}

func (p *person) SetName(name1 string) {
	p.name = name1
}

func (p *person) SetAge(age1 int) {
	p.age = age1
}

type Dongwu struct {
	name string
}

func (p *Dongwu) paintDongwu() {
	fmt.Println(p.name)
}

func Test2() error {
	// interface 这是万能类型
	map1 := make(map[string]interface{})
	map1["name"] = "yang"
	map1["age"] = 18

	fmt.Println(map1)

	// interface 断言机制  这个明确传入数据的类型
	// 类型断言只能用于接口类型
	name, ok := map1["name"].(string)
	if !ok {
		return errors.New("类型断言失败")
	}
	fmt.Println(name)
	return nil
}

type Read interface {
	ReadBook()
}

type Write interface {
	WriteBook()
}

type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("读书")
}

func (b *Book) WriteBook() {
	fmt.Println("写书")
}

func Test3() error {
	// 内置Pair 断言
	// pair<type = book, value = &book{}>
	book := &Book{}

	var r Read
	// pair<type = book, value = &book{}>
	r = book
	r.ReadBook()

	var w Write
	// pair<type = book, value = &book{}>
	w = r.(Write)
	w.WriteBook()
	return nil
}

// 放射 标签
type Person struct {
	name string
	age  int
}

func (this person) Call() {
	fmt.Println(this.name, this.age)
}

//func Test5() error {
//	// 复杂的放射
//	p := &person{
//		name: "yang",
//        age:  18,
//	}
//
//	return nil
//}
//
//func RefsTest(m interface{}) {
//	reflect.Typeof(m)
//}

// 结构体标签
type reume struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func tagTest(r interface{}) {
	t := reflect.TypeOf(r).Elem() //  这个必须得指定类型
	// 只能用于结构体类型
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		fmt.Printf("json tag is %v\n", tag)
	}
}

func Test6() {
	r := &reume{
		Name: "yang",
		Sex:  "nan",
	}
	tagTest(r)
	// 编码 json格式 struct -> json
	// json 编解码
	json_str, err := json.Marshal(r)
	if err != nil {
		fmt.Printf("jsonMarshal error: %v\n", err)
		return
	}
	fmt.Printf("json : %s\n", json_str)
}

func newTesk() {
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
}

// goroutine 管道
func Test7() {
	//go newTesk()
	//time.Sleep(time.Second * 1)
	//fmt.Println("main")

	go func() {
		defer fmt.Printf("main goroutine ending\n")
		func() {
			defer fmt.Println("second goroutine ending\n")
			runtime.Goexit()
		}()
	}()
	time.Sleep(time.Second * 1)
}
