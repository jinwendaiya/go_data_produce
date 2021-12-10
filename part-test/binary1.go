package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)


type date struct {
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

//func  transform(pa interface{}) *bytes.Buffer {
//	//buffer := new(bytes.Buffer[])
//	var buffer bytes.Buffer
//	encoder := gob.NewEncoder(&buffer)
//	err := encoder.Encode(pa)
//	if err != nil {
//		panic(err)
//	}
//	//err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
//	//if err != nil {
//	//	panic(err)
//	//}
//	return &buffer
//}

func binarytestwrite(pa interface{})  *bytes.Buffer {
	//buf := new(bytes.Buffer)
	//pi := math.Pi
	var buf bytes.Buffer

	err := binary.Write(&buf, binary.BigEndian, pa)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(buf.Bytes())
	return &buf
}

func binaryreadtest()  {
	var pi float64
	b := []byte{0x18,0x2d,0x44,0x54,0xfb,0x21,0x09,0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.BigEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
}

func main() {
	//var pi float64
	//b := []byte{0x18,0x2d,0x44,0x54,0xfb,0x21,0x09,0x40}
	//buf := bytes.NewBuffer(b)
	//err := binary.Read(buf, binary.LittleEndian, &pi)
	//if err != nil {
	//	log.Fatalln("binary.Read failed:", err)
	//}
	//fmt.Println(pi)
	PaGPS := date{MessageId: 4096,Data_length: 467,Heart_status: 0}

	PaGPSbuffer:=binarytestwrite(PaGPS)

	fmt.Println("byte",PaGPSbuffer.Bytes())

	fmt.Println("len:",len(PaGPSbuffer.Bytes()))


}

