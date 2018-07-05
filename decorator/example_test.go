package decorator

import "fmt"

func ExampleBottleMilk_Desc() {
	p := &Packing{name:"普通蛋糕",price:10.00} //普通蛋糕10元
	gp := &GlassPack{pack:p,name:"外加玻璃套",price:2.00} //加一个玻璃盒子包装加2元
	pgp := &PlasticPack{pack:gp,name:"外加塑料套",price:1.00} //加一个塑料盒子包装加1元
	gpgp := &GlassPack{pack:pgp,name:"外加玻璃套",price:2.00} //加一个玻璃盒子包装加2元

	fmt.Printf("名称:%s,价格:%f\n",gpgp.Desc(),gpgp.Price())
	//Output:
	//名称:普通蛋糕外加玻璃套外加塑料套外加玻璃套,价格:15.000000
}