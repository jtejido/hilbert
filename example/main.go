package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"fmt"
	. "github.com/jtejido/hilbert"
	"math/big"
	"math/bits"
)

func main() {

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	fmt.Println("starting at bits = 5, dimension = 2")
	sm, _ := New(uint32(bits.Len(1028)), 2)

	fmt.Println("decode 1028")

	pts := make(plotter.XYs, 1028)

	for i := 0; i < 1028; i++ {
		arr := sm.Decode(big.NewInt(int64(i)))
		fmt.Printf("%v \n", arr)
		pts[i].X = float64(arr[0])
		pts[i].Y = float64(arr[1])
	}

	err = plotutil.AddLinePoints(p,
		"hilbert", pts,
	)
	if err != nil {
		panic(err)
	}

	// A bug in plotter shows a point on the lower right part (doesn't matter whether you just show 4-index or higher)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}
