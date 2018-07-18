//桥接模式
package bridge

import "fmt"

//意图：将抽象部分与实现部分分离，使它们都可以独立的变化

//解决:通过桥接模式，是抽象和实现之间建立一个关联关系

//实例:有抽象的消息和消息的实现

//区别：
//1.  Facade模式是为一个复杂系统提供一个简单的接口。
//比如你要去博物馆参观，很多东西，你一个个到处去问每个东西的管理员很麻烦，所以你找一个导游，让他给你一个个介绍，你只要找到导游就好了。导游就是门面。
//2. 适配器模式，引用一下GOF95中的话:
//适配器模式是把一个类的接口变换成客户端所期待的另一种接口，从而使原本因接口不匹配而无法工作的两个类能够工作到一起。
//举个例子，例如变压器
//3. Bridge模式:
//GOF95中的桥梁模式的描述：桥梁模式的用意是"将抽象化与实现化脱耦，使的二者可以独立变化。
//例如AWT的实现

//不同的出发点：
//桥接是先有桥，才有两端的东西,分离抽象化和实现，使两者的接口可以不同，目的是分离
//适配是先有两边的东西，才有适配器,改变已有的两个接口，让他们相容

//抽象的消息接口
type MessageAbstract interface {
	SendMessage(text, to string)
}

//实现的消息接口
type MessageImplementer interface {
	Send(text, to string)
}

//SMS消息
type MessageSMS struct{}

func (m *MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}
func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

//Email消息
type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

//抽象的实现
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
