package plug

import "github.com/xiaomeng79/go-design-pattern/adapter/plug/threeplug"

//通过适配器，在不改变3孔充电方法的情况下，将3孔充电适配2孔充电方法的接口

type ThreePlugAdapter struct {
	threeplug.ThreePlug
}

//实现充电接口，不需要关心2孔插头内部的实现
func (this ThreePlugAdapter) Charge() {
	this.ThreePlug.ThreePlugCharge()
}

