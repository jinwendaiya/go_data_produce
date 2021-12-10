package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/willf/bitset"
)

func  GetErrDate1(start ,end int) []byte{
	rand.Seed(time.Now().UnixNano())

	b:=bitset.New(3647)
	//var k int



	//i<5  进行控制生成随机数位
	for i:=0;i<5;i++ {
		r := rand.Intn(end) + start
		fmt.Println(r)
		b.Set(uint(r))
	}






	y,_:=b.MarshalBinary()
	fmt.Println("故障位byte数据：",y)

	//100是start  ，900是end
	//r := rand.Intn(end) + start
	//t := rand.Intn(end) + start

	//fmt.Println(r)
	return y

}


func main() {


	GetErrDate1(1,3648)


}
