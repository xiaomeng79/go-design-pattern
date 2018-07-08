//单例模式2
package single

//意图：保证一个类仅有一个实例，并提供一个访问它的全局访问点
//解决:一个全局使用的类频繁地创建与销毁

//实例:利用go的init 实现一个单例模式

import "net/http"

var (
	instance2 *http.Client
)

func init() {
	instance2 = &http.Client{
		Timeout:30,//超时30s
		Transport:http.DefaultTransport,
	}
}
