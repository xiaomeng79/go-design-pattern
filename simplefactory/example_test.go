package simplefactory

func ExampleMouseFactory_CreateMouse() {
	var mf MouseFactory
	m := mf.CreateMouse(0)//戴尔鼠标
	m.SayMouseBrand()
	m = mf.CreateMouse(1)//惠普鼠标
	m.SayMouseBrand()
	//OutPut:
	//Dell Mouse
	//Hp Mouse

}
