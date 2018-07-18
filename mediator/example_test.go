package mediator

func ExampleUser_SendMessage() {
	//先实例化一个聊天室
	c := &ChatRoom1{}
	//实例化2个用户
	u1 := &User{Id: 1}
	u2 := &User{Id: 2}
	//加入聊天室
	u1.Join(c)
	u2.Join(c)
	//发送消息
	u1.SendMessage("大家好")
	//OutPut:
	//发送给用户：1,消息为:大家好
	//发送给用户：2,消息为:大家好
}
