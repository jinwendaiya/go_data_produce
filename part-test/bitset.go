package main

import (
	"fmt"
	"github.com/willf/bitset"
)

//replace github.com/bits-and-blooms/bitset v1.2.1 =>  github.com/willf/bitset v1.1.11

//var  b bitset.BitSet

func main()  {
	//var b bitset.BitSet // 定义一个BitSet对象
	//fmt.Println(b.Bytes())
	//b.Set(0).Set(2).Set(5)
	//fmt.Println(b.Bytes(),0,2,5)
	//x,_:=b.MarshalBinary()  //byte[]
	//fmt.Println(x)


	cc:=bitset.New(3647)

	cc.Set(192).Set(3646)

	fmt.Println("cc-byte:",cc.Bytes())

	y,_:=cc.MarshalBinary()


	fmt.Println("----转换后",y)

	fmt.Println("lens",len(y))

	//fmt.Println("chang :"cap(y ))
	//b.Set(10) // 给这个set新增两个值10
	//fmt.Println(b.Bytes(),0,10)
	//b.Set(64)
	//fmt.Println(b.Bytes(),0,10,64)
	//fmt.Println(b.Bytes(),63)

	//z:=b.Test(1000)
	//fmt.Println(z)  //布尔
	//y,_:=b.MarshalBinary()
	//fmt.Println(y)
	//if b.Test(1000) { // 查看set中是否有1000这个值（我觉得Test这个名字起得是真差劲，为啥不叫Exist）
	//	b.Clear(1000) // 情况set
	//}
}










//
//func main() {
//	fmt.Printf("Hello from BitSet!\n")
//	var b bitset.BitSet
//	// play some Go Fish
//	for i := 0; i < 100; i++ {
//		card1 := uint(rand.Intn(52))
//		card2 := uint(rand.Intn(52))
//		b.Set(card1)
//		if b.Test(card2) {
//			fmt.Println("Go Fish!")
//		}
//		b.Clear(card1)
//	}
//
//	// Chaining
//	b.Set(10).Set(11)
//
//	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
//		fmt.Println("The following bit is set:", i)
//	}
//	if b.Intersection(bitset.New(100).Set(10)).Count() == 1 {
//		fmt.Println("Intersection works.")
//	} else {
//		fmt.Println("Intersection doesn't work???")
//	}
//}
//func main() {
//	 var start int
//	//s:=make([]byte,1024)
//	var s []byte
//	for i:=0 ;i<3648 ;i++{
//		if i==start {
//			b.Set(1)
//		}
//
//	}
//	s=b.safeSet()
//b.MarshalBinary()
//
//}
