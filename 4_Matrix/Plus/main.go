package main

import (
	"fmt"
	"errors"
)

func Plus(x, y [][]float32) ([][]float32, error) {

	if (len(x[0]) != len(y[0]) ) || (len(x) != len(y)) {
		return nil, errors.New("Can't do matrix Plus.")
	}

	out := make([][]float32, len(x))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(y[0]); j += 1 {
			if len(out[i]) < 1 {
				out[i] = make([]float32, len(y[0]))
			}
			out[i][j] += x[i][j] + y[i][j]
		}
	}
	return out, nil
}

func main() {

	x := [][]float32{
		[]float32{1.0, 2.0, 3.0},
		[]float32{4.0, 5.0, 6.0},
	}

	y := [][]float32{
		[]float32{0.5, 0.2, 0.7},
		[]float32{0.5, 0.8, 0.3},
	}

	out, err := Plus(x,y)

	if err!=nil {
		panic(err)
	}


	for _,x:=range out {

		fmt.Println(x)

	}


}