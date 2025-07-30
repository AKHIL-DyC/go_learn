package main

import ("fmt"

"runtime")
var abc=4 
var abd int //abd will have 0
var reality bool //default 0->false
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

	count:=1
	for{
		if i:="help";count>80{
			fmt.Println(i)
			break;
		}
		//i=i+1 cant access variables initialised in 'if' outside 'if'
		count=count*2
		
	}

	fmt.Println(count)


	//swith  case

	var os=runtime.GOOS;
	switch os{
	case "darwin":
		fmt.Println("darwin")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Println("eew windows?")		
	}

//switch{case:}at as ifelse if else else

defer fmt.Println("akhil");//when defer is found it is not executed rather pushed to a stack ,then after whole function is over ,one by one popped 
defer fmt.Println("is");
fmt.Println("my name");

p:=10
q:=&p
fmt.Println((*q))//prints 10




type coordinates struct{
	x int
	y int
}
v:=coordinates{1,2}//{}if empty implicit values taken
fmt.Println(v.x)
	
}
