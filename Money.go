package main

import "fmt"

func main(){
	var rupiah float32 = 15.518
	var dollar float32 = 10

	convert := dollar * rupiah
	fmt.Printf("10 Dollar sama dengan Rp %f \n", convert)

}