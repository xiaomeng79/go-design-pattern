//适配器模式
package adapter

import "github.com/xiaomeng79/go-design-pattern/adapter/plug"

//意图：使得原本由于接口不兼容而不能一起工作的那些类可以一起工作
//解决:主要解决在软件系统中，常常要将一些"现存的对象"放到新的环境中，而新环境要求的接口是现对象不能满足的

//实例:电源有2孔插头，3孔插头，现在只有2孔的插排，需要一个适配器将3孔的插头转化成2孔的插头

//实现实例
//c := plug.TwoPlug{}//使用2孔插头充电
//c := plug.ThreePlugAdapter{} //使用3孔插头充电
//虽然3孔插头没实现充电接口，但是3孔插头的适配器实现了充电接口，所有都可以使用ToCharge方法

//使用插头充电
func ToCharge(c plug.IPlug) {
	c.Charge()
}
