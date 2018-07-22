//备忘录模式
package memento

import (
	"encoding/gob"
	"fmt"
	"os"
)

//意图：保存程序内部状态到外部，又不希望暴露内部状态的情形,可以离线保存内部状态，如保存到数据库，文件等
//解决:捕获一个对象的内部状态，并在该对象之外保存这个状态，这样可以在以后将对象恢复到原先保存的状态。
//使用场景： 游戏存副本等

//实例:游戏状态恢复

//游戏备份
type Memento interface {
	Save(*Game)  //存档
	Load() *Game //载入
}

//游戏
type Game struct {
	Hp int
	Mp int
}

//初始化游戏
func GameInit() *Game {
	return &Game{}
}

//玩游戏
func (g *Game) Play(mpDelta, hpDelta int) {
	g.Mp += mpDelta
	g.Hp += hpDelta
}

//游戏状态
func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.Hp, g.Mp)
}

//游戏备份
type GameMemento struct {
}

func (gm *GameMemento) Save(game *Game) {
	f, err := os.Create("game.bak")
	if err != nil {
		panic(err)
	}
	enc := gob.NewEncoder(f)
	err = enc.Encode(game)
	if err != nil {
		panic(err)
	}
}

func (gm *GameMemento) Load() *Game {
	game := &Game{}
	f, err := os.Open("game.bak")
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(f)
	err = dec.Decode(game)
	if err != nil {
		panic(err)
	}
	return game
}
