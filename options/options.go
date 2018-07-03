//选项模式
package options

import (
	"log"
	"io"
	"os"
)

//由于go语言没有默认值，继承，多态，所有使用选项模式来达到默认值，多态的效果

//选项模式 配置一个日志
type Option func(*Options)

type Options struct {
	Out io.Writer //日志输出
	Prefix string //日志前缀
	Flag int //日志标记
}

//设置输出
func WithOut(out io.Writer) Option {
	return func(o *Options){
		o.Out = out
	}
}
//设置前缀
func WithPrefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}
//设置标记
func WithFlag(flag int) Option {
	return func(o *Options) {
		o.Flag = flag
	}
}

//新建一个日志
func NewLog(opts ...Option) *log.Logger {
	//设置默认选项
	o := &Options{
		Out:os.Stderr,
		Prefix:"",
		Flag:log.Ltime,
	}
	//使用自定义选项覆盖默认选项
	for _,opt := range opts {
		opt(o)
	}
	return log.New(o.Out,o.Prefix,o.Flag)
}