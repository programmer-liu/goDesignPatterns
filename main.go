package main

import (
	"design/OOP"
	"fmt"
)

func main(){
	calculator()//计算器（面向对象三大特性：封装/继承/多态）
}

func calculator(){

	add :=&OOP.Add{NumberA: 10,NumberB: 5}
	fmt.Println(add.NumberA,"+",add.NumberB,"=",add.GetResult())

	sub :=&OOP.Sub{NumberA: 10,NumberB: 5}
	fmt.Println(sub.NumberA,"-",sub.NumberB,"=",sub.GetResult())

	multi :=&OOP.Multi{NumberA: 10,NumberB: 5}
	fmt.Println(multi.NumberA,"*",multi.NumberB,"=",multi.GetResult())

	divis :=&OOP.Divis{NumberA: 10,NumberB: 5}
	fmt.Println(divis.NumberA,"/",divis.NumberB,"=",divis.GetResult())
}
