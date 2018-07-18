//中介者模式
package mediator

import "fmt"

//意图：用一个中介对象来封装一系列的对象交互，中介者使各对象不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互
//解决:对象与对象之间存在大量的关联关系，这样势必会导致系统的结构变得很复杂，同时若一个对象发生改变，我们也需要跟踪与之相关联的对象，同时做出相应的处理,由网状结构变为星状结构
//使用场景： 1、系统中对象之间存在比较复杂的引用关系，导致它们之间的依赖关系结构混乱而且难以复用该对象。 2、想通过一个中间类来封装多个类中的行为，而又不想生成太多的子类

//实例:一群用户之间信息交流，可以通过建立一个聊天室来进行，而不是每个用户和一群用户之间通讯，其中的聊天室就是中介者，如果一个用户向聊天室(中介者)发消息，聊天室就广播所有人消息

//聊天室
type IChatRoom interface {
	SendAllUserMessage(string) //向所有用户发送消息
	JoinUsers(...int)
}

//用户
type IUser interface {
	Join(IChatRoom)     //加入聊天室
	SendMessage(string) //发送消息
}

//聊天室
type ChatRoom1 struct {
	userIds []int
}

func (c *ChatRoom1) SendAllUserMessage(msg string) {
	for _, v := range c.userIds {
		fmt.Printf("发送给用户：%d,消息为:%s\n", v, msg)
	}
}
func (c *ChatRoom1) JoinUsers(uid ...int) {
	c.userIds = append(c.userIds, uid...)
}

//用户
type User struct {
	Id int
	c  IChatRoom
}

func (u *User) Join(c IChatRoom) {
	c.JoinUsers(u.Id)
	u.c = c
}
func (u *User) SendMessage(msg string) {
	u.c.SendAllUserMessage(msg)
}
