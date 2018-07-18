package composite

func ExampleNewRealCompany() {
	//新建一个总公司
	root := NewRealCompany("华夏集团")
	//新建1个部门
	rs := NewDepartment("人事部")
	//添加到总公司
	root.Add(rs)
	//再建2个分公司
	c1 := NewRealCompany("华夏集团北京分公司")
	c2 := NewRealCompany("华夏集团上海分公司")
	//添加到总公司
	root.Add(c1)
	root.Add(c2)
	//为北京公司新建2个部门
	rs1 := NewDepartment("人事部")
	cw1 := NewDepartment("财务部")
	//加入到北京公司
	c1.Add(rs1)
	c1.Add(cw1)
	//为上海公司新建一个部门
	cw2 := NewDepartment("财务部")
	c2.Add(cw2)
	//打印集团结构树
	root.Display(0)
	//OutPut:
	//华夏集团
	//--   人事部
	//--   华夏集团北京分公司
	//----   人事部
	//----   财务部
	//--   华夏集团上海分公司
	//----   财务部

}