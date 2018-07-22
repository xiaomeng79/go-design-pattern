//责任链模式
package chain

import "fmt"

//意图：分离不同职责，并且动态组合相关职责
//解决:职责链上的处理者负责处理请求，客户只需要将请求发送到职责链上即可，无须关心请求的处理细节和请求的传递，所以职责链将请求的发送者和请求的处理者解耦了

//角色：
//抽象处理者角色(Handler):定义出一个处理请求的接口
//具体处理者角色（ConcreteHandler）：具体处理者接受到请求后，可以选择将该请求处理掉，或者将请求传给下一个处理者。因此，每个具体处理者需要保存下一个处理者的引用，以便把请求传递下去。

//实例:采购货物，如果小于等于1000，部门经理审批，如果小于等于5000，副总经理审批，如果大于5000总经理审批

//采购请求
type PurchaseRequest struct {
	Amount float64 //采购金额
}

//处理请求接口
type IHandler interface {
	ProcessRequest(request *PurchaseRequest) //处理请求
}

//部门清理处理
type ManagerHandler struct {
	Successor IHandler //继承者
}

func (m *ManagerHandler) ProcessRequest(request *PurchaseRequest) {
	if m == nil {
		return
	}
	if request.Amount <= 1000.00 {
		fmt.Println("我是部门经理，我可以处理!")
	} else {
		m.Successor.ProcessRequest(request)
	}
}

//副总经理
type ViceGeneralManagerHandler struct {
	Successor IHandler
}

func (m *ViceGeneralManagerHandler) ProcessRequest(request *PurchaseRequest) {
	if m == nil {
		return
	}
	if request.Amount <= 5000.00 {
		fmt.Println("我是副总经理，我可以处理!")
	} else {
		m.Successor.ProcessRequest(request)
	}
}

//总经理

type GeneralManagerHandler struct {
	Successor IHandler
}

func (m *GeneralManagerHandler) ProcessRequest(request *PurchaseRequest) {
	if m == nil {
		return
	}
	if request.Amount > 5000.00 {
		fmt.Println("我是总经理，我可以处理!")
	} else {
		m.Successor.ProcessRequest(request)
	}
}
