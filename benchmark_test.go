package vmath

//Benchmark Array based vectors and matrices against
// struct based
import (
	"testing"
)

//Struct Based
type SVector struct {
	X, Y, Z float32
}

type SMatrix struct {
	Col0, Col1, Col2 SVector
}

func SV3MakeFromElems(result *SVector, x, y, z float32) {
	result.X = x
	result.Y = y
	result.Z = z
}
func SM3MakeFromCols(result *SMatrix, col0, col1, col2 *SVector) {
	SV3Copy(&result.Col0, col0)
	SV3Copy(&result.Col1, col1)
	SV3Copy(&result.Col2, col2)
}
func SM3MulV3(result *SVector, mat *SMatrix, vec *SVector) {
	tmpX := ((mat.Col0.X * vec.X) + (mat.Col1.X * vec.Y)) + (mat.Col2.X * vec.Z)
	tmpY := ((mat.Col0.Y * vec.X) + (mat.Col1.Y * vec.Y)) + (mat.Col2.Y * vec.Z)
	tmpZ := ((mat.Col0.Z * vec.X) + (mat.Col1.Z * vec.Y)) + (mat.Col2.Z * vec.Z)
	SV3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func SM3Mul(result, mat0, mat1 *SMatrix) {
	var tmpResult SMatrix
	SM3MulV3(&tmpResult.Col0, mat0, &mat1.Col0)
	SM3MulV3(&tmpResult.Col1, mat0, &mat1.Col1)
	SM3MulV3(&tmpResult.Col2, mat0, &mat1.Col2)
	SM3Copy(result, &tmpResult)
}

func SM3Copy(result *SMatrix, mat *SMatrix) {
	SV3Copy(&result.Col0, &mat.Col0)
	SV3Copy(&result.Col1, &mat.Col1)
	SV3Copy(&result.Col2, &mat.Col2)
}

func SV3Copy(result *SVector, vec *SVector) {
	result.X = vec.X
	result.Y = vec.Y
	result.Z = vec.Z
}

//Array Based
type AVector [3]float32
type AMatrix [3 * 3]float32

func (result *AVector) SetElems(x, y, z float32) {
	result[0] = x
	result[1] = y
	result[2] = z
}

func AM3MakeFromCols(result *AMatrix, col0, col1, col2 *AVector) {
	result[0] = col0[0]
	result[1] = col0[1]
	result[2] = col0[2]
	result[3] = col1[0]
	result[4] = col1[1]
	result[5] = col1[2]
	result[6] = col2[2]
	result[7] = col2[2]
	result[8] = col2[2]
}

func AM3MulV3(result *AVector, mat *AMatrix, vec *AVector) {
	tmpX := ((mat[0] * vec[0]) + (mat[3] * vec[1])) + (mat[6] * vec[2])
	tmpY := ((mat[1] * vec[0]) + (mat[4] * vec[1])) + (mat[7] * vec[2])
	tmpZ := ((mat[2] * vec[0]) + (mat[5] * vec[1])) + (mat[8] * vec[2])
	result.SetElems(tmpX, tmpY, tmpZ)
}

func (result *AMatrix) AM3Mul(mat0, mat1 *AMatrix) {

	for c := 0; c < 3; c++ {
		result[(c*3)+0] = ((mat0[0] * mat1[(c*3)+0]) + (mat0[3] * mat1[(c*3)+1])) + (mat0[6] * mat1[(c*3)+2])
		result[(c*3)+1] = ((mat0[1] * mat1[(c*3)+0]) + (mat0[4] * mat1[(c*3)+1])) + (mat0[7] * mat1[(c*3)+2])
		result[(c*3)+2] = ((mat0[2] * mat1[(c*3)+0]) + (mat0[5] * mat1[(c*3)+1])) + (mat0[8] * mat1[(c*3)+2])
	}
}

func (result *AMatrix) AM3MulSelf(mat *AMatrix) {
	tmp := *result
	result.AM3Mul(&tmp, mat)
}

func M3Mul(mat0, mat1 *AMatrix) *AMatrix {
	return &AMatrix{
		((mat0[0] * mat1[(0*3)+0]) + (mat0[3] * mat1[(0*3)+1])) + (mat0[6] * mat1[(0*3)+2]),
		((mat0[1] * mat1[(0*3)+0]) + (mat0[4] * mat1[(0*3)+1])) + (mat0[7] * mat1[(0*3)+2]),
		((mat0[2] * mat1[(0*3)+0]) + (mat0[5] * mat1[(0*3)+1])) + (mat0[8] * mat1[(0*3)+2]),

		((mat0[0] * mat1[(1*3)+0]) + (mat0[3] * mat1[(1*3)+1])) + (mat0[6] * mat1[(1*3)+2]),
		((mat0[1] * mat1[(1*3)+0]) + (mat0[4] * mat1[(1*3)+1])) + (mat0[7] * mat1[(1*3)+2]),
		((mat0[2] * mat1[(1*3)+0]) + (mat0[5] * mat1[(1*3)+1])) + (mat0[8] * mat1[(1*3)+2]),

		((mat0[0] * mat1[(2*3)+0]) + (mat0[3] * mat1[(2*3)+1])) + (mat0[6] * mat1[(2*3)+2]),
		((mat0[1] * mat1[(2*3)+0]) + (mat0[4] * mat1[(2*3)+1])) + (mat0[7] * mat1[(2*3)+2]),
		((mat0[2] * mat1[(2*3)+0]) + (mat0[5] * mat1[(2*3)+1])) + (mat0[8] * mat1[(2*3)+2]),
	}
}

//Benchmarks

func BenchmarkStructM3MulV3(b *testing.B) {
	vec := new(SVector)
	mat := new(SMatrix)

	SV3MakeFromElems(vec, 3.59024, -11.3123, 342.2111)

	SM3MakeFromCols(mat, &SVector{23.41, 21.12, 0},
		&SVector{214.23, 213.9821, -32.02},
		&SVector{991.90, 75.123, -231.02})

	for i := 0; i < b.N; i++ {
		SM3MulV3(vec, mat, vec)
	}

}

func BenchmarkArrayM3MulV3(b *testing.B) {

	vec := &AVector{3.59024, -11.3123, 342.2111}

	mat := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	for i := 0; i < b.N; i++ {
		AM3MulV3(vec, mat, vec)
	}

}

func BenchmarkStructM3Mul(b *testing.B) {
	mat0 := new(SMatrix)
	mat1 := new(SMatrix)

	SM3MakeFromCols(mat0, &SVector{23.41, 21.12, 0},
		&SVector{214.23, 213.9821, -32.02},
		&SVector{991.90, 75.123, -231.02})

	SM3MakeFromCols(mat1, &SVector{23.41, 21.12, 0},
		&SVector{214.23, 213.9821, -32.02},
		&SVector{991.90, 75.123, -231.02})

	for i := 0; i < b.N; i++ {
		SM3Mul(mat0, mat0, mat1)
	}

}

func BenchmarkArrayM3Mul(b *testing.B) {

	mat0 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	mat1 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	result := &AMatrix{}

	for i := 0; i < b.N; i++ {
		result.AM3Mul(mat0, mat1)
	}

}

func BenchmarkArrayM3MulSelf(b *testing.B) {

	mat0 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	mat1 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	for i := 0; i < b.N; i++ {
		mat0.AM3MulSelf(mat1)
	}

}

func BenchmarkArrayM3MulReturn(b *testing.B) {

	mat0 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	mat1 := &AMatrix{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	for i := 0; i < b.N; i++ {
		_ = M3Mul(mat0, mat1)
	}

}

func BenchmarkVectorConversion(b *testing.B) {
	v := &Vector3{1, 2, 3}
	a := v.Array()
	_ = a
}
