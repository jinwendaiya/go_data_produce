package main

import (
	"fmt"
	"github.com/imroc/biu"
	"strings"
)

type Packages struct {
	Header1  uint32   //4
	Header2	uint32
	MessageId	uint32
	Header3	uint32
	Header4 	uint32
	Data_length	uint32
	Health_status	uint32
	Heart_status	uint32
	Date	[]byte
	SVPC	uint32
	Transmit_offset	uint32
	Receive_offset	uint32
	Datepump	uint32
	VPC	uint32
	Data_CRC	uint32
}

//var s [456]byte
//s:=make([]string,3648)
//for i:=start;i<end;i++ {
//
//	s[i]=string(1)
//}
//s1:=make([]byte,456)
//
//s1=bytes.Repeat()


func (Pa *Packages)  GetErrDate(start ,end int)  {
	var result []byte
	if start <0 || end > 3648 {
		return
	}

	if start ==0 {
		front := strings.Repeat("1",end)
		rear:= strings.Repeat("0",3647-end)
		result =biu.BinaryStringToBytes(front+rear)
	} else if end == 3647 {
		front := strings.Repeat("0",start)
		rear := strings.Repeat("1",3647-start)
		result = biu.BinaryStringToBytes(front+rear)


	}else {
		front:=strings.Repeat("0",start)
		mid:=strings.Repeat("1",end-start)
		rear:=strings.Repeat("0",3647-end)
		result = biu.BinaryStringToBytes(front+mid+rear)
	}

	Pa.Date=result
	fmt.Println(Pa.Date)
	//return s
	//return result
}


func main() {


	//fconf:= Conf.Yaml{}




	PaGPS := Packages{MessageId: 4096,Data_length: 467,Heart_status: 0}
	PaGPS.Heart_status++
	PaGPS.GetErrDate(0,3000)



}
