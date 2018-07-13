package flyweight

import "fmt"

func ExampleMysqlConnect_Do() {
	//初始化连接池
	mysqlpool := NewMysqlConnPool(2)
	//获取一个连接
	conn := mysqlpool.Get()
	//执行操作
	 fmt.Println(conn.Do())
	//放入连接池
	mysqlpool.Put(conn)
	//OutPut:
	//done
}
