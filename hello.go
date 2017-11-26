package main

import (
  "fmt"
  "os"
  "strings"
  "path/filepath"
)

type Person struct {
  name string
  address string
}

type IHuman interface {
  PrintName()
  PrintAddress()
}

type Tiger struct {
  size float64
  strength float64
  biteForce float64
  longivity float64
}

type IAnimal interface {
  PrintSize()
  PrintStrength()
  PrintBiteForce()
  PrintLongivity()
}

type ILiving interface{}

func (p Person) PrintName(){
  fmt.Println("Name is ",p.name)
}

func (p Person) PrintAddress(){
  fmt.Println("Address is ",p.address)
}

func (t Tiger) PrintSize(){
  fmt.Println("Size is ",t.size)
}

func (t Tiger) PrintStrength(){
  fmt.Println("Strength is", t.strength)
}

func (t Tiger) PrintBiteForce(){
  fmt.Println("BiteForce is ", t.biteForce)
}

func (t Tiger) PrintLongivity(){
  fmt.Println("Longivity is", t.longivity)
}

func Receive(i ILiving) {
  fmt.Println("Instance received successfully, Details are:")
} 

func main(){
  who:="World"
  if len(os.Args)==1{
    fmt.Println("usage: ",filepath.Base(os.Args[0]))
    os.Exit(1)
  }
  if len(os.Args)>1 {
    who+=" from "+strings.Join(os.Args[1:]," ")
  }
  fmt.Println("Hello ",who)
  PrintMsg("Util message")

  person:=Person{"abhi","ranchi"}
  person.PrintName()
  person.PrintAddress()
  
  tiger:=Tiger{6.5,3200,6500,25}
  tiger.PrintSize()
  tiger.PrintStrength()
  tiger.PrintBiteForce()
  tiger.PrintLongivity()

  Receive(person)
  Receive(tiger)
}



