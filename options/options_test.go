package options

import (
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestNewLog(t *testing.T) {
	//新建一个日志，并修改默认选项
	l := NewLog(
		WithPrefix("noprefix"),
	)
	assert.Equal(t,"noprefix",l.Prefix(),"prefix not equal")
	//flag 使用默认值
	assert.Equal(t,log.Ltime,l.Flags(),"flag not equal")
}
