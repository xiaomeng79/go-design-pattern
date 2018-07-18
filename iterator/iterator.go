//迭代器模式
package iterator

import "fmt"

//意图：提供一种方法顺序访问一个聚合对象中各个元素, 而又无须暴露该对象的内部表示。
//主要解决：不同的方式来遍历整个整合对象。
//何时使用：遍历一个聚合对象。
//如何解决：把在元素之间游走的责任交给迭代器，而不是聚合对象。
//关键代码：定义接口：hasNext, next。
//优点： 1、它支持以不同的方式遍历一个聚合对象。 2、迭代器简化了聚合类。 3、在同一个聚合上可以有多个遍历。 4、在迭代器模式中，增加新的聚合类和迭代器类都很方便，无须修改原有代码。
//缺点：由于迭代器模式将存储数据和遍历数据的职责分离，增加新的聚合类需要对应增加新的迭代器类，类的个数成对增加，这在一定程度上增加了系统的复杂性。
//使用场景： 1、访问一个聚合对象的内容而无须暴露它的内部表示。 2、需要为聚合对象提供多种遍历方式。 3、为遍历不同的聚合结构提供一个统一的接口。
//注意事项：迭代器模式就是分离了集合对象的遍历行为，抽象出一个迭代器类来负责，这样既可以做到不暴露集合的内部结构，又可让外部代码透明地访问集合内部的数据。
//实例: 迭代一个聚合的人类的信息

type Person struct {
	Name   string
	Age    int
	Enjoys string
}

type peopleIterator interface {
	first() *Person
	next() *Person
}

type Peoples struct {
	peoples []*Person
}

func (p *Peoples) Add(people ...*Person) {
	if p == nil {
		return
	}
	p.peoples = append(p.peoples, people...)
}

//创建迭代器
func (p *Peoples) CreateIterator() *PeoplesIterator {
	if p == nil {
		return nil
	}
	return &PeoplesIterator{p, 0}
}

//人迭代器
type PeoplesIterator struct {
	ps    *Peoples
	index int
}

//first
func (pi *PeoplesIterator) first() *Person {
	if pi == nil {
		return nil
	}
	if len(pi.ps.peoples) > 0 {
		pi.index = 0
		return pi.ps.peoples[pi.index]
	}
	return nil
}

//next
func (pi *PeoplesIterator) next() *Person {
	if pi == nil {
		return nil
	}
	if len(pi.ps.peoples) > pi.index+1 {
		pi.index++
		return pi.ps.peoples[pi.index]
	}
	return nil
}

func (pi *PeoplesIterator) Iterator() {
	for b := pi.first(); b != nil; b = pi.next() {
		fmt.Printf("姓名:%s,年龄:%d,爱好:%s\n", b.Name, b.Age, b.Enjoys)
	}
}
