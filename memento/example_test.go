package memento

func ExampleGame_Play() {
	//初始化游戏
	g := GameInit()
	//查看游戏状态
	g.Status()
	//玩游戏
	g.Play(5, 5)
	//查看游戏状态
	g.Status()
	//存档
	gm := &GameMemento{}
	gm.Save(g)
	//初始化游戏
	g = GameInit()
	//查看游戏状态
	g.Status()
	//恢复
	g = gm.Load()
	//查看游戏状态
	g.Status()

	//OutPut:
	//Current HP:0, MP:0
	//Current HP:5, MP:5
	//Current HP:0, MP:0
	//Current HP:5, MP:5

}
