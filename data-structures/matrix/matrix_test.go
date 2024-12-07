package matrix

import (
	"everything-go/data-structures/vector"
	"testing"
)

type vec = vector.Vector[int]

func TestNewMatrix(t *testing.T) {
	rows, cols := 3, 3
	mat := New[int](rows, cols)

	if len(mat) != rows {
		t.Errorf("expected %d rows, got %d", rows, len(mat))
	}

	for i := 0; i < rows; i++ {
		if len(mat[i]) != cols {
			t.Errorf("expected %d columns in row %d, got %d", cols, i, len(mat[i]))
		}
	}
}

func TestMatrixSize(t *testing.T) {
	rows, cols := 4, 5
	mat := New[int](rows, cols)

	if mat.SizeRows() != rows {
		t.Errorf("expected %d rows, got %d", rows, mat.SizeRows())
	}

	if mat.SizeCols() != cols {
		t.Errorf("expected %d columns, got %d", cols, mat.SizeCols())
	}
}

func TestMatrixClone(t *testing.T) {
	rows, cols := 2, 2
	mat := New[int](rows, cols)
	mat[0][0] = 1
	mat[1][1] = 2

	clone := mat.Clone()

	if clone.SizeRows() != rows || clone.SizeCols() != cols {
		t.Errorf("clone size mismatch")
	}

	if clone[0][0] != 1 || clone[1][1] != 2 {
		t.Errorf("clone content mismatch")
	}
}

func TestMatrixMultiply(t *testing.T) {
	a := New[int](2, 3)
	b := New[int](3, 2)

	a[0][0], a[0][1], a[0][2] = 1, 2, 3
	a[1][0], a[1][1], a[1][2] = 4, 5, 6

	b[0][0], b[0][1] = 7, 8
	b[1][0], b[1][1] = 9, 10
	b[2][0], b[2][1] = 11, 12

	expected := New[int](2, 2)
	expected[0][0], expected[0][1] = 58, 64
	expected[1][0], expected[1][1] = 139, 154

	result, err := Multiply(a, b)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < result.SizeRows(); i++ {
		for j := 0; j < result.SizeCols(); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("expected %d at position (%d, %d), got %d", expected[i][j], i, j, result[i][j])
			}
		}
	}
}

func TestMatrixPower(t *testing.T) {
	mat := New[int](2, 2)
	mat[0][0], mat[0][1] = 1, 2
	mat[1][0], mat[1][1] = 3, 4

	expected := New[int](2, 2)
	expected[0][0], expected[0][1] = 7, 10
	expected[1][0], expected[1][1] = 15, 22

	result, err := Power(mat, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < result.SizeRows(); i++ {
		for j := 0; j < result.SizeCols(); j++ {
			if result[i][j] != expected[i][j] {
				t.Errorf("expected %d at position (%d, %d), got %d", expected[i][j], i, j, result[i][j])
			}
		}
	}
}

func TestFibonacci(t *testing.T) {
	mat := Matrix[int]{
		vec{1, 1},
		vec{1, 0},
	}

	fib := Matrix[int]{
		vec{1},
		vec{0},
	}

	// test fibonacci up to 30
	res := []int{0, 1}
	for i := 2; i < 30; i++ {
		res = append(res, res[i-1]+res[i-2])
		poweredMat, err := Power(mat, i-1)
		if err != nil {
			t.Errorf("Unexpected err from Power call: %v", err)
		}
		matrixRes, err := Multiply(poweredMat, fib)
		if err != nil {
			t.Errorf("Unexpected err from Multiply call: %v", err)
		}
		if matrixRes[0][0] != res[i] {
			t.Errorf("expected %d, got %d", res[i], matrixRes[0][0])
		}
	}
}

func TestMatrixCopy(t *testing.T) {
	src := New[int](2, 2)
	src[0][0], src[0][1] = 1, 2
	src[1][0], src[1][1] = 3, 4

	dest := New[int](2, 2)
	err := Copy(dest, src)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for i := 0; i < src.SizeRows(); i++ {
		for j := 0; j < src.SizeCols(); j++ {
			if dest[i][j] != src[i][j] {
				t.Errorf("expected %d at position (%d, %d), got %d", src[i][j], i, j, dest[i][j])
			}
		}
	}
}

func TestIdentityMatrix(t *testing.T) {
	tc := []struct {
		size     int
		expected Matrix[int]
	}{
		{
			size: 2,
			expected: Matrix[int]{
				vec{1, 0},
				vec{0, 1},
			},
		},
		{
			size: 3,
			expected: Matrix[int]{
				vec{1, 0, 0},
				vec{0, 1, 0},
				vec{0, 0, 1},
			},
		},
	}

	for _, test := range tc {
		result := Identity[int](test.size)

		if result.SizeRows() != test.size || result.SizeCols() != test.size {
			t.Errorf("expected size %d, got rows=%d cols=%d",
				test.size, result.SizeRows(), result.SizeCols())
		}

		for i := 0; i < test.size; i++ {
			for j := 0; j < test.size; j++ {
				expected := 0
				if i == j {
					expected = 1
				}
				if result[i][j] != expected {
					t.Errorf("at position (%d,%d): expected %d, got %d",
						i, j, expected, result[i][j])
				}
			}
		}
	}
}

func TestMatrixFill(t *testing.T) {
	mat := New[int](3, 2)
	fillValue := 42

	mat.Fill(fillValue)

	for i := 0; i < mat.SizeRows(); i++ {
		for j := 0; j < mat.SizeCols(); j++ {
			if mat[i][j] != fillValue {
				t.Errorf("expected %d at position (%d, %d), got %d", fillValue, i, j, mat[i][j])
			}
		}
	}
}

func TestEmptyMatrix(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "empty matrix size",
			test: func(t *testing.T) {
				mat := New[int](0, 0)
				if mat.SizeRows() != 0 {
					t.Errorf("expected 0 rows, got %d", mat.SizeRows())
				}
				if mat.SizeCols() != 0 {
					t.Errorf("expected 0 columns, got %d", mat.SizeCols())
				}
			},
		},
		{
			name: "multiply with empty matrix",
			test: func(t *testing.T) {
				a := New[int](0, 0)
				b := New[int](2, 2)
				_, err := Multiply(a, b)
				if err == nil {
					t.Error("expected error when multiplying empty matrix, got nil")
				}
				if err.Error() != "cannot multiply empty matrix" {
					t.Errorf("expected 'cannot multiply empty matrix' error, got '%s'", err.Error())
				}
			},
		},
		{
			name: "copy to smaller destination",
			test: func(t *testing.T) {
				src := New[int](3, 3)
				dest := New[int](2, 2)
				err := Copy(dest, src)
				if err == nil {
					t.Error("expected error when copying to smaller destination, got nil")
				}
				if err.Error() != "destination matrix does not have enough size" {
					t.Errorf("expected 'destination matrix does not have enough size' error, got '%s'", err.Error())
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(t)
		})
	}
}

func TestMatrixPrint(t *testing.T) {
	mat := New[int](2, 2)
	mat[0][0], mat[0][1] = 1, 2
	mat[1][0], mat[1][1] = 3, 4

	// Since Print() writes to stdout, we can't easily test the output
	// This is more of a smoke test to ensure it doesn't panic
	mat.Print()
}
