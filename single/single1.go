//单例模式
//利用go包sync 和 首字母小写对外不可见 实现一个单例模式

package single

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


