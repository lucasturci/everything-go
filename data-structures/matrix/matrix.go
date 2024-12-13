package matrix

import (
	"errors"
	"fmt"

	"github.com/lucasturci/everything-go/data-structures/vector"

	"golang.org/x/exp/constraints"
)

type number interface {
	constraints.Integer | constraints.Float
}

type Matrix[T any] vector.Vector[vector.Vector[T]]

// constructors
func New[T any](rows, cols int) Matrix[T] {
	mat := vector.NewWithSize[vector.Vector[T]](rows)
	for i := 0; i < rows; i++ {
		mat[i] = vector.NewWithSize[T](cols)
	}
	return Matrix[T](mat)
}

func Identity[T number](n int) Matrix[T] {
	mat := New[T](n, n)
	for i := 0; i < n; i++ {
		mat[i][i] = T(1)
	}
	return mat
}

func (m Matrix[T]) Print() {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Printf("%v ", m[i][j])
		}
		fmt.Println()
	}
}

func (m Matrix[T]) SizeRows() int {
	return len(m)
}

func (m Matrix[T]) SizeCols() int {
	if m.SizeRows() == 0 {
		return 0
	}
	return len(m[0])
}

func (m Matrix[T]) Clone() Matrix[T] {
	ans := New[T](m.SizeRows(), m.SizeCols())
	Copy(ans, m)
	return ans
}

func (m Matrix[T]) Fill(val T) {
	for i := 0; i < m.SizeRows(); i++ {
		for j := 0; j < m.SizeCols(); j++ {
			m[i][j] = val
		}
	}
}

func Copy[T any](dest Matrix[T], src Matrix[T]) error {
	if len(dest) < len(src) {
		return errors.New("destination matrix does not have enough size")
	}
	// to copy a matrix, you have to copy all the rows
	for i := 0; i < len(dest); i++ {
		vector.Copy(dest[i], src[i])
	}
	return nil
}

// Algorithms

func multiplyImpl[T number](a Matrix[T], b Matrix[T], fast bool) (Matrix[T], error) {
	if a.SizeRows() == 0 || a.SizeCols() == 0 || b.SizeRows() == 0 || b.SizeCols() == 0 {
		return Matrix[T]{}, errors.New("cannot multiply empty matrix")
	}
	if a.SizeCols() != b.SizeRows() {
		return Matrix[T]{}, fmt.Errorf(
			"# cols of first matrix (%v) != # rows of second matrix (%v)", a.SizeCols(), b.SizeRows())
	}
	res := New[T](a.SizeRows(), b.SizeCols())

	done := make(chan bool)
	for i := 0; i < res.SizeRows(); i++ {
		for j := 0; j < res.SizeCols(); j++ {
			if fast {
				go func(i, j int) {
					for k := 0; k < a.SizeCols(); k++ {
						res[i][j] += a[i][k] * b[k][j]
					}
					done <- true
				}(i, j)
			} else {
				for k := 0; k < a.SizeCols(); k++ {
					res[i][j] += a[i][k] * b[k][j]
				}
			}
		}
	}

	if fast { // wait for all goroutines to be concluded
		for i := 0; i < res.SizeRows()*res.SizeCols(); i++ {
			<-done
		}
	}
	return res, nil
}

func Multiply[T number](a Matrix[T], b Matrix[T]) (Matrix[T], error) {
	return multiplyImpl(a, b, false /*fast mult*/)
}

// Fast Mult does multiplication using goroutines
func FastMult[T number](a Matrix[T], b Matrix[T]) (Matrix[T], error) {
	return multiplyImpl(a, b, true /*fast mult*/)
}

func powerImpl[T number](m Matrix[T], b int, fast bool) (Matrix[T], error) {
	if m.SizeRows() != m.SizeCols() {
		return Matrix[T]{}, errors.New("Matrix must be a square matrix to power")
	}

	mult := Multiply[T]
	if fast {
		mult = FastMult[T]
	}

	ans := Identity[T](m.SizeRows())
	for ; b > 0; b >>= 1 {
		var err error
		if b&1 == 1 {
			ans, err = mult(ans, m)
			if err != nil {
				return Matrix[T]{}, err
			}
		}
		m, err = mult(m, m)
		if err != nil {
			return Matrix[T]{}, err
		}
	}
	return ans, nil
}

func Power[T number](m Matrix[T], b int) (Matrix[T], error) {
	return powerImpl[T](m, b, false /*fast power*/)
}

func FastPower[T number](m Matrix[T], b int) (Matrix[T], error) {
	return powerImpl[T](m, b, true /*fast power*/)
}
