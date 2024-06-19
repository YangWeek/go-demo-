package day1

import (
	"fmt"
	"os"
)

func foo() (int, string) {
	return 1, "jasjdjad"

	// iota 常量计数器
}

// const (
// 	n1 = iota
// 	n2 = iota
// 	n3 = iota
// )

// go基本数据类型
func Test() {
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 在 Go 语言中，字符串是一种原生数据类型，而不是像在某些其他语言中那样是类似数组的对象。
	//在 Go 中，字符串是不可变的字节序列，用于存储文本数据

	s1 := `第一行
	第二行
	第三行
	`
	fmt.Println(s1)

	var str string = "yang"
	// 底层是byte数组
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v(%c) ", str[i], str[i])
	}
	fmt.Println()
	str1 := []byte(str)
	str1[1] = 'c'
	fmt.Println(string(str1))
}

func Test1() {
	var a int = 1
	if a == 1 {
		fmt.Println(1)
	} else if a == 2 {
		fmt.Println(2)
	}

	// Go 语言中的所有循环类型均可以使用for关键字
	for i := 0; i <= a; i++ {
		fmt.Println(i)
	}
}

func Test2() {

	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	name := "沙河小王子"

	// File结构体实现了io.Write 实现了多态
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	s2 := fmt.Sprintf("name:%s,age:%d", name, 20) // 返回来指针
	fmt.Println(s2)

}

func Test3() {
	p := &Person{
		name: "yang",
		sex:  "nan",
	}

	p.sayName()
	Test4(p)
}

// 以接口为类型的方法
// 多态实现
func Test4(n name) {
	n.sayName()
}

type name interface {
	sayName()
}

type Person struct {
	name string
	sex  string
}

func (p *Person) sayName() {
	fmt.Println(p.name)
	fmt.Println(p.sex)
}

func Test5() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}

	for _, v := range a {
		// 这个v 也是数组
		for _, v1 := range v {
			fmt.Printf("%s\t", v1)
		}
		fmt.Println()
	}
}

func Test6() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("t:%v len(t):%v cap(t):%v\n", a, len(a), cap(a))
	fmt.Printf("%p\n", &a)
	t := a[1:3:5]
	fmt.Printf("%p\n", &t)
	fmt.Println(t)

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
}
