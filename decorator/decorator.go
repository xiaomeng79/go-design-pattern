//装饰器模式
package decorator

//意图：动态地给一个对象添加一些额外的职责,不改变对象的主要功能
//解决:对现有的对象添加新的功能，同时又不改变其结构,现有类的包装

//实例:对蛋糕进行包装，包装：普通包装，塑料包装 玻璃包装

//包装
type IPack interface {
	Desc() string
	Price() float32
}

//普通包装
type Packing struct {
	name  string
	price float32
}

func (p Packing) Desc() string {
	return p.name
}

func (p Packing) Price() float32 {
	return p.price
}

//塑料盒包装
type PlasticPack struct {
	pack  IPack
	name  string
	price float32
}

func (p PlasticPack) Desc() string {
	return p.pack.Desc() + p.name
}
func (p PlasticPack) Price() float32 {
	return p.pack.Price() + p.price
}

//玻璃盒包装
type GlassPack struct {
	pack  IPack
	name  string
	price float32
}

func (p GlassPack) Desc() string {
	return p.pack.Desc() + p.name
}
func (p GlassPack) Price() float32 {
	return p.pack.Price() + p.price
}

////测试
//func Test() (string,float32){
//	p := &Packing{name:"普通蛋糕",price:10.00} //普通蛋糕10元
//	gp := &GlassPack{pack:p,name:"外加玻璃套",price:2.00} //加一个玻璃盒子包装加2元
//	pgp := &PlasticPack{pack:gp,name:"外加塑料套",price:1.00} //加一个塑料盒子包装加1元
//	gpgp := &GlassPack{pack:pgp,name:"外加玻璃套",price:2.00} //加一个玻璃盒子包装加2元
//	return gpgp.Desc(),gpgp.Price()
//}
