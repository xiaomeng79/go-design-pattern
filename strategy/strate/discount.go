package strate

//折扣

type Discount struct {
	OriginalPrice float64 //原价
	Discount float64//折扣
}

//计算
func (d *Discount) Compute() float64 {
	if d.Discount <= 0 {
		return  d.OriginalPrice
	}
	return d.OriginalPrice * d.Discount/10
}
