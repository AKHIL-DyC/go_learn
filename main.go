package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x,y string)(string,string){//x and y time same so we can do like this
	return y,x
}
func split(sum int)(x,y int){//when we do this x and y are initialised with 0,no need to return in this ,x,y returned automatically
	x=sum*10
	y=sum/10 //we use = ,because = is for assigning 
	return
}
func main() {
	sum:=40//we do this := to initialize  //inside a function we dont need to mention type but in package we need to declare like var age=40(no need for :=)
	fmt.Println(add(42, 13))
	a,b:=swap("hello","world")
	fmt.Println(a,b)

	t,y:=split(sum)
	fmt.Println(t,y)

}
