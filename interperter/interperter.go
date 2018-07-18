//解释器模式
package interperter

import "strings"

//意图：给定一个语言，定义它的文法表示，并定义一个解释器，这个解释器使用该标识来解释语言中的句子。
//主要解决：对于一些固定文法构建一个解释句子的解释器。
//何时使用：如果一种特定类型的问题发生的频率足够高，那么可能就值得将该问题的各个实例表述为一个简单语言中的句子。这样就可以构建一个解释器，该解释器通过解释这些句子来解决该问题。
//如何解决：构件语法树，定义终结符与非终结符。
//关键代码：构件环境类，包含解释器之外的一些全局信息，一般是 HashMap。
//应用实例：编译器、运算表达式计算。SQL 解析、符号处理引擎等
//优点： 1、可扩展性比较好，灵活。 2、增加了新的解释表达式的方式。 3、易于实现简单文法。
//缺点： 1、可利用场景比较少。 2、对于复杂的文法比较难维护。 3、解释器模式会引起类膨胀。 4、解释器模式采用递归调用方法。
//使用场景： 1、可以将一个需要解释执行的语言中的句子表示为一个抽象语法树。 2、一些重复出现的问题可以用一种简单的语言来进行表达。 3、一个简单语法需要解释的场景。

//实例:判断特定的文本是否符合正确

type Expression interface {
	Interpret(context string) bool
}

type TerminalExpression struct {
	Word string
}

// 终结符
func (te *TerminalExpression) Interpret(context string) bool {
	if strings.Contains(context, te.Word) {
		return true
	}
	return false
}

// 或
type OrExpression struct {
	A Expression
	B Expression
}

func (oe *OrExpression) Interpret(context string) bool {
	return oe.A.Interpret(context) || oe.B.Interpret(context)
}

// 与
type AndExpression struct {
	A Expression
	B Expression
}

func (ae *AndExpression) Interpret(context string) bool {
	return ae.A.Interpret(context) && ae.B.Interpret(context)
}
