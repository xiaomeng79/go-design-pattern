//工厂方法模式
package factorymethod

import "fmt"

//意图：定义一个创建对象的接口，让其子类自己决定实例化哪一个工厂类，工厂模式使其创建过程延迟到子类进行
//解决:主要解决接口选择的问题

//不足:选择判断问题还是存在的。也就是说，工厂方法把简单的工厂内部逻辑判断转移到了客户端来运行。

//实例:实现一个鼠标工厂的实现的接口，不同品牌的鼠标工厂实现它

//鼠标接口
type IMouse interface {
	SayMouseBrand()
}

//戴尔鼠标
type DellMouse struct{}

func (d DellMouse) SayMouseBrand() {
	fmt.Println("Dell Mouse")
}

//惠普鼠标
type HpMouse struct{}

func (h HpMouse) SayMouseBrand() {
	fmt.Println("Hp Mouse")
}

//鼠标工厂接口
type IMouseFactory interface {
	Create() IMouse
}

//戴尔鼠标工厂实现
type DellMouseFactory struct{}

func (d DellMouseFactory) Create() IMouse {
	return DellMouse{}
}

//惠普鼠标工厂
type HpMouseFactory struct{}

func (h HpMouseFactory) Create() IMouse {
	return HpMouse{}
}
