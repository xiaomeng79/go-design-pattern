package threeplug

import "fmt"

//3孔插头，未实现2孔插头

type ThreePlug struct {
}

//充电方法，未实现充电接口
func (this ThreePlug) ThreePlugCharge() {
	fmt.Println("我是3孔插头，我使用方法：ThreePlugCharge充电")
}
