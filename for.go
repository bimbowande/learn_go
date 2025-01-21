package main

import "fmt"

func main(){
	i:=1
	for i<=3{
		fmt.Println(i)
		i = i + 1
	}

	// for j:=0 ; j<3; j++ {
	// 	fmt.Println(j)
	// }

	for i:= range 8 {
		fmt.Println("range",i)
	}

	for i,c := range  "go"{
		fmt.Println(i,c)
	}
	for{
		fmt.Println("loop")
		break
	}

	for n := range 8 {
		if(n > 0){
			if (n%2 == 0 ){
				continue
			}
			fmt.Println(n)
		}
	}
}
