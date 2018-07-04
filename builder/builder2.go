//建造者模式
package builder
//组件不变，组合经常变化

//组合一个大门，有锁，把手，门，不同的材料不同的价格

type Item interface {
	Name() string
	Price() float32
}
//锁
type Lock struct {
}

func (Lock) Name() string {
	return "锁"
}
func(Lock) Price() float32 {
	return 22.22
}

//把手
type Handle struct {
}

func (Handle) Name() string {
	return "把手"
}
func (Handle) Price() float32 {
	return 11.11
}

//门
type Door struct {
}

func (Door) Name() string {
	return "门"
}
func (Door) Price() float32 {
	return 33.33
}

type BigDoor []Item

func (this *BigDoor) AddItem(item ...Item) {
	*this = append(*this,item...)
}

func (this BigDoor) GetCost() (cost float32) {
	for _,v := range this {
		cost += v.Price()
	}
	return
}

func (this BigDoor) ShowItems() (msg string) {
	for _,v := range this {
		msg += "  组件:" + v.Name()
	}
	return
}

//建造
type DoorBuilder struct {
}

func (DoorBuilder) Build(item ...Item) (*BigDoor) {
	bigdoor := new(BigDoor)
	bigdoor.AddItem(item...)
	return bigdoor
}

