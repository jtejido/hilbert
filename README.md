![hilbert_3d](https://c1.staticflickr.com/3/2666/4245088030_b00b1351de.jpg)

# hilbert

[![Build Status](https://travis-ci.org/jtejido/hilbert.svg?branch=master)](https://travis-ci.org/jtejido/hilbert) 
[![codecov](https://codecov.io/gh/jtejido/hilbert/branch/master/graph/badge.svg)](https://codecov.io/gh/jtejido/hilbert)
[![MIT license](http://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT) 


Transform N-dimensional points to and from a 1-dimensional Hilbert fractal curve index.

While there has been a number of 2-d & 3-d implementation present online, it's somehow hard to find one that generalizes the notion of Hilbert curves to arbitrary dimensions of hypercube programmatically.

This is helpful (in my personal use-case) for N-dimensional tree variants (B-Tree/R-Tree).

The core algorithm is a port to Golang from the C code in a paper written by John Skilling published in the journal "Programming the Hilbert curve", (c) 2004 American Institute of Physics.

This package was meant to work on large set/s of integers(big.Int).

## Usage:
 
```golang
package main

import (
	"fmt"
	"github.com/jtejido/hilbert"
	"math/big"
)

func main() {
	fmt.Println("starting at bits = 5, dimension = 2")
	sm, err := hilbert.New(5, 2)

	if err != nil {
		fmt.Println(err)
	}

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
```

## Floating point data

 This library applies to non-negative integer data. To apply it to floating point numbers, you need to do the following:

 1. Decide how much resolution in bits you require for each coordinate. 
    The more bits of precision you use, the higher the cost of the transformation.

 2. Write methods to perform a two-way mapping from your coordinate system to the non-negative integers.
    This transform may require shifting and scaling each dimension a different amount in order to yield a desirable distance metric and origin. 
