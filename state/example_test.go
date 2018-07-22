package state

func ExampleLight_PressSwitch() {
	//初始化灯的状态:开灯状态
	light := &Light{State: &OpenLightState{}}
	//按下开关,每次按下开关，改变状态
	light.PressSwitch()
	light.PressSwitch()
	light.PressSwitch()

	//OutPut:
	//open light
	//close light
	//open light
}
