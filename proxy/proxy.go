//代理模式
package proxy

import "github.com/xiaomeng79/go-design-pattern/proxy/readfile"

//意图：为其他对象提供一种代理以控制对这个对象的访问
//解决:想在访问一个类时做一些[安全]控制或者直接访问对象很麻烦
//注意事项：
// 1、和适配器模式的区别：适配器模式主要改变所考虑对象的接口，而代理模式不能改变所代理类的接口。
// 2、和装饰器模式的区别：装饰器模式为了增强功能，而代理模式是为了加以控制

//实例:有一个访问本地磁盘文件的接口，但是现在需要权限认证后才可以访问
//实例2:也可以实现一个代理，可以读取本地文件，可以读取网络文件(未实现)

type IReadFileProxy interface {
	ReadFileContext(string,string) string
}

type ReadFileProxy struct {
}

func (r ReadFileProxy) auth(name string) bool {
	if name == "boos" {
		return true
	}
	return false
}

func (r ReadFileProxy) ReadFileContext(name,filename string) string {
	if !r.auth(name) {
		return "无权查看"
	}
	//代理去调用读文件接口读取
	read := readfile.ReadFile{}
	return read.ReadFile(filename)
}


