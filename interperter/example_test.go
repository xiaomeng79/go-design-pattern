package interperter

import "fmt"

func ExampleTerminalExpression_Interpret() {
	isMale := &OrExpression{&TerminalExpression{"Robert"}, &TerminalExpression{"John"}}
	fmt.Println("John is male?", isMale.Interpret("John"))
	isMarriedWoman := &AndExpression{&TerminalExpression{"Julie"}, &TerminalExpression{"Married"}}
	fmt.Println("Julie is a married women?", isMarriedWoman.Interpret("Married Julie"))
	//OutPut:
	//John is male? true
	//Julie is a married women? true
}
