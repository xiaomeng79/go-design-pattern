package command

func ExampleInvoker_Do() {
	//电视实例
	tv := &TV{Name:"熊猫"}
	//实例化打开命令
	opencmd := OpenCommand{Tv:tv}
	//执行打开命令
	invoker := Invoker{cmd:opencmd}
	invoker.Do()
	//实例化关闭命令
	closecmd := CloseCommand{Tv:tv}
	//执行关闭命令
	invoker.SetCommand(closecmd)
	invoker.Do()

	//更换电视实例
	tv = &TV{Name:"黑熊"}
	opencmd = OpenCommand{Tv:tv}
	invoker.SetCommand(opencmd)
	invoker.Do()


	//OutPut:
	//打开熊猫电视
	//关闭熊猫电视
	//打开黑熊电视
}
