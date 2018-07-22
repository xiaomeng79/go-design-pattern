package visitor

func ExampleObjectStructure_Accept() {
	object := ObjectStructure{}
	object.Attach(&ConcreteElementA{"A"})
	object.Attach(&ConcreteElementB{"B"})

	cva := ConcreteVisitorA{"vA"}
	cvb := ConcreteVisitorB{"vB"}

	object.Accept(&cva)
	object.Accept(&cvb)

	//OutPut:
	//A vA
	//OperatorA
	//B vA
	//OperatorB
	//A vB
	//OperatorA
	//B vB
	//OperatorB

}
