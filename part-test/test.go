package main

import (
	"fmt"
	"goweb-test/date-produce/Conf"
	"time"

	//"goweb-test/date-produce/Conf"
	//"strings"
	//"time"

)

//
//type Packages struct {
//	a  uint32
//	Fir_exist	[456]byte
//}


//故障位处理
//var start int
//var end int



func main() {

	//pa:=new(Packages)
	//
	//pa.a=4096
	//
	//fmt.Println(pa)



	cp:=make(chan Conf.Yaml)

	conf:=Conf.Yaml{}

	go func() {
		go 	Conf.GetYaml("e:\\conf.yaml",cp)
		for  {
			select {
			case conf = <-cp:
			//	fmt.Println("配置已读取刷新")
			}
		}
	}()
	t1:=time.NewTicker(2*time.Second)

	for {
		select {
		case <-t1.C:
			fmt.Println(conf)
		//	fmt.Println(conf.ACU.Start)

		}
	}
	//fmt.Println(conf.ACU.Start)

	//Conf.GetErrDate(4,69)

	//time.Sleep(1*time.Second)









}
