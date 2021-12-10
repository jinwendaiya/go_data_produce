package main

import "fmt"

type Subject interface {
	AddObserver(observer Observer)
	RemoveObserver(observer Observer)
	Notify()
}

type Observer interface {
	Update()
}

type Weather struct {
	observers []Observer
}

func(this *Weather) AddObserver(observer Observer)  {
	this.observers = append(this.observers, observer)
}

func(this *Weather) RemoveObserver(observer Observer)  {
	for i,v := range this.observers {
		if v == observer {
			this.observers = append(this.observers[:i], this.observers[i+1:]...)
		}
	}
}

func(this *Weather) Notify()  {
	for _,v :=range this.observers {
		v.Update()
	}
}

type Ob1 struct {}
func(this Ob1) Update()  {
	fmt.Println("this is Ob1")
}

type Ob2 struct {}
func(this Ob2) Update()  {
	fmt.Println("this is Ob2")
}

type Ob3 struct {}
func(this Ob3) Update()  {
	fmt.Println("this is Ob3")
}
var (
	Sub Subject
	ob1, ob2, ob3 Observer
)
func main()  {
	Sub, ob1, ob2, ob3 =  &Weather{}, Ob1{}, Ob2{},Ob3{}
	Sub.AddObserver(ob1)
	Sub.AddObserver(ob2)
	Sub.AddObserver(ob3)
	Sub.RemoveObserver(ob2)
	Sub.Notify()
}
