package main

import (
    "fmt"
    "reflect"
)

func main() {
	t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
	v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值

	tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
	name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	/*
	error  反射的字段必须是可读写的意思是，如果下面这样写，那么会发生错误
	*/
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)

	/*
	如果要修改相应的值，必须这样写
	*/
	var x float64 = 3.4
	p := reflect.ValueOf(&x)
	v := p.Elem()
	v.SetFloat(7.1)

}