package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/willf/bitset"
	"goweb-test/date-produce/Conf"
	"log"

	//"log"
	"math/rand"
	"net"
	"time"
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

var start ,end int

func send(conn net.Conn ,a bytes.Buffer,o ,p string)  {
	for  {


		fmt.Println("大小",len(a.Bytes()))
		n,err :=conn.Write(a.Bytes())

		if err!=nil{
			fmt.Println("conn.Write err=", err)
		}
		fmt.Printf("客户端发送了%d字节的%s类%s数据\n",n,o,p)

		time.Sleep(500 * time.Millisecond)    //设置发送速率保证程序不崩溃
		return
	}
}


func (pa Packages) transform(messageid  uint32) *bytes.Buffer {

	pa.MessageId = messageid
	pa.Data_length = 467
	pa.Heart_status=0
	pa.Heart_status++
	//pa.Heart_status++
	//pa.Date=pa.GetErrDate(start,end)

//r:=make([]byte,0)
//buf:=bytes.NewBuffer(r)
	var buf bytes.Buffer

	binary.Write(&buf,binary.BigEndian,pa.Header1)
	binary.Write(&buf,binary.BigEndian,pa.Header2)
	binary.Write(&buf,binary.BigEndian,pa.MessageId)
	binary.Write(&buf,binary.BigEndian,pa.Header3)
	binary.Write(&buf,binary.BigEndian,pa.Header4)
	binary.Write(&buf,binary.BigEndian,pa.Data_length)
	binary.Write(&buf,binary.BigEndian,pa.Health_status)
	binary.Write(&buf,binary.BigEndian,pa.Heart_status)
	buf.Write(pa.Date)
	binary.Write(&buf,binary.BigEndian,pa.SVPC)
	binary.Write(&buf,binary.BigEndian,pa.Transmit_offset)
	binary.Write(&buf,binary.BigEndian,pa.Receive_offset)
	binary.Write(&buf,binary.BigEndian,pa.Datepump)
	binary.Write(&buf,binary.BigEndian,pa.VPC)
	binary.Write(&buf,binary.BigEndian,pa.Data_CRC)

	return &buf
}

func (Pa *Packages)  GetErrDate(start ,end int) []byte {

	b:=bitset.New(3647)
	rand.Seed(time.Now().UnixNano())

	if start <0 || end > 3647 {
		return nil
	}

	if start==0 && end==3647{
		for i :=0;i<3648;i++ {
			b.Set(uint(i))
		}
	}else if start>0 && end < 3648 {
		//i<5  进行控制生成随机数位
		for i:=0;i<5;i++ {
			r := rand.Intn(end) + start
			fmt.Println(r)
			b.Set(uint(r))
		}

	}
	y,_:=b.MarshalBinary()
	//注意更改writeto方法，否则包大小有问题


	Pa.Date=y
	fmt.Println(Pa.Date)
	fmt.Println("len:",len(Pa.Date))
	return Pa.Date

}

func notify(transferconf,fconf Conf.Yaml)  {

	fconf=transferconf

}

func main()  {
	//建立连接
	conn, err:=net.Dial("tcp", "127.0.0.1:8888")
	if err!=nil{
		fmt.Println("client dial err=",err)
		return
	}

	transferconf:=Conf.Yaml{}

	fconf:= Conf.Yaml{}

	cp:=make(chan Conf.Yaml,1024)

	//conf:=Conf.Yaml{}

	go func() {
		go 	Conf.GetYaml("e:\\conf.yaml",cp)
		for  {
			select {
			case transferconf = <-cp:
				fmt.Println("配置已读取刷新")
			}
		}
	}()

	go func() {
		//创建一个监控对象
		watch, err := fsnotify.NewWatcher();
		if err != nil {
			log.Fatal(err);
		}
		defer watch.Close();
		//添加要监控的对象，文件或文件夹
		err = watch.Add("e:\\conf.yaml");
		if err != nil {
			log.Fatal(err);
		}
		//我们另启一个goroutine来处理监控对象的事件
		go func() {
			for {
				select {
				case ev := <-watch.Events:
					{
						//判断事件发生的类型，如下5种
						// Create 创建
						// Write 写入
						// Remove 删除
						// Rename 重命名
						// Chmod 修改权限
						//if ev.Op&fsnotify.Create == fsnotify.Create {
						//	log.Println("创建文件 : ", ev.Name);
						//}
						if ev.Op&fsnotify.Write == fsnotify.Write {
							log.Println("写入文件 : ", ev.Name);
							notify(transferconf,fconf)

						}
						if ev.Op&fsnotify.Remove == fsnotify.Remove {
							log.Println("删除文件 : ", ev.Name);
						}
						//if ev.Op&fsnotify.Rename == fsnotify.Rename {
						//	log.Println("重命名文件 : ", ev.Name);
						//}
						//if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						//	log.Println("修改权限 : ", ev.Name);
						//}
					}
				case err := <-watch.Errors:
					{
						log.Println("error : ", err);
						return;
					}
				}
			}
		}();
		fmt.Println("监测中...")
		//循环
		select {};


	}()


	time.Sleep(3*time.Second)
	fmt.Println(fconf)

	//var PaGPS = &Packages{}
	//PaGPS.MessageId=4096
	//PaGPS.Data_length=467
	//PaGPS.Heart_status=0
	//
	//PaGPS.Heart_status++  //可能要单独放到函数内
	//PaGPS.Date=PaGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)

	//a:=make([]byte,456)
	PaGPS := Packages{}
	//PaGPS := Packages{MessageId: 4096,Data_length: 467,Heart_status: 0}
	//PaGPS.Heart_status++
	//PaGPS.Date=a
 	PaGPS.Date=PaGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)


	PaIMU := Packages{}
	PaIMU.Date=PaIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PaBD :=Packages{}
	PaBD.Date=PaBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PaGLONASS :=Packages{}
	PaGLONASS.Date=PaGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PaADS := Packages{}
	PaADS.Date=PaADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PaACU :=Packages{}
	PaACU.Date=PaACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)


	PbGPS := Packages{}
	PbGPS.Date=PbGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)

	PbIMU := Packages{}
	PbIMU.Date=PbIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PbBD :=Packages{}
	PbBD.Date=PbBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PbGLONASS :=Packages{}
	PbGLONASS.Date=PbGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PbADS := Packages{}
	PbADS.Date=PbADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PbACU :=Packages{}
	PbACU.Date=PbACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)



	PcGPS := Packages{}
	PcGPS.Date=PcGPS.GetErrDate(fconf.GPS.Start,fconf.GPS.End)

	PcIMU := Packages{}
	PcIMU.Date=PcIMU.GetErrDate(fconf.IMU.Start,fconf.IMU.End)

	PcBD :=Packages{}
	PcBD.Date=PcBD.GetErrDate(fconf.BD.Start,fconf.BD.End)

	PcGLONASS :=Packages{}
	PcGLONASS.Date=PcGLONASS.GetErrDate(fconf.GLONASS.Start,fconf.GLONASS.End)

	PcADS := Packages{}
	PcADS.Date=PcADS.GetErrDate(fconf.ADS.Start,fconf.ADS.End)

	PcACU :=Packages{}
	PcACU.Date=PcACU.GetErrDate(fconf.ACU.Start,fconf.ACU.End)

	time.Sleep(100*time.Microsecond)

	//var P1 = Packages{}

	PaGPSbuffer:=PaGPS.transform(4096)
	PaIMUbuffer:= PaIMU.transform(4097)
	PaBDbuffer :=PaBD.transform(4098)
	PaGLONASSbuffer:=PaGLONASS.transform(4099)
	PaADSbuffer:=PaADS.transform(4100)
	PaACUbuffer:=PaACU.transform(4101)

	PbGPSbuffer:=PbGPS.transform(8192)
	PbIMUbuffer:= PbIMU.transform(8193)
	PbBDbuffer :=PbBD.transform(8194)
	PbGLONASSbuffer:=PbGLONASS.transform(8195)
	PbADSbuffer:=PbADS.transform(8196)
	PbACUbuffer:=PbACU.transform(8197)

	PcGPSbuffer:=PcGPS.transform(12288)
	PcIMUbuffer:= PcIMU.transform(12289)
	PcBDbuffer :=PcBD.transform(12290)
	PcGLONASSbuffer:=PcGLONASS.transform(12291)
	PcADSbuffer:=PcADS.transform(12292)
	PcACUbuffer:=PcACU.transform(12293)

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










