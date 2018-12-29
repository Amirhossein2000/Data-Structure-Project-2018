package main

import (
	"errors"
	"fmt"
)

func multiplication(x, y [][]float32) ([][]float32, error) {
	if len(x[0]) != len(y) {
		return nil, errors.New("Can't do matrix multiplication.")
	}

	out := make([][]float32, len(x))

	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(y[0]); j += 1 {
			for z := 0; z < len(y); z += 1 {

				if len(out[i]) < 1 {
					out[i] = make([]float32, len(y[0]))
				}

				out[i][j] += x[i][z] * y[z][j]
			}
		}
	}
	return out, nil
}

func main() {

	fmt.Println("====================")

	out, err := multiplication([][]float32{
		[]float32{1, 1, 1},
		[]float32{0, 1, 1},
		[]float32{0, 0, 1},
	}, [][]float32{
		[]float32{1, 1, 1},
		[]float32{0, 1, 1},
		[]float32{0, 0, 1},
	})

	errcheck(err)

	for _, x := range out {

		fmt.Println(x)

	}

	fmt.Println("====================")

	out, err = multiplication([][]float32{
		[]float32{1, 0, 0},
		[]float32{1, 1, 0},
		[]float32{1, 1, 1},
	}, [][]float32{
		[]float32{1, 0, 0},
		[]float32{1, 1, 0},
		[]float32{1, 1, 1},
	})

	errcheck(err)

	for _, x := range out {

		fmt.Println(x)

	}

	fmt.Println("====================")

	out, err = multiplication([][]float32{
		[]float32{1, 0, 0},
		[]float32{1, 1, 0},
		[]float32{1, 1, 1},
	}, [][]float32{
		[]float32{1, 1, 1},
		[]float32{0, 1, 1},
		[]float32{0, 0, 1},
	})

	errcheck(err)

	for _, x := range out {

		fmt.Println(x)

	}

	fmt.Println("====================")

	out, err = multiplication([][]float32{
		[]float32{1, 0, 0},
		[]float32{1, 0, 0},
		[]float32{1, 0, 0},
	}, [][]float32{
		[]float32{1, 1, 1},
		[]float32{0, 0, 0},
		[]float32{0, 0, 0},
	})

	errcheck(err)

	for _, x := range out {

		fmt.Println(x)

	}

}

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}
