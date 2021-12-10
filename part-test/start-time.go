package main

import (
	"fmt"
	"time"
)

func StartCac() {
	t1 := time.Now()//.Add(time.Microsecond) // get current time
	//logic handlers
	//for i := 0; i < 1000; i++ {
	//	fmt.Print("*")
	//}
	for i:=0;i<4 ;i++ {
		time.Sleep(time.Second *1 )
	}

	elapsed := time.Since(t1).Microseconds()
	fmt.Println("App elapsed: ", elapsed)

	t2:= elapsed%7
	if t2==0 {
		fmt.Println("@@0",t2)
	}else if t2==1 {
		fmt.Println("@@1",t2)
	}else if t2==2 {
		fmt.Println("@@02",t2)
	}else if t2==3 {
		fmt.Println("@@03",t2)
	}else if t2==4 {
		fmt.Println("@@04",t2)
	}else if t2==5 {
		fmt.Println("@@05",t2)
	}else  {
		fmt.Println("@@06",t2)
	}

	t:=time.Now().Nanosecond()



	fmt.Println(t)

}

func main(){
	StartCac()
}


