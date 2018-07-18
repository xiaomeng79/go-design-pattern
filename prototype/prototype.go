//原型模式
package prototype

import "fmt"

//意图：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象
//解决:利用已有的一个原型对象，快速地生成和原型对象一样的实例

//实例:人类有吃饭，睡觉方法，克隆一个人，使其拥有同样的方法

type IPerson interface {
	Eat()
	Sleep()
}

type People struct {
	name string
}

func (p *People) SetName(name string) {
	p.name = name
}

func (p *People) Eat() {
	fmt.Println(p.name, "在吃饭")
}

func (p *People) Sleep() {
	fmt.Println(p.name, "在睡觉")
}

func (p *People) Clone() *People {
	if p == nil {
		return nil
	}
	new_obj := (*p)
	return &new_obj
}
