package main

import "fmt"

func getPrime(n int) string{
	if(n<1){
		return "not prime"
	}
	for i:=2; i<n; i++{
		if(n%i == 0){
			return "not prime"
		}
	}
	return "it is prime"
}

func main() {
	fmt.Println(getPrime(13));
}
