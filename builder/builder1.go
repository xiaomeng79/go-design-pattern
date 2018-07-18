//建造者模式

package builder

//意图：将一个复杂的构建与其表示相分离，使得同样的构建过程可以创建不同的表示
//解决:复杂对象的各个部分经常面临着剧烈的变化，但是将它们组合在一起的算法却相对稳定

//实例:建造一个房子,都需要，不同材料搭建，涂颜色，建设面积不同，但都需要这3个属性

import ()

//建造一个房子

//房子颜色
//type Color uint32
//var (
//	ErrorNoColor = errors.New("no this color")
//)
//const (
//	Blue Color = iota
//	Red
//	Green
//)
//func (color Color) String() string {
//	switch color {
//	case Blue:
//		return "blue"
//	case Red:
//		return "red"
//	case Green:
//		return "green"
//	}
//	return "unknown"
//}
//
//func ParseColor(color string) (Color,error) {
//	switch strings.ToLower(color) {
//	case "blue":
//		return Blue,nil
//	case "red":
//		return Red,nil
//	case "green":
//		return Green,nil
//	}
//	var c Color
//	return c,ErrorNoColor
//}
//可以随便输入参数的链式调用
type IHouse interface {
	Color(string)    //房子颜色
	Material(string) //房子材料
	Area(float32)    //房子面积
}

type House struct {
	color    string
	material string
	area     float32
}

func (h *House) Color(c string) {
	h.color = c
}
func (h *House) Material(m string) {
	h.material = m
}
func (h *House) Area(a float32) {
	h.area = a
}

type IBuilder interface {
	Color(string) IBuilder
	Material(string) IBuilder
	Area(float32) IBuilder
	Build() IHouse
}
type Builder struct {
	house IHouse
}

func (b Builder) Color(c string) IBuilder {
	b.house.Color(c)
	return b
}

func (b Builder) Material(m string) IBuilder {
	b.house.Material(m)
	return b
}

func (b Builder) Area(a float32) IBuilder {
	b.house.Area(a)
	return b
}
func (b Builder) Build() IHouse {
	return b.house
}
