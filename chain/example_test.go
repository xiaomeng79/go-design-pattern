package chain

func ExampleManagerHandler_ProcessRequest() {
	//新建一个采购请求,4999在副总经理审批的范围
	req := &PurchaseRequest{Amount: 4999.00}
	//建立处理链
	h := &ManagerHandler{Successor: &ViceGeneralManagerHandler{Successor: &GeneralManagerHandler{}}}
	h.ProcessRequest(req)
	//OutPut:
	//我是副总经理，我可以处理!
}
