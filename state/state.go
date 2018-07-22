//状态模式
package state

import "fmt"

//意图：当对象的状态改变时，同时改变其行为
//解决:对象的行为依赖于它的状态（属性），并且可以根据它的状态改变而改变它的相关行为
//使用场景： 1、行为随状态改变而改变的场景。 2、条件、分支语句的代替者

//实例:灯开关，当灯的状态为关着的时候，按开关，灯就会开，反之... ，开关就是行为，当状态不一样的时候，行为也不同

//灯
type Light struct {
	State ILightState
}

//开关
func (l *Light) PressSwitch() {
	if l.State != nil {
		l.State.PressSwitch(l)
	}
}

//灯状态接口
type ILightState interface {
	PressSwitch(*Light)
}

//开灯
type OpenLightState struct{}

func (o *OpenLightState) PressSwitch(l *Light) {
	fmt.Println("open light")
	//改变灯的下一个状态为关灯
	l.State = &CloseLightState{}
}

//关灯
type CloseLightState struct{}

func (c *CloseLightState) PressSwitch(l *Light) {
	fmt.Println("close light")
	//改变灯的下一个状态为开灯
	l.State = &OpenLightState{}
}
