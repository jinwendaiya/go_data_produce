package sub

import (
	"fmt"
)

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
//A
type DateAGPS struct {

}
func (this DateAGPS) Update(){
//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateAIMU struct {

}
func (this DateAIMU) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateABD struct {

}
func (this DateABD) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateADateBGLASS struct {

}
func (this DateADateBGLASS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateAADS struct {

}
func (this DateAADS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateAACU struct {

}
func (this DateAACU) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}


//B
type DateBGPS struct {


}
func ( this DateBGPS) Update()  {
	fmt.Println("配置已更新")
}
type DateBIMU struct {

}
func (this DateBIMU) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateBBD struct {

}
func (this DateBBD) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateBDateBGLASS struct {

}
func (this DateBDateBGLASS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateBADS struct {

}
func (this DateBADS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateBACU struct {

}
func (this DateBACU) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}



//c
type DateCGPS struct {

}
func (this DateCGPS) Update() {
	fmt.Println("配置已更新")
}
type DateCIMU struct {

}
func (this DateCIMU) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateCBD struct {

}
func (this DateCBD) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateCGLASS struct {

}
func (this DateCGLASS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateCADS struct {

}
func (this DateCADS) Update(){
	//	GetErrDate1
	fmt.Println("配置已更新")
}
type DateCACU struct {

}
func (this DateCACU) Update(){
	//	GetErrDate1
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

func main()  {








	//Sub, ob1, ob2, ob3 =  &Weather{}, Ob1{}, Ob2{},Ob3{}
	//Sub.AddObserver(ob1)
	//Sub.AddObserver(ob2)
	//Sub.AddObserver(ob3)
	//Sub.RemoveObserver(ob2)
	//Sub.Notify()
}
