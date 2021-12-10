package main

import (
	"fmt"
	"goweb-test/date-produce/Conf"
	"time"
)

type Package struct {
	Date []byte
}





type Yaml struct {

	GPS struct {
		Start int `yaml:"start"`
		End int`yaml:"end"`
	}

	IMU struct {
		Start int `yaml:"start"`
		End int`yaml:"end""`
	}

	BD struct {
		Start int `yaml:"start"`
		End int`yaml:"end"`
	}
	GLONASS struct {
		Start int `yaml:"start"`
		End int`yaml:"end"`
	}
	ADS struct {
		Start int `yaml:"start"`
		End int`yaml:"end"`
	}
	ACU struct {
		Start int `yaml:"start"`
		End int`yaml:"end"`
	}
}
func (Pa *Package)  GetErrDate(start ,end int) []byte {

	if start <0 || end > 3648 {
		return nil
	}
	//var s [456]byte
	s:=make([]byte,456)
	for i:=start;i<end;i++ {
		s[i]=byte(1)

	}
	//byteContent := "\x00"+ strings.Join(s, "\x02\x00")  // x20 = space and x00 = null
	//fmt.Println(byteContent)
	fmt.Println(len(s))
	//fmt.Println(s)
	//copy(ErrDate,s[:])
	//fmt.Println(ErrDate)
	//copy(Pa.Date,s)
	Pa.Date=s
	fmt.Println(Pa.Date)
	//if start ==0  {
	//	//Repeat 返回的是string类型，后面需要转byte[]类型
	//	result1:=strings.Repeat("1",end)
	//	result2:=strings.Repeat("0",3648-end)
	//	//ErrDate1:=strings.Join([]string{result1,result2},"")
	//
	//	ErrDate=[]byte(result1+result2)
	//	//ErrDate=[]byte(ErrDate1)
	//
	//}else if end == {
	//
	//}

	return s
}



func main() {

	//fconf:= &Yaml{}

	//cp:=make(chan Conf.Yaml)

	//Conf.GetYaml("e:\\conf.yaml",cp)


	//yamlFile,_:= ioutil.ReadFile("e:\\conf.yaml")
	////yamlFile,_:= ioutil.ReadFile("e:\\aaa.TXT")
	////fmt.Println(string(yamlFile))
	//
	//yaml.Unmarshal(yamlFile,fconf)
	//
	//fmt.Println(fconf)

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








	//var cp chan Yaml
	//
	//cp <- *fconf
	//ErrDate:=make([]byte,456)

	//Conf.GetErrDate(fconf.ACU.Start,fconf.ACU.End,ErrDate)

	time.Sleep(5*time.Second)

	//fmt.Println(ErrDate)

	Pa:=&Package{}

	//ticker1:=time.NewTicker(2*time.Second)

	//
	//for  {
	//	select {
	//	case <-ticker1.C:
	//
	//	 	Pa.GetErrDate(conf.ACU.Start,conf.ACU.End)
	//	}
	//}

 	// for  {
	//	time.Sleep(1*time.Second)
	//
	//go 	Pa.GetErrDate(conf.ACU.Start,conf.ACU.End)
	//
	//}



	Pa.GetErrDate(conf.ACU.Start,conf.ACU.End)

	//fmt.Println(Pa.Date)
	time.Sleep(1*time.Second)
}
