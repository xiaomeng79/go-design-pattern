//访问者模式
package visitor

import "fmt"

//意图:主要将数据结构与数据操作分离,不同的访问者，操作同一个元素，结果不同
//解决：稳定的数据结构和易变的操作耦合问题

//访问者模式和策略模式的区别:
//策略模式:上下文和算法是一对多的关系 如：算2个数的加减乘除，上下文就是者2个数，算法就是加减乘除
//访问者模式:上下和算法是多对多的关系

//实例:老板和会计查看这个月的总账，老板查看的是盈利多少钱，会计查看的是交了多少税

// 访问接口
type IVisitor interface {
	VisitConcreteElementA(ConcreteElementA)
	VisitConcreteElementB(ConcreteElementB)
}

// 具体访问者A
type ConcreteVisitorA struct {
	name string
}

func (c *ConcreteVisitorA) VisitConcreteElementA(ce ConcreteElementA) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorA) VisitConcreteElementB(ce ConcreteElementB) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}

// 具体访问者B
type ConcreteVisitorB struct {
	name string
}

func (c *ConcreteVisitorB) VisitConcreteElementA(ce ConcreteElementA) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorB) VisitConcreteElementB(ce ConcreteElementB) {
	if c == nil {
		return
	}
	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}

// 元素接口
type IElement interface {
	Accept(IVisitor)
}

// 具体元素A
type ConcreteElementA struct {
	name string
}

func (c *ConcreteElementA) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementA(*c)
}
func (c *ConcreteElementA) OperatorA() {
	if c == nil {
		return
	}
	fmt.Println("OperatorA")
}

// 具体元素B
type ConcreteElementB struct {
	name string
}

func (c *ConcreteElementB) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementB(*c)
}
func (c *ConcreteElementB) OperatorB() {
	if c == nil {
		return
	}
	fmt.Println("OperatorB")
}

// 维护元素集合
type ObjectStructure struct {
	list []IElement
}

func (o *ObjectStructure) Attach(e IElement) {
	if o == nil || e == nil {
		return
	}
	o.list = append(o.list, e)
}

func (o *ObjectStructure) Detach(e IElement) {
	if o == nil || e == nil {
		return
	}
	for i, val := range o.list {
		if val == e {
			o.list = append(o.list[:i], o.list[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructure) Accept(v IVisitor) {
	if o == nil {
		return
	}
	for _, val := range o.list {
		val.Accept(v)
	}
}
