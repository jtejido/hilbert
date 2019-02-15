package main

import (
	"fmt"
	. "github.com/jtejido/hilbert"
	"math/big"
)

func main() {
	fmt.Println("starting at bits = 5, dimension = 2")
	sm, _ := New(5, 2)

	fmt.Println("decode 1074")
	arr := sm.Decode(big.NewInt(1074))
	fmt.Printf("%v \n", arr)
	t := sm.Encode(arr[0], arr[1])
	fmt.Println("encode arr[0], arr[1]")
	fmt.Println(t)
	fmt.Println("decode back to 1074")
	arr2 := sm.Decode(t)
	fmt.Printf("%v \n", arr2)
}
