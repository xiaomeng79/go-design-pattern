package plug

import "fmt"

//2孔插头
type TwoPlug struct {
}

//充电方法，实现充电接口
func (this TwoPlug) Charge() {
	fmt.Println("我是2孔插头，我使用方法：Charge充电")
}
