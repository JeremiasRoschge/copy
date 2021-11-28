package main 

import (
	"fmt"
)

func main()  {

	/* 
		La función copy copia el minimo de elementos en ambos arreglos
	*/
	slice := []int{1,2,3,4}
	arr_2 := make([]int, len(slice) cap(slice)*2)

	copy(arr_2, slice)

	fmt.Println(slice)
	fmt.Println(arr_2)


}