//单例模式

package single

//意图：保证一个类仅有一个实例，并提供一个访问它的全局访问点
//解决:一个全局使用的类频繁地创建与销毁

//实例:利用go包sync 和 首字母小写对外不可见 实现一个单例模式

import (
	"net/http"
	"sync"
)

//实现一个全局唯一的http客户端
var (
	once sync.Once
	instance *http.Client
)

func New() *http.Client{
	once.Do(func() {
		instance = &http.Client{
			Timeout:30,//超时30s
			Transport:http.DefaultTransport,
		}
	})
	return instance
}


