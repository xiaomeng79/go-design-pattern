package facade

func ExampleComputer_StartUp() {
	c := NewComputer()
	c.StartUp()
	//OutPut:
	//cpu startup
	//memory startup
}

func ExampleComputer_ShutDown() {
	c := NewComputer()
	c.ShutDown()
	//OutPut:
	//cpu shutdown
	//membory shutdown

}
