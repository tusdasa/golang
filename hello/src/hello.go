package main

import (
	"fmt"
	"time"
)

func main(){
  var a int
  var b int = 100
  var c  string = "abc"
  d := 100
  fmt.Println("a=",a)
  fmt.Println("a=%T,b=%T,c=%s,d=%T",a,b,c,d)
	fmt.Println("hello Go!")
	time.Sleep(1*time.Second)
	fmt.Println("end")
  fmt.Println("Ok!")
}
