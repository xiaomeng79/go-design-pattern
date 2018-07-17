package abstractfactory

func ExampleDellFactory_CreateMouse() {
	df := new(DellFactory)
	df.CreateMouse()
	df.CreateKeybo()
	df.Mouse.SayMouseBrand()
	df.Keybo.SayKeyBoBrand()
	//OutPut:
	//Hp Mouse
	//Dell Keybo
}
