package proxy

import "fmt"

func ExampleReadFileProxy_ReadFileContext() {
	proxy := ReadFileProxy{}
	c1 := proxy.ReadFileContext("boy", "hello")
	c2 := proxy.ReadFileContext("boos", "hello")
	fmt.Println(c1)
	fmt.Println(c2)
	//OutPut:
	//无权查看
	//文件内容为：保密内容
}
