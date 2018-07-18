//简单工厂模式,不是23种之一
package simplefactory

import "fmt"

//意图：将对象的创建过程进行了封装，用户不需要知道具体的创建过程，只需要调用工厂类获取对象即可
//解决:创建过程的复杂性

//不足:简单工厂的写法是通过switch-case来判断对象创建过程的。在实际使用过程中，违背了 开放-关闭原则，当然有些情况下可以通过反射调用来弥补这种不足。

//实例:实现一个生产鼠标工厂的类，通过输入参数不同，实例化不同品牌的鼠标工厂

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

//实例化鼠标工厂
type MouseFactory struct{}

func (m MouseFactory) CreateMouse(i int) IMouse {
	switch i {
	case 0: //戴尔
		return new(DellMouse)
	case 1: //惠普
		return new(HpMouse)
	}
	return nil
}
