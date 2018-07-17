package factorymethod


func ExampleDellMouseFactory_Create() {
	var dmf DellMouseFactory
	dm := dmf.Create()
	dm.SayMouseBrand()
	//OutPut:
	//Dell Mouse
}

func ExampleHpMouseFactory_Create() {
	var hmf HpMouseFactory
	hm := hmf.Create()
	hm.SayMouseBrand()
	//OutPut:
	//Hp Mouse
}
