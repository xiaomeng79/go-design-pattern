//外观模式
package facade

import "fmt"

//意图：为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用
//解决: 降低访问复杂系统的内部子系统时的复杂度，简化客户端与之的接口

//实例:经典的实例：一台电脑，包含cpu,内存,硬盘等部件,若想要启动电脑，则先后必须启动cpu,memory,Disk等

//心得: 外观模式就是接待人员，导游


//cpu
type cpu struct {}
func (c cpu) startup() {
	fmt.Println("cpu startup")
}
func (c cpu) shutdown() {
	fmt.Println("cpu shutdown")
}

//memory
type memory struct{}
func (m memory) startup() {
	fmt.Println("memory startup")
}
func (m memory) shutdown() {
	fmt.Println("membory shutdown")
}

//使用外观模式

type Computer struct{
	cpu cpu
	memory memory
}
//实例化computer
func NewComputer() *Computer {
	return &Computer{cpu:cpu{},memory:memory{}}
}
func (c *Computer)StartUp() {
	c.cpu.startup()
	c.memory.startup()
}

func (c *Computer) ShutDown() {
	c.cpu.shutdown()
	c.memory.shutdown()
}

