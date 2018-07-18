//命令模式
package command

import "fmt"

//意图：将一个请求封装成一个对象，从而使您可以用不同的请求对客户进行参数化
//解决: 在软件系统中，行为请求者与行为实现者通常是一种紧耦合的关系

//实例:人使用遥控器按下开机键打开电视的过程，人就是客户端,电视遥控器就是调用者，按键就是具体的调用的命令，电视就是接收执行命令者

//心得: 此实例中，电视为行为实现者,可以随时更换,而不影响行为请求者

type ICommand interface {
	Press() //按键操作，执行命令
}

//按键
//开机命令
type OpenCommand struct {
	Tv *TV
}

func (o OpenCommand) Press() {
	o.Tv.Open()
}

//关机命令
type CloseCommand struct {
	Tv *TV
}

func (c CloseCommand) Press() {
	c.Tv.Close()
}

//电视
type ITV interface {
	Open()
	Close()
}
type TV struct {
	Name string //电视名称
}

func (t *TV) Open() {
	fmt.Println("打开" + t.Name + "电视")
}

func (t *TV) Close() {
	fmt.Println("关闭" + t.Name + "电视")
}

//遥控器
//按键按下就是设置命令，松开就是执行命令
type IInvoker interface {
	SetCommand(ICommand) //设置命令
	Do()                 //执行命令
}

type Invoker struct {
	cmd ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.cmd = command
}

func (i *Invoker) Do() {
	i.cmd.Press()
}
