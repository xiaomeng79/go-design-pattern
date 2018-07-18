//外观模式
package facade

import "fmt"

//意图：为子系统中的一组接口提供一个一致的界面，外观模式定义了一个高层接口，这个接口使得这一子系统更加容易使用
//解决: 降低访问复杂系统的内部子系统时的复杂度，简化客户端与之的接口

//实例:经典的实例：一台电脑，包含cpu,内存,硬盘等部件,若想要启动电脑，则先后必须启动cpu,memory,Disk等

//心得: 外观模式就是接待人员，导游

//1.外观模式的优点：
//
//（1）由于外观模式对外屏蔽了子系统的细节，因此外观模式降低了客户端对子系统使用的复杂性。
//
//（2）外观模式松散了客户端与子系统的耦合关系，让子系统内部的模块更易维护和扩展。
//
//（3）通过合理的使用外观模式，可以帮我们更好的划分访问的层次。
//
//2.外观模式的缺点：
//
//过多的或者不合理的使用外观模式容易让人产生困惑，到底是调用Facade好呢还是直接调用模块好？因此一定要合理使用外观模式。
//
//3.外观模式的应用场景：
//
//（1）当系统需要进行分层设计时可以考虑使用Facade模式。
//
//（2）在开发阶段，子系统可能会因为重构变得越发复杂，此时可以通过外观模式对外提供一个简单的接口，减少系统之间的依赖。
//
//（3）在维护一个遗留的大型系统时，可能这个系统已经变得非常难以维护和扩展，此时可以考虑为新系统开发一个Facade类，来提供遗留系统的比较清晰简单的接口，让新系统与Facade类交互。

//cpu
type cpu struct{}

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

type Computer struct {
	cpu    cpu
	memory memory
}

//实例化computer
func NewComputer() *Computer {
	return &Computer{cpu: cpu{}, memory: memory{}}
}
func (c *Computer) StartUp() {
	c.cpu.startup()
	c.memory.startup()
}

func (c *Computer) ShutDown() {
	c.cpu.shutdown()
	c.memory.shutdown()
}
