package prototype

func ExamplePeople_Clone() {
	//先实例化一个对象
	p := new(People)
	p.SetName("小明")
	p.Eat()
	//克隆一个对象小红
	p1 := p.Clone()
	p1.SetName("小红")
	p1.Eat()
	p1.Sleep()
	//OutPut:
	//小明 在吃饭
	//小红 在吃饭
	//小红 在睡觉

}
