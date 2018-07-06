package strate

//满减
type FullSub struct {
	OriginalPrice float64 //原价
	Full float64//满
	Sub float64//减
}

//计算
func (f *FullSub) Compute() float64 {
	//原价小于要进行满减的要求
	if f.OriginalPrice < f.Full {
		return  f.OriginalPrice
	}
	return f.OriginalPrice - f.Sub
}