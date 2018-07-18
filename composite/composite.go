//组合模式
package composite

import (
	"fmt"
	"strings"
)

//意图：统一对象和对象集，使得使用相同接口使用对象和对象集,使得接口具有一致性
//解决: 客户程序可以向处理简单元素一样来处理复杂元素，从而使得客户程序与复杂元素的内部结构解耦。
//优点:  1、高层模块调用简单。 2、节点自由增加
//缺点：在使用组合模式时，其叶子和树枝的声明都是实现类，而不是接口

//实例:公司/部门的实例，统一的接口公司管理，公司相当于树枝，部门相当于叶子
//公司可以增加/删除部门(对象),部门不可以

//组合接口(公司管理接口)

type ICompany interface {
	Add(ICompany)//增加
	Remove(ICompany)//移除
	Display(int)//展示下级
}

//公共的部分
type Company struct {
	Name string
}
//具体公司(树枝节点[根])
type RealCompany struct {
	Company
	list []ICompany
}
//新建一个公司
func NewRealCompany(name string) *RealCompany {
	return &RealCompany{Company{Name:name},[]ICompany{}}
}
//实现接口
func (r *RealCompany) Add(c ICompany) {
	if r == nil {
		return
	}
	r.list = append(r.list,c)
}
func (r *RealCompany) Remove(c ICompany) {
	if r == nil {
		return
	}
	for i,v := range r.list {
		if v == c {
			r.list = append(r.list[:i],r.list[i+1:]...)
		}
	}
	return
}
func (r *RealCompany) Display(depth int) {
	if r == nil {
		return
	}
	fmt.Println(strings.Repeat("-",depth)," ",r.Name)
	for _,val := range r.list {
		val.Display(depth + 2)
	}
}

//具体的部门(叶子节点)
//新建一个部门
type Department struct {
	Company
}
//新建一个公司
func NewDepartment(name string) *Department {
	return &Department{Company{Name:name}}
}
//实现接口
func (r *Department) Add(c ICompany) {}
func (r *Department) Remove(c ICompany) {}
func (r *Department) Display(depth int) {
	if r == nil {
		return
	}
	fmt.Println(strings.Repeat("-",depth)," ",r.Name)
}


