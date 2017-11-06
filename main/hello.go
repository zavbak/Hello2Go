package main

import (
	"fmt"
	"time"
)

var g int

func test(i...int) (res int, r bool){

	for _,v := range i{
		fmt.Println(v)
		res+=1;
	}

	r = true

	return
}

func printSeparator(name string){

	var s string

	for i:=0; i<10; i++{
		s +="-"
	}
	fmt.Println()

	fmt.Println(s," ", name," ", s)
}

func main() {

	//var sl []int;
	//sl = append(sl, 100)
	//sl = append(sl, 105)
	//sl = append(sl, 105)
	//
	//fmt.Println(sl, len(sl))
	//
	//s2 := []int{1, 2, 5}
	//
	//fmt.Println(s2, s2[0], len(s2))
	//
	//s2 = append(s2, sl...)
	//
	//fmt.Println(s2, s2[0], len(s2))
	//
	//var slm [][]int
	//slm = append(slm, s2)
	//
	//fmt.Println(slm, slm[0], len(slm))
	//
	//slmake := make([]int, 40)
	//fmt.Println(slmake)
	//
	//nm := map[string]string{
	//	"h1": "Hello",
	//	"h2": "world",
	//}
	//
	//fmt.Println(nm)
	//
	//hello, ok := nm["h1"]
	//fmt.Println(hello, ok)
	//
	//fmt.Println("_____________________")
	//for key,val :=range nm{
	//	fmt.Println(key, val)
	//}
	//
	////fmt.Println(Hello1)
	//
	//
	//fmt.Println("*********Pointer*****************")
	//
	//a:=5
	//b:=&a
	//
	//*b=4
	//fmt.Println(a,b)

    // 1 *************************************************************************************************************
	printSeparator("func 1")

	//sl3 :=[]int{1,2,3}
	res,t := test(3,2,4,5)

	fmt.Println("mount: ", res,t)


	// 2 *************************************************************************************************************
	printSeparator("init")

	fmt.Println("init g = ", g)

	// 3 *************************************************************************************************************

	printSeparator("Замыкания")


	mTimer1 := getTimer("f1")
	mTimer2 := getTimer("f2")
	//defer  mTimer1()

	f1:= func() {
		mTimer1()
	}


	f2 := func(){
		mTimer2()
	}

	f1()
	f2()
	f1()
	f1()
	f2()
	f2()
	f1()



}

func getTimer(name string) func()  {
	start := time.Now()
	count :=0

	return func() {
		count++
		fmt.Printf("\n name %v  time %v count %v",name,time.Since(start),count)
	}
}

func init()  {
	g = 5
}
