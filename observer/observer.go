//观察者模式
package observer

import (
	"fmt"
)

//意图：定义对象间的一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都得到通知并被自动更新
//解决:一个对象状态改变给其他对象通知的问题，而且要考虑到易用和低耦合，保证高度的协作

//实例:当程序接收到用户的信号的时候，通知其他监听的程序做响应处理

type (
	//事件
	Event struct {
		Data int64
	}
	//观察者
	Observer interface {
		OnNotify(Event)
	}
	//通知者
	Notifier interface {
		//注册
		Register(Observer)
		//取消注册
		Deregister(Observer)
		//通知
		Notify(Event)
	}
)

type (

	eventObserver struct {
		id int
	}
	eventNotifier struct {
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received:%d\n",o.id,e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers,l)
}

func (o *eventNotifier) Notify(e Event) {
	for l := range o.observers {
		l.OnNotify(e)
	}
}

//测试
func Test() {
	//观察者2个
	ob1 := &eventObserver{id:1}
	ob2 := &eventObserver{id:2}
	//通知者
	o := eventNotifier{observers:make(map[Observer]struct{},2)}
	//注册
	o.Register(ob1)
	o.Register(ob2)
	//发送
	o.Notify(Event{Data:222})
}

