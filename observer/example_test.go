package observer

func ExampleTest() {
	ob1 := &eventObserver{id: 1}
	ob2 := &eventObserver{id: 2}
	//通知者
	o := eventNotifier{observers: make(map[Observer]struct{}, 2)}
	//注册
	o.Register(ob1)
	o.Register(ob2)
	//发送
	o.Notify(Event{Data: 222})
	//OutPut:
	//*** Observer 1 received:222
	//*** Observer 2 received:222
}
