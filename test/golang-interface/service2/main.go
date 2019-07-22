package main

import (
	"fmt"
)

//声明一个USB的接口
type USB interface {
	Name() string
	Connecter //嵌入Connecter，从而USB就拥有Connecter的方法Connect()
}

type Connecter interface {
	Connect()
}

//声明一个PhoneConnect的结构去实现USB接口
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println(pc.Name() + "已经成功连接")
}

//TVConnecter只实现Connecter，没有实现USB,所以只能调用Connect()方法，因此转换成USB时会报错
type TVConnecter struct {
	name string
}

func (tv TVConnecter) Connect() {
	fmt.Println(tv.name + "已经成功连接")
}

func main() {
	tv := TVConnecter{"森森的电视"}

	var c Connecter = tv
	c.Connect()

	var a USB = USB(tv) //报错：cannot convert tv (type TVConnecter) to type USB:
	a.Connect()
}
