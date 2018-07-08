//模板模式
package template

import "fmt"

//意图：定义一个骨架，子类可以根据自身特性，重写方法
//解决:一些通用的方法，在每个子类都写一遍，减少代码量和修改难度

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

func (p *Person)SetSex() {
	p.sex = "unknow"
}

func (p *Person) Eat(s string) {
	fmt.Println("吃食物 ",s)
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
	fmt.Println(m.p.sex,"化妆, ","刮胡子")
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
	fmt.Println(w.p.sex,"化妆, ","涂口红")
}


