//单例模式2 利用go的init
package single

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
