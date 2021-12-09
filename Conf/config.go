package Conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
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

//故障位处理
//var start int
//var end int

//var ErrDate []byte



//func (Pa *Package)  GetErrDate(start ,end int,ErrDate []byte) []byte {
//
//	if start <0 || end > 3648 {
//		return nil
//	}
//	//var s [456]byte
//	s:=make([]byte,456)
//	for i:=start;i<end;i++ {
//		s[i]=byte(1)
//
//	}
//	//byteContent := "\x00"+ strings.Join(s, "\x02\x00")  // x20 = space and x00 = null
//	//fmt.Println(byteContent)
//
//	//fmt.Println(s)
//	copy(ErrDate,s[:])
//	//fmt.Println(ErrDate)
//	copy(Pa.Date,s)
//	//if start ==0  {
//	//	//Repeat 返回的是string类型，后面需要转byte[]类型
//	//	result1:=strings.Repeat("1",end)
//	//	result2:=strings.Repeat("0",3648-end)
//	//	//ErrDate1:=strings.Join([]string{result1,result2},"")
//	//
//	//	ErrDate=[]byte(result1+result2)
//	//	//ErrDate=[]byte(ErrDate1)
//	//
//	//}else if end == {
//	//
//	//}
//
//	return ErrDate
//}

//读取Yaml配置文件,
//并转换成conf对象
func GetYaml(fillpath string,cp chan Yaml) {

	fconf:= &Yaml{}
	t:= time.NewTicker(1*time.Second)
	for  {
		select {
		case <-t.C:
			//看配置是否存在
			_,err:= os.Stat(fillpath)
			if err != nil {
				fmt.Printf("fillpath not exist%#v\n",err)
			}
			//读外部的yaml配置文件
			yamlFile,err2:= ioutil.ReadFile(fillpath)
			if err2!=nil {
				fmt.Println(err2.Error())
			}
			//fmt.Println(yamlFile)
			err3:= yaml.Unmarshal(yamlFile,fconf)
			if err3 != nil {
				panic(err3)
			}

			cp <- *fconf

		}

	}
	////应该是 绝对地址
	//yamlFile, err := ioutil.ReadFile("conf.yaml")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = yaml.Unmarshal(yamlFile, cp)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

}



//func main() {
//
//
//
//
//
//
//
//
//
//}
