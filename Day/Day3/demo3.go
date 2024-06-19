package Day3

import (
	"fmt"
	"reflect"
)

// 这个代码有问题  []int{1, 2, 3, 4} 与  []interface{} 不能使用
//
//	func ReverseData(data []interface{}) []interface{} {
//		len1 := len(data)
//		r := make([]interface{}, len1)
//		for i, e := range data {
//			r[len1-1-i] = e
//		}
//		return data
//	}
func Test8() {
	r := []int{1, 2, 3, 4}
	ReverseData(r)
	strings := []string{"1", "2", "3", "4", "5"} // 字符串切片
	ReverseData(strings)

}

// 用反射实现 反转数组
func ReverseData(data interface{}) {
	// Typeof 这个是获取type的值   Kind 对应reflect
	if reflect.TypeOf(data).Kind() == reflect.Slice {
		len1 := reflect.ValueOf(data).Len()
		r := make([]interface{}, len1)
		for i := 0; i < len1; i++ {
			r[i] = reflect.ValueOf(data).Index(i).Interface() // 放回
		}
		//value, ok := data[0].(int)
		////if !ok {
		////	fmt.Printf("is not int %v\n", ok)
		////	return
		////}
		////fmt.Println(value)
		for i := 0; i < len1; i++ {
			fmt.Printf("Type is %v,  vaule is %v\n", reflect.TypeOf(r[i]), reflect.ValueOf(r[i]))
		}
	}
}

// 用泛型实现 反转数组
func reverseWithGenerics[T any](s []T) []T {
	l := len(s)
	r := make([]T, l)
	for i, e := range s {
		r[l-i-1] = e
	}
	return r
}
