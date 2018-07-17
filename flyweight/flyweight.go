//享元模式
package flyweight


//意图：享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元，使多个对象共享，从而节省内存以及减少对象数量
//解决:在有大量对象时，有可能会造成内存溢出，我们把其中共同的部分抽象出来，如果有相同的业务请求，直接返回在内存中已有的对象，避免重新创建

//实例:数据库的连接池

//DB接口
type IDBConnect interface {
	Do() string
}
//连接池接口
type IDBPool interface {
	Get() IDBConnect
	Put(IDBConnect)
}

//数据库连接
type MysqlConnect struct {
}
//数据库操作
func (db *MysqlConnect) Do() string {
	return "done"
}


//连接池实例 mysql
type MysqlConnPool struct {
	ConnChan chan *MysqlConnect
}
//初始化连接池
func NewMysqlConnPool(num int) *MysqlConnPool{
	return &MysqlConnPool{ConnChan:make(chan *MysqlConnect,num)}
}
//获取连接
func (p *MysqlConnPool) Get() IDBConnect {
	select {
	case conn := <-p.ConnChan:
		return conn
	default:
		return new(MysqlConnect)
	}
}
//存入连接
func (p *MysqlConnPool) Put(conn IDBConnect) {
	select {
	case p.ConnChan <- conn.(*MysqlConnect):
		return
	default:
		//丢弃
		return
	}
}

