//模板模式
package template

import "fmt"

//意图：在一个抽象类中定义一个操作中的算法骨架（对应于生活中的大家下载的模板），而将一些步骤延迟到子类中去实现（对应于我们根据自己的情况向模板填充内容）。
// 模板方法使得子类可以不改变一个算法的结构前提下，重新定义算法的某些特定步骤，模板方法模式把不变行为搬到超类中，从而去除了子类中的重复代码。

//解决:一些通用的方法，在每个子类都写一遍，减少代码量和修改难度

//优点:1.实现了代码复用 2.能够灵活应对子步骤的变化，符合开放-封闭原则

//实例:人类有吃饭，睡觉，化妆的方法，男女人吃饭睡觉都相同，但化妆不同，我们只需要在男女类中重写化妆方法即可

type IPerson interface {
	SetSex()
	Eat(string)
	Dress()
}

//人类
type Person struct {
	sex string
}

func (p *Person) SetSex() {
	p.sex = "unknow"
}

func (p *Person) Eat(s string) {
	fmt.Println("吃食物 ", s)
}

func (p *Person) Dress() {
	fmt.Println("unknow")
}

//男人

type Man struct {
	p *Person
}

func (m Man) SetSex() {
	m.p.sex = "男人"
}

//重写化妆
func (m Man) Dress() {
	fmt.Println(m.p.sex, "化妆, ", "刮胡子")
}

//女人
type Women struct {
	p *Person
}

func (w Women) SetSex() {
	w.p.sex = "女人"
}

//重写化妆
func (w Women) Dress() {
	fmt.Println(w.p.sex, "化妆, ", "涂口红")
}
