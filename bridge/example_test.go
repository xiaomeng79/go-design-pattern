package bridge

func ExampleCommonMessage_SendMessage() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	//OutPut:
	//send have a drink? to bob via SMS
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send [Urgency] have a drink? to bob via SMS
}
