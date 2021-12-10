package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/imroc/biu"
	"net"
	"strings"
	"time"
	"goweb-test/date-produce/Conf"
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

type IConfig interface {
	AddObserver()
	DelObserver()
	Notify()
}

type Iobserver interface {
	Update()
}


type Config struct {
	Pa []Iobserver

}


//send
type DateA struct {


}
func (this DateA) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}

type DateB struct {

}
func ( this DateB) Update()  {
	fmt.Println("配置已更新")
}

type DateC struct {

}
func (this DateC) Update() {
	fmt.Println("配置已更新")
}


func( this Config) Notify()  {
	for _,v :=range this.Pa {
		v.Update()
	}
}

func (this Config) AddObserver(iobserver Iobserver)  {
	this.Pa = append(this.Pa, iobserver)
}

func (this Config) RemoveObserver(iobserver Iobserver)  {
	for i,v := range this.Pa {
		if v == iobserver {
			this.Pa = append(this.Pa[:i], this.Pa[i+1:]...)
		}
	}
}

var (

)







func send(conn net.Conn ,a bytes.Buffer,o ,p string)  {
	for  {


		//fmt.Println("大小",len(a.Bytes()))
		n,err :=conn.Write(a.Bytes())

		if err!=nil{
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("客户端发送了%d字节的%s类%s数据\n",n,o,p)

		time.Sleep(500 * time.Millisecond)    //设置发送速率保证程序不崩溃
		return
	}
}


func  transform(pa interface{}) *bytes.Buffer {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(pa)
	if err != nil {
		panic(err)
	}

	return &buffer
}

func (Pa *Packages)  GetErrDate(start ,end int) []byte {

	var result []byte
	if start <0 || end > 3648 {
		return nil
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
	return Pa.Date

}

func main()  {
	//建立连接
	conn, err:=net.Dial("tcp", "127.0.0.1:8888")
	if err!=nil{
		fmt.Println("client dial err=",err)
		return
	}

	fconf:= Conf.Yaml{}

	cp:=make(chan Conf.Yaml,1024)

	//conf:=Conf.Yaml{}

	go func() {
		go 	Conf.GetYaml("e:\\conf.yaml",cp)
		for  {
			select {
			case fconf = <-cp:
				fmt.Println("配置已读取刷新")
			}
		}
	}()

	time.Sleep(3*time.Second)
	fmt.Println(fconf)


	PaGPS := Packages{MessageId: 4096,Data_length: 467,Heart_status: 0}
	PaGPS.Heart_status++
	PaGPS.Date=PaGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)



	PaIMU := Packages{MessageId: 4097,Data_length: 467,Heart_status: 0}
	PaIMU.Heart_status++
	PaIMU.Date=PaIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PaBD :=Packages{MessageId: 4098,Data_length: 467,Heart_status: 0}
	PaBD.Heart_status++
	PaBD.Date=PaBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PaGLONASS :=Packages{MessageId: 4099,Data_length: 467,Heart_status: 0}
	PaGLONASS.Heart_status++
	PaGLONASS.Date=PaGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PaADS := Packages{MessageId: 4100,Data_length: 467,Heart_status: 0}
	PaADS.Heart_status++
	PaADS.Date=PaADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PaACU :=Packages{MessageId: 4101,Data_length: 467,Heart_status: 0}
	PaACU.Heart_status++
	PaACU.Date=PaACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)


	PbGPS := Packages{MessageId: 8192,Data_length: 467,Heart_status: 0}
	PbGPS.Heart_status++
	PbGPS.Date=PbGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)

	PbIMU := Packages{MessageId: 8193,Data_length: 467,Heart_status: 0}
	PbIMU.Heart_status++
	PbIMU.Date=PbIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PbBD :=Packages{MessageId: 8194,Data_length: 467,Heart_status: 0}
	PbBD.Heart_status++
	PbBD.Date=PbBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PbGLONASS :=Packages{MessageId: 8195,Data_length: 467,Heart_status: 0}
	PbGLONASS.Heart_status++
	PbGLONASS.Date=PbGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PbADS := Packages{MessageId: 8196,Data_length: 467,Heart_status: 0}
	PbADS.Heart_status++
	PbADS.Date=PbADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PbACU :=Packages{MessageId: 8197,Data_length: 467,Heart_status: 0}
	PbACU.Heart_status++
	PbACU.Date=PbACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)




	PcGPS := Packages{MessageId: 12288,Data_length: 467,Heart_status: 0}
	PcGPS.Heart_status ++
	PcGPS.Date=PcGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)

	PcIMU := Packages{MessageId: 12289,Data_length: 467,Heart_status: 0}
	PcIMU.Heart_status ++
	PcIMU.Date=PcIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PcBD :=Packages{MessageId: 12290,Data_length: 467,Heart_status: 0}
	PcBD.Heart_status ++
	PcBD.Date=PcBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PcGLONASS :=Packages{MessageId: 12291,Data_length: 467,Heart_status: 0}
	PcGLONASS.Heart_status ++
	PcGLONASS.Date=PcGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PcADS := Packages{MessageId: 12292,Data_length: 467,Heart_status: 0}
	PcADS.Heart_status ++
	PcADS.Date=PcADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PcACU :=Packages{MessageId: 12293,Data_length: 467,Heart_status: 0}
	PcACU.Heart_status ++
	PcACU.Date=PcACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)

	time.Sleep(100*time.Microsecond)

	PaGPSbuffer:=transform(PaGPS)
	PaIMUbuffer:= transform(PaIMU)
	PaBDbuffer :=transform(PaBD)
	PaGLONASSbuffer:=transform(PaGLONASS)
	PaADSbuffer:=transform(PaADS)
	PaACUbuffer:=transform(PaACU)

	PbGPSbuffer:=transform(PbGPS)
	PbIMUbuffer:= transform(PbIMU)
	PbBDbuffer :=transform(PbBD)
	PbGLONASSbuffer:=transform(PbGLONASS)
	PbADSbuffer:=transform(PbADS)
	PbACUbuffer:=transform(PbACU)

	PcGPSbuffer:=transform(PcGPS)
	PcIMUbuffer:= transform(PcIMU)
	PcBDbuffer :=transform(PcBD)
	PcGLONASSbuffer:=transform(PcGLONASS)
	PcADSbuffer:=transform(PcADS)
	PcACUbuffer:=transform(PcACU)

	var A string = "A"
	var B string = "B"
	var C string = "C"

	var GPS string= "GPS"
	var IMU string = "IMU"
	var BD string ="BD"
	var GLONASS = "GLONSS"
	var ADS string = "ADS"
	var ACU string ="ACU"

	for  {

		time2:=time.NewTimer(1000 * time.Microsecond)

		<-time2.C

		send(conn,*PaGPSbuffer,A,GPS)

		send(conn,*PbGPSbuffer,B,GPS)

		send(conn,*PcGPSbuffer,C,GPS)

		time.Sleep(1000 *time.Microsecond)

		send(conn,*PaIMUbuffer,A,IMU)

		send(conn,*PbIMUbuffer,B,IMU)

		send(conn,*PcIMUbuffer,C,IMU)

		time.Sleep(1000 *time.Microsecond)

		send(conn,*PaBDbuffer,A,BD)

		send(conn,*PbBDbuffer,B,BD)

		send(conn,*PcBDbuffer,C,BD)

		time.Sleep(1000 *time.Microsecond)

		send(conn,*PaGLONASSbuffer,A,GLONASS)

		send(conn,*PbGLONASSbuffer,B,GLONASS)

		send(conn,*PcGLONASSbuffer,C,GLONASS)

		time.Sleep(1000 *time.Microsecond)

		send(conn,*PaADSbuffer,A,ADS)

		send(conn,*PbADSbuffer,B,ADS)

		send(conn,*PcADSbuffer,C,ADS)

		time.Sleep(1000 *time.Microsecond)

		send(conn,*PaACUbuffer,A,ACU)

		send(conn,*PbACUbuffer,B,ACU)

		send(conn,*PcACUbuffer,C,ACU)

		time.Sleep(1000 *time.Microsecond)


	}

}










