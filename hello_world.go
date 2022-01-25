package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	name string
	age int
	male bool
}

func (p Person) gao() {
	fmt.Println("gao")
}
func main(){
	fmt.Println("hello world")
	a := 0
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())

	b := &a
	fmt.Println(reflect.TypeOf(b), reflect.TypeOf(b).Kind(), reflect.TypeOf(b).Elem())

	p := Person{}
	fmt.Println(reflect.TypeOf(p), reflect.TypeOf(p).Kind())
	typeOfP := reflect.TypeOf(p)
	for i:=0; i< typeOfP.NumField(); i++ {
		fmt.Println('1', typeOfP.Field(i).Name, typeOfP.Field(i).Type, typeOfP.Field(i).Type.Kind())
	}
	for i:=0; i< typeOfP.NumMethod(); i++ {
		fmt.Println('2', typeOfP.Method(i).Name, typeOfP.Method(i).Type, typeOfP.Method(i).Type.Kind())
	}
}

