package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
)
type Packagess struct {
	Header1  uint32
	Header2	uint32
	MessageId	uint32
	Header3	uint32
	Header4 	uint32
	Data_length	uint32
	Health_status	uint32
	Heart_status	uint32
	Data	[]byte
	SVPC	uint32
	Transmit_offset	uint32
	Receive_offset	uint32
	Datepump	uint32
	VPC	uint32
	Data_CRC	uint32
}

func decode(Sbyte []byte)  {
	var pas =new(Packagess)

	//var b bytes.Buffer
	buf:=bytes.NewBuffer(Sbyte)  //构造出bytes.Buffer{}
	err := binary.Read(buf, binary.BigEndian, pas)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}



}



func process(conn net.Conn)  {
	//延迟关闭目前连接
	defer conn.Close()

	for{
		sbyte:=make([]byte,512)
		fmt.Printf("\n服务器在等待客户端%s 发送信息\n", conn.RemoteAddr().String())
		n, err:=conn.Read(sbyte)
		if err!=nil{
			fmt.Printf("客户端退出 err=%v\n",err)
			return
		}
		//var pas =new(Packagess)
		//var b bytes.Buffer
		b:=bytes.NewBuffer(sbyte)  //构造出bytes.Buffer{}
		pas:=Packagess{}
		binary.Read(b, binary.BigEndian, &pas.Header1)
		binary.Read(b, binary.BigEndian, &pas.Header2)
		binary.Read(b, binary.BigEndian, &pas.MessageId)
		binary.Read(b, binary.BigEndian, &pas.Header3)
		binary.Read(b, binary.BigEndian, &pas.Header4)
		binary.Read(b, binary.BigEndian, &pas.Data_length)
		binary.Read(b, binary.BigEndian, &pas.Health_status)
		binary.Read(b, binary.BigEndian, &pas.Heart_status)
		pas.Data=sbyte[32:488]
		//bytes.ReplaceAll(pas.Date,pas.Date[0],)
		//copy(pas.Date,sbyte[32:488])
		binary.Read(b, binary.BigEndian, pas.SVPC)
		binary.Read(b, binary.BigEndian, pas.Transmit_offset)
		binary.Read(b, binary.BigEndian, pas.Receive_offset)
		binary.Read(b, binary.BigEndian, pas.Datepump)
		binary.Read(b, binary.BigEndian, pas.VPC)
		binary.Read(b, binary.BigEndian, pas.Data_CRC)
		fmt.Println(sbyte[32])
		//fmt.Println("pas.data",pas.Data)
		fmt.Println("sbyte:",len(sbyte))
		//fmt.Println("pas",len(pas))
		fmt.Println(pas)
		//dec := gob.NewDecoder(b)
		//err=dec.Decode(pas)
		//if err != nil {
		//	panic(err)
		//}

		fmt.Println(n)
		//fmt.Printf("长度",len(b.Bytes()))
		//fmt.Println(b.Bytes())
		//fmt.Printf("%#v\n",pas)
		//fmt.Println(pas)


	}
}

func main(){
	fmt.Printf("服务器开始监听...\n")
	//监听一个端口
	listen, err:=net.Listen("tcp", ":8888")
	//延迟关闭连接
	if err!=nil{
		fmt.Println("listen err=",err)
		return
	}


	defer listen.Close()

	//循环等待客户端来连接
	for{
		fmt.Println("\n等待客户端来连接...\n")
		conn, err:=listen.Accept()
		//错误处理和输出当前连接客户端信息



		if err!=nil{
			fmt.Println("Accept() err=", err)
		}else{
			fmt.Printf("Accept() suc con=%v 客户端ip=%v\n", conn, conn.RemoteAddr().String())
		}

		//开一个协程专门处理当前连接的客户端
		go process(conn)
	}
}