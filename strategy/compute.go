//选择不同的策略计算价格
package strategy

import (
	"github.com/xiaomeng79/go-design-pattern/strategy/strate"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//价格计算
type PriceCompute struct {
	OriginalPrice float64          //原价
	PromotionType string           //促销类型字符串：如：几折 满几减几
	strate        strate.IStrategy //策略
}

//reg1 := `(^[0-9].*)折`
//reg2 := `^满([\d]+.?[\d]*)减([\d]+.?[\d]*)`

//由于浮点运算不准确，所以返回2位小数
//返回计算后的价格
func (p *PriceCompute) Do() float64 {
	r1 := regexp.MustCompile(`(^[0-9].*)折`)                     //打折 ：8.8折
	r2 := regexp.MustCompile(`^满([\d]+.?[\d]*)减([\d]+.?[\d]*)`) //满减 ： 满200.35减100
	p.PromotionType = strings.Replace(p.PromotionType, " ", "", -1)
	switch {
	case r1.MatchString(p.PromotionType): //打折
		s := r1.FindStringSubmatch(p.PromotionType)
		f, _ := strconv.ParseFloat(s[1], 64)
		p.strate = &strate.Discount{p.OriginalPrice, f}

	case r2.MatchString(p.PromotionType): //满减
		s := r2.FindStringSubmatch(p.PromotionType)
		f1, _ := strconv.ParseFloat(s[1], 64)
		f2, _ := strconv.ParseFloat(s[2], 64)
		p.strate = &strate.FullSub{p.OriginalPrice, f1, f2}
	default:
		return p.OriginalPrice //没有匹配的促销模式
	}
	return Round(p.strate.Compute(), 2) //取2位小数

}

//计算结果取小数的n位
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
