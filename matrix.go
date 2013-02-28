//Copyright (C) 2006, 2007 Sony Computer Entertainment Inc.
//  All rights reserved.
// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

const (
	m3col0 = 0
	m3col1 = 3
	m3col2 = 6
)

const g_PI_OVER_2 = 1.570796327

func (result *Matrix3) MakeFromScalar(scalar float32) {
	result[m3col0+x] = scalar
	result[m3col0+y] = scalar
	result[m3col0+z] = scalar

	result[m3col1+x] = scalar
	result[m3col1+y] = scalar
	result[m3col1+z] = scalar

	result[m3col2+x] = scalar
	result[m3col2+y] = scalar
	result[m3col2+z] = scalar
}

func (result *Matrix3) MakeFromQ(unitQuat *Quaternion) {
	qx := unitQuat[x]
	qy := unitQuat[x]
	qz := unitQuat[x]
	qw := unitQuat[x]
	qx2 := qx + qx
	qy2 := qy + qy
	qz2 := qz + qz
	qxqx2 := qx * qx2
	qxqy2 := qx * qy2
	qxqz2 := qx * qz2
	qxqw2 := qw * qx2
	qyqy2 := qy * qy2
	qyqz2 := qy * qz2
	qyqw2 := qw * qy2
	qzqz2 := qz * qz2
	qzqw2 := qw * qz2

	result[m3col0+x] = ((1.0 - qyqy2) - qzqz2)
	result[m3col0+y] = (qxqy2 + qzqw2)
	result[m3col0+z] = (qxqz2 - qyqw2)

	result[m3col1+x] = (qxqy2 - qzqw2)
	result[m3col1+y] = ((1.0 - qxqx2) - qzqz2)
	result[m3col1+z] = (qyqz2 + qxqw2)

	result[m3col2+x] = (qxqz2 + qyqw2)
	result[m3col2+y] = (qyqz2 - qxqw2)
	result[m3col2+z] = ((1.0 - qxqx2) - qyqy2)
}

func (result *Matrix3) MakeFromCols(col0, col1, col2 *Vector3) {
	result.SetCol(0, col0)
	result.SetCol(1, col1)
	result.SetCol(2, col2)
}

func (m *Matrix3) SetCol(col int, vec *Vector3) {
	switch col {
	case 0:
		m[m3col0+x] = vec[x]
		m[m3col0+y] = vec[y]
		m[m3col0+z] = vec[z]
	case 1:
		m[m3col1+x] = vec[x]
		m[m3col1+y] = vec[y]
		m[m3col1+z] = vec[z]
	case 2:
		m[m3col2+x] = vec[x]
		m[m3col2+y] = vec[y]
		m[m3col2+z] = vec[z]
	}
}

func (m *Matrix3) SetRow(row int, vec *Vector3) {
	m[m3col0+row] = vec[x]
	m[m3col1+row] = vec[y]
	m[m3col2+row] = vec[z]
}

func (m *Matrix3) SetElem(col, row int, val float32) {
	m[col*3+row] = val
}

func (m *Matrix3) GetElem(col, row int) float32 {
	return m[col*3+row]
}

func (m *Matrix3) GetCol(result *Vector3, col int) {
	switch col {
	case 0:
		copy(result[:], m[m3col0:m3col1-1])
	case 1:
		copy(result[:], m[m3col1:m3col2-1])
	case 2:
		copy(result[:], m[m3col2:])
	}
}

func (mat *Matrix3) GetRow(result *Vector3, row int) {
	result[x] = mat[m3col0+row]
	result[y] = mat[m3col1+row]
	result[z] = mat[m3col2+row]
}

func (result *Matrix3) Transpose(mat *Matrix3) {
	result[m3col0+x] = mat[m3col0+x]
	result[m3col0+y] = mat[m3col1+x]
	result[m3col0+z] = mat[m3col2+x]

	result[m3col1+x] = mat[m3col0+y]
	result[m3col1+y] = mat[m3col1+y]
	result[m3col1+z] = mat[m3col2+y]

	result[m3col2+x] = mat[m3col0+z]
	result[m3col2+y] = mat[m3col1+z]
	result[m3col2+z] = mat[m3col2+z]

}

func (m *Matrix3) TransposeSelf() {
	tmp := *m
	m.Transpose(&tmp)
}

func (result *Matrix3) Inverse(mat *Matrix3) {
	var col0, col1, col2 Vector3
	var tmp0, tmp1, tmp2 Vector3

	mat.GetCol(&col0, 0)
	mat.GetCol(&col1, 1)
	mat.GetCol(&col2, 2)

	tmp0.Cross(&col1, &col2)
	tmp1.Cross(&col2, &col0)
	tmp2.Cross(&col0, &col1)

	detinv := 1.0 / col2.Dot(&tmp2)

	result[m3col0+x] = tmp0[x] * detinv
	result[m3col0+y] = tmp1[x] * detinv
	result[m3col0+z] = tmp2[x] * detinv

	result[m3col1+x] = tmp0[y] * detinv
	result[m3col1+y] = tmp1[y] * detinv
	result[m3col1+z] = tmp2[y] * detinv

	result[m3col1+x] = tmp0[z] * detinv
	result[m3col1+y] = tmp1[z] * detinv
	result[m3col1+z] = tmp2[z] * detinv

}

func (m *Matrix3) InvertSelf() {
	m.Inverse(m)
}

func (m *Matrix3) Determinant() float32 {
	var col0, col1, col2, tmp Vector3
	m.GetCol(&col0, 0)
	m.GetCol(&col1, 0)
	m.GetCol(&col2, 0)

	tmp.Cross(&col0, &col1)

	return col2.Dot(&tmp)
}

func (result *Matrix3) Add(mat0, mat1 *Matrix3) {
	result[m3col0+x] = mat0[m3col0+x] + mat1[m3col0+x]
	result[m3col0+y] = mat0[m3col0+y] + mat1[m3col0+y]
	result[m3col0+z] = mat0[m3col0+z] + mat1[m3col0+z]

	result[m3col1+x] = mat0[m3col1+x] + mat1[m3col1+x]
	result[m3col1+y] = mat0[m3col1+y] + mat1[m3col1+y]
	result[m3col1+z] = mat0[m3col1+z] + mat1[m3col1+z]

	result[m3col2+x] = mat0[m3col2+x] + mat1[m3col2+x]
	result[m3col2+y] = mat0[m3col2+y] + mat1[m3col2+y]
	result[m3col2+z] = mat0[m3col2+z] + mat1[m3col2+z]
}

func (result *Matrix3) AddSelf(mat *Matrix3) {
	result.Add(result, mat)
}

func (result *Matrix3) Sub(mat0, mat1 *Matrix3) {
	result[m3col0+x] = mat0[m3col0+x] - mat1[m3col0+x]
	result[m3col0+y] = mat0[m3col0+y] - mat1[m3col0+y]
	result[m3col0+z] = mat0[m3col0+z] - mat1[m3col0+z]

	result[m3col1+x] = mat0[m3col1+x] - mat1[m3col1+x]
	result[m3col1+y] = mat0[m3col1+y] - mat1[m3col1+y]
	result[m3col1+z] = mat0[m3col1+z] - mat1[m3col1+z]

	result[m3col2+x] = mat0[m3col2+x] - mat1[m3col2+x]
	result[m3col2+y] = mat0[m3col2+y] - mat1[m3col2+y]
	result[m3col2+z] = mat0[m3col2+z] - mat1[m3col2+z]
}

func (result *Matrix3) SubSelf(mat *Matrix3) {
	result.Sub(result, mat)
}

func (result *Matrix3) Neg(mat *Matrix3) {
	result[m3col0+x] = -mat[m3col0+x]
	result[m3col0+y] = -mat[m3col0+y]
	result[m3col0+z] = -mat[m3col0+z]

	result[m3col1+x] = -mat[m3col1+x]
	result[m3col1+y] = -mat[m3col1+y]
	result[m3col1+z] = -mat[m3col1+z]

	result[m3col2+x] = -mat[m3col2+x]
	result[m3col2+y] = -mat[m3col2+y]
	result[m3col2+z] = -mat[m3col2+z]
}

func (result *Matrix3) NegSelf() {
	result.Neg(result)
}

func (result *Matrix3) AbsPerElem(mat *Matrix3) {
	result[m3col0+x] = abs(mat[m3col0+x])
	result[m3col0+y] = abs(mat[m3col0+y])
	result[m3col0+z] = abs(mat[m3col0+z])

	result[m3col1+x] = abs(mat[m3col1+x])
	result[m3col1+y] = abs(mat[m3col1+y])
	result[m3col1+z] = abs(mat[m3col1+z])

	result[m3col2+x] = abs(mat[m3col2+x])
	result[m3col2+y] = abs(mat[m3col2+y])
	result[m3col2+z] = abs(mat[m3col2+z])
}

func (result *Matrix3) AbsPerElemSelf() {
	result.AbsPerElem(result)
}

func (result *Matrix3) ScalarMul(mat *Matrix3, scalar float32) {
	result[m3col0+x] = mat[m3col0+x] * scalar
	result[m3col0+y] = mat[m3col0+y] * scalar
	result[m3col0+z] = mat[m3col0+z] * scalar

	result[m3col1+x] = mat[m3col1+x] * scalar
	result[m3col1+y] = mat[m3col1+y] * scalar
	result[m3col1+z] = mat[m3col1+z] * scalar

	result[m3col2+x] = mat[m3col2+x] * scalar
	result[m3col2+y] = mat[m3col2+y] * scalar
	result[m3col2+z] = mat[m3col2+z] * scalar
}

func (result *Matrix3) ScalarMulSelf(scalar float32) {
	result.ScalarMul(result, scalar)
}

func (result *Vector3) MulM3(vec *Vector3, mat *Matrix3) {
	result[x] = ((mat[m3col0+x] * vec[x]) + (mat[m3col1+x] * vec[y])) + (mat[m3col2+x] * vec[z])
	result[y] = ((mat[m3col0+y] * vec[x]) + (mat[m3col1+y] * vec[y])) + (mat[m3col2+y] * vec[z])
	result[z] = ((mat[m3col0+z] * vec[x]) + (mat[m3col1+z] * vec[y])) + (mat[m3col2+z] * vec[z])
}

func (result *Vector3) MulM3Self(mat *Matrix3) {
	temp := *result
	result.MulM3(&temp, mat)
}

func (result *Matrix3) Mul(mat0, mat1 *Matrix3) {
	result[m3col0+x] = ((mat0[m3col0+x] * mat1[m3col0+x]) + (mat0[m3col1+x] * mat1[m3col0+y])) + (mat0[m3col2+x] * mat1[m3col0+z])
	result[m3col0+y] = ((mat0[m3col0+y] * mat1[m3col0+x]) + (mat0[m3col1+y] * mat1[m3col0+y])) + (mat0[m3col2+y] * mat1[m3col0+z])
	result[m3col0+z] = ((mat0[m3col0+z] * mat1[m3col0+x]) + (mat0[m3col1+z] * mat1[m3col0+y])) + (mat0[m3col2+z] * mat1[m3col0+z])

	result[m3col1+x] = ((mat0[m3col0+x] * mat1[m3col1+x]) + (mat0[m3col1+x] * mat1[m3col1+y])) + (mat0[m3col2+x] * mat1[m3col1+z])
	result[m3col1+y] = ((mat0[m3col0+y] * mat1[m3col1+x]) + (mat0[m3col1+y] * mat1[m3col1+y])) + (mat0[m3col2+y] * mat1[m3col1+z])
	result[m3col1+z] = ((mat0[m3col0+z] * mat1[m3col1+x]) + (mat0[m3col1+z] * mat1[m3col1+y])) + (mat0[m3col2+z] * mat1[m3col1+z])

	result[m3col2+x] = ((mat0[m3col0+x] * mat1[m3col2+x]) + (mat0[m3col1+x] * mat1[m3col2+y])) + (mat0[m3col2+x] * mat1[m3col2+z])
	result[m3col2+y] = ((mat0[m3col0+y] * mat1[m3col2+x]) + (mat0[m3col1+y] * mat1[m3col2+y])) + (mat0[m3col2+y] * mat1[m3col2+z])
	result[m3col2+z] = ((mat0[m3col0+z] * mat1[m3col2+x]) + (mat0[m3col1+z] * mat1[m3col2+y])) + (mat0[m3col2+z] * mat1[m3col2+z])
}

func (result *Matrix3) MulSelf(mat *Matrix3) {
	temp := *result
	result.Mul(&temp, mat)
}

func (result *Matrix3) MulPerElem(mat0, mat1 *Matrix3) {
	result[m3col0+x] = mat0[m3col0+x] * mat1[m3col0+x]
	result[m3col0+y] = mat0[m3col0+y] * mat1[m3col0+y]
	result[m3col0+z] = mat0[m3col0+z] * mat1[m3col0+z]

	result[m3col1+x] = mat0[m3col1+x] * mat1[m3col1+x]
	result[m3col1+y] = mat0[m3col1+y] * mat1[m3col1+y]
	result[m3col1+z] = mat0[m3col1+z] * mat1[m3col1+z]

	result[m3col2+x] = mat0[m3col2+x] * mat1[m3col2+x]
	result[m3col2+y] = mat0[m3col2+y] * mat1[m3col2+y]
	result[m3col2+z] = mat0[m3col2+z] * mat1[m3col2+z]
}

func (result *Matrix3) MulPerElemSelf(mat *Matrix3) {
	result.MulPerElem(result, mat)
}

func (result *Matrix3) MakeIdentity() {
	//x axis
	result[m3col0+x] = 1.0
	result[m3col0+y] = 0.0
	result[m3col0+z] = 0.0

	//y axis
	result[m3col1+x] = 0.0
	result[m3col1+y] = 1.0
	result[m3col1+z] = 0.0

	//z axis
	result[m3col2+x] = 0.0
	result[m3col2+y] = 0.0
	result[m3col2+z] = 1.0
}

func (result *Matrix3) MakeRotationX(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[m3col0+x] = 1.0
	result[m3col0+y] = 0.0
	result[m3col0+z] = 0.0

	result[m3col1+x] = 0.0
	result[m3col1+y] = c
	result[m3col1+z] = s

	result[m3col1+x] = 0.0
	result[m3col1+y] = -s
	result[m3col1+z] = c

}

func (result *Matrix3) MakeRotationY(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[m3col0+x] = c
	result[m3col0+y] = 0.0
	result[m3col0+z] = -s

	result[m3col2+x] = 0.0
	result[m3col2+y] = 1.0
	result[m3col2+z] = 0.0

	result[m3col2+x] = s
	result[m3col2+y] = 0.0
	result[m3col2+z] = c
}

func (result *Matrix3) MakeRotationZ(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[m3col0+x] = c
	result[m3col0+y] = s
	result[m3col0+z] = 0.0

	result[m3col1+x] = -s
	result[m3col1+y] = c
	result[m3col1+z] = 0.0

	result[m3col2+x] = 0.0
	result[m3col2+y] = 0.0
	result[m3col2+z] = 1.0
}

func (result *Matrix3) MakeRotationZYX(radiansXYZ *Vector3) {
	sX := sin(radiansXYZ[x])
	cX := cos(radiansXYZ[x])
	sY := sin(radiansXYZ[y])
	cY := cos(radiansXYZ[y])
	sZ := sin(radiansXYZ[z])
	cZ := cos(radiansXYZ[z])
	tmp0 := cZ * sY
	tmp1 := sZ * sY

	result[m3col0+x] = (cZ * cY)
	result[m3col0+y] = (sZ * cY)
	result[m3col0+z] = -sY

	result[m3col1+x] = ((tmp0 * sX) - (sZ * cX))
	result[m3col1+y] = ((tmp1 * sX) + (cZ * cX))
	result[m3col1+z] = (cY * sX)

	result[m3col2+x] = ((tmp0 * cX) + (sZ * sX))
	result[m3col2+y] = ((tmp1 * cX) - (cZ * sX))
	result[m3col2+z] = (cY * cX)
}

func (result *Matrix3) MakeRotationAxis(radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	X := unitVec[x]
	Y := unitVec[y]
	Z := unitVec[z]
	xy := X * Y
	yz := Y * Z
	zx := Z * X
	oneMinusC := 1.0 - c

	result[m3col0+x] = (((X * X) * oneMinusC) + c)
	result[m3col0+y] = ((xy * oneMinusC) + (Z * s))
	result[m3col0+z] = ((zx * oneMinusC) - (Y * s))

	result[m3col1+x] = ((xy * oneMinusC) - (Z * s))
	result[m3col1+y] = (((Y * Y) * oneMinusC) + c)
	result[m3col1+z] = ((yz * oneMinusC) + (X * s))

	result[m3col2+x] = ((zx * oneMinusC) + (Y * s))
	result[m3col2+y] = ((yz * oneMinusC) - (X * s))
	result[m3col2+z] = (((Z * Z) * oneMinusC) + c)
}

func (result *Matrix3) MakeRotationQ(unitQuat *Quaternion) {
	result.MakeFromQ(unitQuat)
}

func (result *Matrix3) MakeScale(scaleVec *Vector3) {
	result[m3col0+x] = scaleVec[x]
	result[m3col0+y] = 0.0
	result[m3col0+z] = 0.0

	result[m3col1+x] = 0.0
	result[m3col1+y] = scaleVec[y]
	result[m3col1+z] = 0.0

	result[m3col2+x] = 0.0
	result[m3col2+y] = 0.0
	result[m3col2+z] = scaleVec[z]
}

func (result *Matrix3) AppendScale(mat *Matrix3, scaleVec *Vector3) {
	result[m3col0+x] = mat[m3col0+x] * scaleVec[x]
	result[m3col0+y] = mat[m3col0+y] * scaleVec[x]
	result[m3col0+z] = mat[m3col0+z] * scaleVec[x]

	result[m3col1+x] = mat[m3col1+x] * scaleVec[y]
	result[m3col1+y] = mat[m3col1+y] * scaleVec[y]
	result[m3col1+z] = mat[m3col1+z] * scaleVec[y]

	result[m3col2+x] = mat[m3col2+x] * scaleVec[z]
	result[m3col2+y] = mat[m3col2+y] * scaleVec[z]
	result[m3col2+z] = mat[m3col2+z] * scaleVec[z]

}

func (result *Matrix3) AppendScaleSelf(scaleVec *Vector3) {
	result.AppendScale(result, scaleVec)
}

func (result *Matrix3) PrependScale(scaleVec *Vector3, mat *Matrix3) {
	result[m3col0+x] = mat[m3col0+x] * scaleVec[x]
	result[m3col0+y] = mat[m3col0+y] * scaleVec[y]
	result[m3col0+z] = mat[m3col0+z] * scaleVec[z]

	result[m3col1+x] = mat[m3col1+x] * scaleVec[x]
	result[m3col1+y] = mat[m3col1+y] * scaleVec[y]
	result[m3col1+z] = mat[m3col1+z] * scaleVec[z]

	result[m3col2+x] = mat[m3col2+x] * scaleVec[x]
	result[m3col2+y] = mat[m3col2+y] * scaleVec[y]
	result[m3col2+z] = mat[m3col2+z] * scaleVec[z]
}

func (result *Matrix3) PrependScaleSelf(scaleVec *Vector3) {
	result.PrependScale(scaleVec, result)

}

func (result *Matrix3) Select(mat0, mat1 *Matrix3, select1 int) {
	if select1 != 0 {
		result[m3col0+x] = mat1[m3col0+x]
		result[m3col0+y] = mat1[m3col0+y]
		result[m3col0+z] = mat1[m3col0+z]

		result[m3col1+x] = mat1[m3col1+x]
		result[m3col1+y] = mat1[m3col1+y]
		result[m3col1+z] = mat1[m3col1+z]

		result[m3col2+x] = mat1[m3col2+x]
		result[m3col2+y] = mat1[m3col2+y]
		result[m3col2+z] = mat1[m3col2+z]

	} else {
		result[m3col0+x] = mat0[m3col0+x]
		result[m3col0+y] = mat0[m3col0+y]
		result[m3col0+z] = mat0[m3col0+z]

		result[m3col1+x] = mat0[m3col1+x]
		result[m3col1+y] = mat0[m3col1+y]
		result[m3col1+z] = mat0[m3col1+z]

		result[m3col2+x] = mat0[m3col2+x]
		result[m3col2+y] = mat0[m3col2+y]
		result[m3col2+z] = mat0[m3col2+z]

	}

}

func (result *Matrix3) SelectSelf(mat *Matrix3, select1 int) {
	result.Select(result, mat, select1)
}

/*
//Matrix 4
const (
	m4col0 = 0
	m4col1 = 4
	m4col2 = 8
	m4Col3 = 12
)

func M4MakeFromScalar(result *Matrix4, scalar float32) {
	V4MakeFromScalar(&result.Col0, scalar)
	V4MakeFromScalar(&result.Col1, scalar)
	V4MakeFromScalar(&result.Col2, scalar)
	V4MakeFromScalar(&result.Col3, scalar)
}

func M4MakeFromT3(result *Matrix4, mat *Transform3) {
	V4MakeFromV3Scalar(&result.Col0, &mat.Col0, 0.0)
	V4MakeFromV3Scalar(&result.Col1, &mat.Col1, 0.0)
	V4MakeFromV3Scalar(&result.Col2, &mat.Col2, 0.0)
	V4MakeFromV3Scalar(&result.Col3, &mat.Col3, 1.0)
}

func M4MakeFromCols(result *Matrix4, col0, col1, col2, col3 *Vector4) {
	V4Copy(&result.Col0, col0)
	V4Copy(&result.Col1, col1)
	V4Copy(&result.Col2, col2)
	V4Copy(&result.Col3, col3)
}

func M4MakeFromM3V3(result *Matrix4, mat *Matrix3, translateVec *Vector3) {
	V4MakeFromV3Scalar(&result.Col0, &mat.Col0, 0.0)
	V4MakeFromV3Scalar(&result.Col1, &mat.Col1, 0.0)
	V4MakeFromV3Scalar(&result.Col2, &mat.Col2, 0.0)
	V4MakeFromV3Scalar(&result.Col3, translateVec, 1.0)
}

func M4MakeFromQV3(result *Matrix4, unitQuat *Quaternion, translateVec *Vector3) {
	var mat *Matrix3
	M3MakeFromQ(mat, unitQuat)
	V4MakeFromV3Scalar(&result.Col0, &mat.Col0, 0.0)
	V4MakeFromV3Scalar(&result.Col1, &mat.Col1, 0.0)
	V4MakeFromV3Scalar(&result.Col2, &mat.Col2, 0.0)
	V4MakeFromV3Scalar(&result.Col3, translateVec, 1.0)
}

func (m *Matrix4) SetCol(col int, vec *Vector4) {
	switch col {
	case 0:
		V4Copy(&m.Col0, vec)
	case 1:
		V4Copy(&m.Col1, vec)
	case 2:
		V4Copy(&m.Col2, vec)
	case 3:
		V4Copy(&m.Col3, vec)
	}
}

func (m *Matrix4) SetRow(row int, vec *Vector4) {
	m.Col0.SetElem(row, vec.X)
	m.Col1.SetElem(row, vec.Y)
	m.Col2.SetElem(row, vec.Z)
	m.Col3.SetElem(row, vec.W)
}

func (m *Matrix4) SetElem(col, row int, val float32) {
	var tmpV3_0 Vector4
	M4GetCol(&tmpV3_0, m, col)
	tmpV3_0.SetElem(row, val)
	m.SetCol(col, &tmpV3_0)
}

func (m *Matrix4) GetElem(col, row int) float32 {
	var tmpV4_0 Vector4
	M4GetCol(&tmpV4_0, m, col)
	return tmpV4_0.GetElem(row)
}

func M4GetCol(result *Vector4, mat *Matrix4, col int) {
	switch col {
	case 0:
		V4Copy(result, &mat.Col0)
	case 1:
		V4Copy(result, &mat.Col1)
	case 2:
		V4Copy(result, &mat.Col2)
	case 3:
		V4Copy(result, &mat.Col3)
	}
}

func M4GetRow(result *Vector4, mat *Matrix4, row int) {
	V4MakeFromElems(result, mat.Col0.GetElem(row), mat.Col1.GetElem(row), mat.Col2.GetElem(row), mat.Col3.GetElem(row))
}

func M4Transpose(result, mat *Matrix4) {
	var tmpResult Matrix4
	V4MakeFromElems(&tmpResult.Col0, mat.Col0.X, mat.Col1.X, mat.Col2.X, mat.Col3.X)
	V4MakeFromElems(&tmpResult.Col1, mat.Col0.Y, mat.Col1.Y, mat.Col2.Y, mat.Col3.Y)
	V4MakeFromElems(&tmpResult.Col2, mat.Col0.Z, mat.Col1.Z, mat.Col2.Z, mat.Col3.Z)
	V4MakeFromElems(&tmpResult.Col3, mat.Col0.W, mat.Col1.W, mat.Col2.W, mat.Col3.W)
	M4Copy(result, &tmpResult)
}

func (m *Matrix4) Transpose() {
	M4Transpose(m, m)
}

func M4Inverse(result, mat *Matrix4) {
	var res0, res1, res2, res3 Vector4
	mA := mat.Col0.X
	mB := mat.Col0.Y
	mC := mat.Col0.Z
	mD := mat.Col0.W
	mE := mat.Col1.X
	mF := mat.Col1.Y
	mG := mat.Col1.Z
	mH := mat.Col1.W
	mI := mat.Col2.X
	mJ := mat.Col2.Y
	mK := mat.Col2.Z
	mL := mat.Col2.W
	mM := mat.Col3.X
	mN := mat.Col3.Y
	mO := mat.Col3.Z
	mP := mat.Col3.W
	tmp0 := ((mK * mD) - (mC * mL))
	tmp1 := ((mO * mH) - (mG * mP))
	tmp2 := ((mB * mK) - (mJ * mC))
	tmp3 := ((mF * mO) - (mN * mG))
	tmp4 := ((mJ * mD) - (mB * mL))
	tmp5 := ((mN * mH) - (mF * mP))
	res0.X = (((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5))
	res0.Y = (((mN * tmp0) - (mP * tmp2)) - (mO * tmp4))
	res0.Z = (((mD * tmp3) + (mC * tmp5)) - (mB * tmp1))
	res0.W = (((mH * tmp2) + (mG * tmp4)) - (mF * tmp0))
	detInv := (1.0 / ((((mA * res0.X) + (mE * res0.Y)) + (mI * res0.Z)) + (mM * res0.W)))
	res1.X = (mI * tmp1)
	res1.Y = (mM * tmp0)
	res1.Z = (mA * tmp1)
	res1.W = (mE * tmp0)
	res3.X = (mI * tmp3)
	res3.Y = (mM * tmp2)
	res3.Z = (mA * tmp3)
	res3.W = (mE * tmp2)
	res2.X = (mI * tmp5)
	res2.Y = (mM * tmp4)
	res2.Z = (mA * tmp5)
	res2.W = (mE * tmp4)
	tmp0 = ((mI * mB) - (mA * mJ))
	tmp1 = ((mM * mF) - (mE * mN))
	tmp2 = ((mI * mD) - (mA * mL))
	tmp3 = ((mM * mH) - (mE * mP))
	tmp4 = ((mI * mC) - (mA * mK))
	tmp5 = ((mM * mG) - (mE * mO))
	res2.X = (((mL * tmp1) - (mJ * tmp3)) + res2.X)
	res2.Y = (((mP * tmp0) - (mN * tmp2)) + res2.Y)
	res2.Z = (((mB * tmp3) - (mD * tmp1)) - res2.Z)
	res2.W = (((mF * tmp2) - (mH * tmp0)) - res2.W)
	res3.X = (((mJ * tmp5) - (mK * tmp1)) + res3.X)
	res3.Y = (((mN * tmp4) - (mO * tmp0)) + res3.Y)
	res3.Z = (((mC * tmp1) - (mB * tmp5)) - res3.Z)
	res3.W = (((mG * tmp0) - (mF * tmp4)) - res3.W)
	res1.X = (((mK * tmp3) - (mL * tmp5)) - res1.X)
	res1.Y = (((mO * tmp2) - (mP * tmp4)) - res1.Y)
	res1.Z = (((mD * tmp5) - (mC * tmp3)) + res1.Z)
	res1.W = (((mH * tmp4) - (mG * tmp2)) + res1.W)
	V4ScalarMul(&result.Col0, &res0, detInv)
	V4ScalarMul(&result.Col1, &res1, detInv)
	V4ScalarMul(&result.Col2, &res2, detInv)
	V4ScalarMul(&result.Col3, &res3, detInv)
}

func (m *Matrix4) Invert() {
	M4Inverse(m, m)
}

func M4AffineInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.Col0)
	V4GetXYZ(&tmpV3_1, &mat.Col1)
	V4GetXYZ(&tmpV3_2, &mat.Col2)
	V4GetXYZ(&tmpV3_3, &mat.Col3)
	affineMat.SetCol(0, &tmpV3_0)
	affineMat.SetCol(1, &tmpV3_1)
	affineMat.SetCol(2, &tmpV3_2)
	affineMat.SetCol(3, &tmpV3_3)
	T3Inverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func (m *Matrix4) AffineInvert() {
	M4AffineInverse(m, m)
}

func M4OrthoInverse(result, mat *Matrix4) {
	var affineMat, tmpT3_0 Transform3
	var tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3 Vector3
	V4GetXYZ(&tmpV3_0, &mat.Col0)
	V4GetXYZ(&tmpV3_1, &mat.Col1)
	V4GetXYZ(&tmpV3_2, &mat.Col2)
	V4GetXYZ(&tmpV3_3, &mat.Col3)
	affineMat.SetCol(0, &tmpV3_0)
	affineMat.SetCol(1, &tmpV3_1)
	affineMat.SetCol(2, &tmpV3_2)
	affineMat.SetCol(3, &tmpV3_3)
	T3OrthoInverse(&tmpT3_0, &affineMat)
	M4MakeFromT3(result, &tmpT3_0)
}

func (m *Matrix4) OrthoInvert() {
	M4OrthoInverse(m, m)
}

func (m *Matrix4) Determinant() float32 {
	mA := m.Col0.X
	mB := m.Col0.Y
	mC := m.Col0.Z
	mD := m.Col0.W
	mE := m.Col1.X
	mF := m.Col1.Y
	mG := m.Col1.Z
	mH := m.Col1.W
	mI := m.Col2.X
	mJ := m.Col2.Y
	mK := m.Col2.Z
	mL := m.Col2.W
	mM := m.Col3.X
	mN := m.Col3.Y
	mO := m.Col3.Z
	mP := m.Col3.W
	tmp0 := ((mK * mD) - (mC * mL))
	tmp1 := ((mO * mH) - (mG * mP))
	tmp2 := ((mB * mK) - (mJ * mC))
	tmp3 := ((mF * mO) - (mN * mG))
	tmp4 := ((mJ * mD) - (mB * mL))
	tmp5 := ((mN * mH) - (mF * mP))
	dx := (((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5))
	dy := (((mN * tmp0) - (mP * tmp2)) - (mO * tmp4))
	dz := (((mD * tmp3) + (mC * tmp5)) - (mB * tmp1))
	dw := (((mH * tmp2) + (mG * tmp4)) - (mF * tmp0))
	return ((((mA * dx) + (mE * dy)) + (mI * dz)) + (mM * dw))
}

func M4Add(result, mat0, mat1 *Matrix4) {
	V4Add(&result.Col0, &mat0.Col0, &mat1.Col0)
	V4Add(&result.Col1, &mat0.Col1, &mat1.Col1)
	V4Add(&result.Col2, &mat0.Col2, &mat1.Col2)
	V4Add(&result.Col3, &mat0.Col3, &mat1.Col3)
}

func M4Sub(result, mat0, mat1 *Matrix4) {
	V4Sub(&result.Col0, &mat0.Col0, &mat1.Col0)
	V4Sub(&result.Col1, &mat0.Col1, &mat1.Col1)
	V4Sub(&result.Col2, &mat0.Col2, &mat1.Col2)
	V4Sub(&result.Col3, &mat0.Col3, &mat1.Col3)
}

func M4Neg(result, mat *Matrix4) {
	V4Neg(&result.Col0, &mat.Col0)
	V4Neg(&result.Col1, &mat.Col1)
	V4Neg(&result.Col2, &mat.Col2)
	V4Neg(&result.Col3, &mat.Col3)
}

func (m *Matrix4) Neg() {
	M4Neg(m, m)
}

func M4AbsPerElem(result, mat *Matrix4) {
	V4AbsPerElem(&result.Col0, &mat.Col0)
	V4AbsPerElem(&result.Col1, &mat.Col1)
	V4AbsPerElem(&result.Col2, &mat.Col2)
	V4AbsPerElem(&result.Col3, &mat.Col3)
}

func M4ScalarMul(result, mat *Matrix4, scalar float32) {
	V4ScalarMul(&result.Col0, &mat.Col0, scalar)
	V4ScalarMul(&result.Col1, &mat.Col1, scalar)
	V4ScalarMul(&result.Col2, &mat.Col2, scalar)
	V4ScalarMul(&result.Col3, &mat.Col3, scalar)
}

func M4MulV4(result *Vector4, mat *Matrix4, vec *Vector4) {
	tmpX := (((mat.Col0.X * vec.X) + (mat.Col1.X * vec.Y)) + (mat.Col2.X * vec.Z)) + (mat.Col3.X * vec.W)
	tmpY := (((mat.Col0.Y * vec.X) + (mat.Col1.Y * vec.Y)) + (mat.Col2.Y * vec.Z)) + (mat.Col3.Y * vec.W)
	tmpZ := (((mat.Col0.Z * vec.X) + (mat.Col1.Z * vec.Y)) + (mat.Col2.Z * vec.Z)) + (mat.Col3.Z * vec.W)
	tmpW := (((mat.Col0.W * vec.X) + (mat.Col1.W * vec.Y)) + (mat.Col2.W * vec.Z)) + (mat.Col3.W * vec.W)
	V4MakeFromElems(result, tmpX, tmpY, tmpZ, tmpW)
}

func M4MulV3(result *Vector4, mat *Matrix4, vec *Vector3) {
	result.X = ((mat.Col0.X * vec.X) + (mat.Col1.X * vec.Y)) + (mat.Col2.X * vec.Z)
	result.Y = ((mat.Col0.Y * vec.X) + (mat.Col1.Y * vec.Y)) + (mat.Col2.Y * vec.Z)
	result.Z = ((mat.Col0.Z * vec.X) + (mat.Col1.Z * vec.Y)) + (mat.Col2.Z * vec.Z)
	result.W = ((mat.Col0.W * vec.X) + (mat.Col1.W * vec.Y)) + (mat.Col2.W * vec.Z)
}

func M4MulP3(result *Vector4, mat *Matrix4, pnt *Point3) {
	result.X = (((mat.Col0.X * pnt.X) + (mat.Col1.X * pnt.Y)) + (mat.Col2.X * pnt.Z)) + mat.Col3.X
	result.Y = (((mat.Col0.Y * pnt.X) + (mat.Col1.Y * pnt.Y)) + (mat.Col2.Y * pnt.Z)) + mat.Col3.Y
	result.Z = (((mat.Col0.Z * pnt.X) + (mat.Col1.Z * pnt.Y)) + (mat.Col2.Z * pnt.Z)) + mat.Col3.Z
	result.W = (((mat.Col0.W * pnt.X) + (mat.Col1.W * pnt.Y)) + (mat.Col2.W * pnt.Z)) + mat.Col3.W
}

func M4Mul(result, mat0, mat1 *Matrix4) {
	var tmpResult Matrix4
	M4MulV4(&tmpResult.Col0, mat0, &mat1.Col0)
	M4MulV4(&tmpResult.Col1, mat0, &mat1.Col1)
	M4MulV4(&tmpResult.Col2, mat0, &mat1.Col2)
	M4MulV4(&tmpResult.Col3, mat0, &mat1.Col3)
	M4Copy(result, &tmpResult)
}

func M4MulT3(result, mat *Matrix4, tfrm1 *Transform3) {
	var tmpResult Matrix4
	var tmpP3_0 Point3
	M4MulV3(&tmpResult.Col0, mat, &tfrm1.Col0)
	M4MulV3(&tmpResult.Col1, mat, &tfrm1.Col1)
	M4MulV3(&tmpResult.Col2, mat, &tfrm1.Col2)
	P3MakeFromV3(&tmpP3_0, &tfrm1.Col3)
	M4MulP3(&tmpResult.Col3, mat, &tmpP3_0)
	M4Copy(result, &tmpResult)
}

func M4MulPerElem(result, mat0, mat1 *Matrix4) {
	V4MulPerElem(&result.Col0, &mat0.Col0, &mat1.Col0)
	V4MulPerElem(&result.Col1, &mat0.Col1, &mat1.Col1)
	V4MulPerElem(&result.Col2, &mat0.Col2, &mat1.Col2)
	V4MulPerElem(&result.Col3, &mat0.Col3, &mat1.Col3)
}

func M4MakeIdentity(result *Matrix4) {
	V4MakeXAxis(&result.Col0)
	V4MakeYAxis(&result.Col1)
	V4MakeZAxis(&result.Col2)
	V4MakeWAxis(&result.Col3)
}

func (m *Matrix4) SetUpper3x3(mat3 *Matrix3) {
	m.Col0.SetXYZ(&mat3.Col0)
	m.Col0.SetXYZ(&mat3.Col1)
	m.Col0.SetXYZ(&mat3.Col2)
}

func (m *Matrix4) Upper3x3(result *Matrix3) {
	V4GetXYZ(&result.Col0, &m.Col0)
	V4GetXYZ(&result.Col1, &m.Col1)
	V4GetXYZ(&result.Col2, &m.Col2)
}
func M4GetUpper3x3(result *Matrix3, mat *Matrix4) {
	V4GetXYZ(&result.Col0, &mat.Col0)
	V4GetXYZ(&result.Col1, &mat.Col1)
	V4GetXYZ(&result.Col2, &mat.Col2)
}

func (m *Matrix4) SetTranslation(translateVec *Vector3) {
	m.Col3.SetXYZ(translateVec)
}

func M4GetTranslation(result *Vector3, mat *Matrix4) {
	V4GetXYZ(result, &mat.Col3)
}

func (m *Matrix4) Translation(result *Vector3) {
	M4GetTranslation(result, m)
}

func M4MakeRotationX(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeXAxis(&result.Col0)
	V4MakeFromElems(&result.Col1, 0.0, c, s, 0.0)
	V4MakeFromElems(&result.Col2, 0.0, -s, c, 0.0)
	V4MakeWAxis(&result.Col3)
}

func M4MakeRotationY(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeFromElems(&result.Col0, c, 0.0, -s, 0.0)
	V4MakeYAxis(&result.Col1)
	V4MakeFromElems(&result.Col2, s, 0.0, c, 0.0)
	V4MakeWAxis(&result.Col3)
}

func M4MakeRotationZ(result *Matrix4, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V4MakeFromElems(&result.Col0, c, s, 0.0, 0.0)
	V4MakeFromElems(&result.Col1, -s, c, 0.0, 0.0)
	V4MakeZAxis(&result.Col2)
	V4MakeWAxis(&result.Col3)
}

func M4MakeRotationZYX(result *Matrix4, radiansXYZ *Vector3) {
	sX := sin(radiansXYZ.X)
	cX := cos(radiansXYZ.X)
	sY := sin(radiansXYZ.Y)
	cY := cos(radiansXYZ.Y)
	sZ := sin(radiansXYZ.Z)
	cZ := cos(radiansXYZ.Z)
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)
	V4MakeFromElems(&result.Col0, (cZ * cY), (sZ * cY), -sY, 0.0)
	V4MakeFromElems(&result.Col1, ((tmp0 * sX) - (sZ * cX)), ((tmp1 * sX) + (cZ * cX)), (cY * sX), 0.0)
	V4MakeFromElems(&result.Col2, ((tmp0 * cX) + (sZ * sX)), ((tmp1 * cX) - (cZ * sX)), (cY * cX), 0.0)
	V4MakeWAxis(&result.Col3)
}

func M4MakeRotationAxis(result *Matrix4, radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	x := unitVec.X
	y := unitVec.Y
	z := unitVec.Z
	xy := x * y
	yz := y * z
	zx := z * x
	oneMinusC := 1.0 - c
	V4MakeFromElems(&result.Col0, (((x * x) * oneMinusC) + c), ((xy * oneMinusC) + (z * s)), ((zx * oneMinusC) - (y * s)), 0.0)
	V4MakeFromElems(&result.Col1, ((xy * oneMinusC) - (z * s)), (((y * y) * oneMinusC) + c), ((yz * oneMinusC) + (x * s)), 0.0)
	V4MakeFromElems(&result.Col2, ((zx * oneMinusC) + (y * s)), ((yz * oneMinusC) - (x * s)), (((z * z) * oneMinusC) + c), 0.0)
	V4MakeWAxis(&result.Col3)
}

func M4MakeRotationQ(result *Matrix4, unitQuat *Quaternion) {
	var tmpT3_0 Transform3
	T3MakeRotationQ(&tmpT3_0, unitQuat)
	M4MakeFromT3(result, &tmpT3_0)
}

func M4MakeScale(result *Matrix4, scaleVec *Vector3) {
	V4MakeFromElems(&result.Col0, scaleVec.X, 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.Col1, 0.0, scaleVec.Y, 0.0, 0.0)
	V4MakeFromElems(&result.Col2, 0.0, 0.0, scaleVec.Z, 0.0)
	V4MakeWAxis(&result.Col3)
}

func M4AppendScale(result, mat *Matrix4, scaleVec *Vector3) {
	V4ScalarMul(&result.Col0, &mat.Col0, scaleVec.X)
	V4ScalarMul(&result.Col1, &mat.Col1, scaleVec.Y)
	V4ScalarMul(&result.Col2, &mat.Col2, scaleVec.Z)
	V4Copy(&result.Col3, &mat.Col3)
}

func M4PrependScale(result *Matrix4, scaleVec *Vector3, mat *Matrix4) {
	var scale4 Vector4
	V4MakeFromV3Scalar(&scale4, scaleVec, 1.0)
	V4MulPerElem(&result.Col0, &mat.Col0, &scale4)
	V4MulPerElem(&result.Col1, &mat.Col1, &scale4)
	V4MulPerElem(&result.Col2, &mat.Col2, &scale4)
	V4MulPerElem(&result.Col3, &mat.Col3, &scale4)
}

func M4MakeTranslation(result *Matrix4, translateVec *Vector3) {
	V4MakeXAxis(&result.Col0)
	V4MakeYAxis(&result.Col1)
	V4MakeZAxis(&result.Col2)
	V4MakeFromV3Scalar(&result.Col3, translateVec, 1.0)
}

func M4MakeLookAt(result *Matrix4, eyePos, lookAtPos *Point3, upVec *Vector3) {
	var m4EyeFrame Matrix4
	var v3X, v3Y, v3Z, tmpV3_0, tmpV3_1 Vector3
	var tmpV4_0, tmpV4_1, tmpV4_2, tmpV4_3 Vector4
	V3Normalize(&v3Y, upVec)
	P3Sub(&tmpV3_0, eyePos, lookAtPos)
	V3Normalize(&v3Z, &tmpV3_0)
	V3Cross(&tmpV3_1, &v3Y, &v3Z)
	V3Normalize(&v3X, &tmpV3_1)
	V3Cross(&v3Y, &v3Z, &v3X)
	V4MakeFromV3(&tmpV4_0, &v3X)
	V4MakeFromV3(&tmpV4_1, &v3Y)
	V4MakeFromV3(&tmpV4_2, &v3Z)
	V4MakeFromP3(&tmpV4_3, eyePos)
	M4MakeFromCols(&m4EyeFrame, &tmpV4_0, &tmpV4_1, &tmpV4_2, &tmpV4_3)
	M4OrthoInverse(result, &m4EyeFrame)
}

func M4MakePerspective(result *Matrix4, fovyRadians, aspect, zNear, zFar float32) {
	f := tan(g_PI_OVER_2 - (0.5 * fovyRadians))
	rangeInv := 1.0 / (zNear - zFar)
	V4MakeFromElems(&result.Col0, (f / aspect), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.Col1, 0.0, f, 0.0, 0.0)
	V4MakeFromElems(&result.Col2, 0.0, 0.0, ((zNear + zFar) * rangeInv), -1.0)
	V4MakeFromElems(&result.Col3, 0.0, 0.0, (((zNear * zFar) * rangeInv) * 2.0), 0.0)
}

func M4MakeFrustum(result *Matrix4, left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))
	n2 := (zNear + zNear)
	V4MakeFromElems(&result.Col0, (n2 * inv_rl), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.Col1, 0.0, (n2 * inv_tb), 0.0, 0.0)
	V4MakeFromElems(&result.Col2, (sum_rl * inv_rl), (sum_tb * inv_tb), (sum_nf * inv_nf), -1.0)
	V4MakeFromElems(&result.Col3, 0.0, 0.0, ((n2 * inv_nf) * zFar), 0.0)
}

func M4MakeOrthographic(result *Matrix4, left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))
	V4MakeFromElems(&result.Col0, (inv_rl + inv_rl), 0.0, 0.0, 0.0)
	V4MakeFromElems(&result.Col1, 0.0, (inv_tb + inv_tb), 0.0, 0.0)
	V4MakeFromElems(&result.Col2, 0.0, 0.0, (inv_nf + inv_nf), 0.0)
	V4MakeFromElems(&result.Col3, (-sum_rl * inv_rl), (-sum_tb * inv_tb), (sum_nf * inv_nf), 1.0)
}

func M4Select(result, mat0, mat1 *Matrix4, select1 int) {
	V4Select(&result.Col0, &mat0.Col0, &mat1.Col0, select1)
	V4Select(&result.Col1, &mat0.Col1, &mat1.Col1, select1)
	V4Select(&result.Col2, &mat0.Col2, &mat1.Col2, select1)
	V4Select(&result.Col3, &mat0.Col3, &mat1.Col3, select1)
}

func (m *Matrix4) String() string {
	var tmp Matrix4
	M4Transpose(&tmp, m)
	return tmp.Col0.String() + tmp.Col1.String() + tmp.Col2.String() + tmp.Col3.String()
}


func T3Copy(result, tfrm *Transform3) {
	V3Copy(&result.Col0, &tfrm.Col0)
	V3Copy(&result.Col1, &tfrm.Col1)
	V3Copy(&result.Col2, &tfrm.Col2)
	V3Copy(&result.Col3, &tfrm.Col3)
}

func T3MakeFromScalar(result *Transform3, scalar float32) {
	V3MakeFromScalar(&result.Col0, scalar)
	V3MakeFromScalar(&result.Col1, scalar)
	V3MakeFromScalar(&result.Col2, scalar)
	V3MakeFromScalar(&result.Col3, scalar)
}

func T3MakeFromCols(result *Transform3, col0, col1, col2, col3 *Vector3) {
	V3Copy(&result.Col0, col0)
	V3Copy(&result.Col1, col1)
	V3Copy(&result.Col2, col2)
	V3Copy(&result.Col3, col3)
}

func T3MakeFromM3V3(result *Transform3, tfrm *Matrix3, translateVec *Vector3) {
	result.SetUpper3x3(tfrm)
	result.SetTranslation(translateVec)
}

func T3MakeFromQV3(result *Transform3, unitQuat *Quaternion, translateVec *Vector3) {
	var tmpM3_0 Matrix3
	M3MakeFromQ(&tmpM3_0, unitQuat)
	result.SetUpper3x3(&tmpM3_0)
	result.SetTranslation(translateVec)
}

func (t *Transform3) SetCol(col int, vec *Vector3) {
	switch col {
	case 0:
		V3Copy(&t.Col0, vec)
	case 1:
		V3Copy(&t.Col1, vec)
	case 2:
		V3Copy(&t.Col2, vec)
	case 3:
		V3Copy(&t.Col3, vec)
	}
}

func (t *Transform3) SetRow(row int, vec *Vector4) {
	t.Col0.SetElem(row, vec.GetElem(0))
	t.Col1.SetElem(row, vec.GetElem(1))
	t.Col2.SetElem(row, vec.GetElem(2))
	t.Col3.SetElem(row, vec.GetElem(3))
}

func (t *Transform3) SetElem(col, row int, val float32) {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, t, col)
	tmpV3_0.SetElem(row, val)
	t.SetCol(col, &tmpV3_0)
}

func (t *Transform3) GetElem(col, row int) float32 {
	var tmpV3_0 Vector3
	T3GetCol(&tmpV3_0, t, col)
	return tmpV3_0.GetElem(row)
}

func T3GetCol(result *Vector3, tfrm *Transform3, col int) {
	switch col {
	case 0:
		V3Copy(result, &tfrm.Col0)
	case 1:
		V3Copy(result, &tfrm.Col1)
	case 2:
		V3Copy(result, &tfrm.Col2)
	case 3:
		V3Copy(result, &tfrm.Col3)
	}
}

func T3GetRow(result *Vector4, tfrm *Transform3, row int) {
	V4MakeFromElems(result, tfrm.Col0.GetElem(row), tfrm.Col1.GetElem(row), tfrm.Col2.GetElem(row), tfrm.Col3.GetElem(row))
}

func T3Inverse(result, tfrm *Transform3) {
	var tmp0, tmp1, tmp2, inv0, inv1, inv2, tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3, tmpV3_4, tmpV3_5 Vector3
	V3Cross(&tmp0, &tfrm.Col1, &tfrm.Col2)
	V3Cross(&tmp1, &tfrm.Col2, &tfrm.Col0)
	V3Cross(&tmp2, &tfrm.Col0, &tfrm.Col1)
	detinv := (1.0 / V3Dot(&tfrm.Col2, &tmp2))
	V3MakeFromElems(&inv0, (tmp0.X * detinv), (tmp1.X * detinv), (tmp2.X * detinv))
	V3MakeFromElems(&inv1, (tmp0.Y * detinv), (tmp1.Y * detinv), (tmp2.Y * detinv))
	V3MakeFromElems(&inv2, (tmp0.Z * detinv), (tmp1.Z * detinv), (tmp2.Z * detinv))
	V3Copy(&result.Col0, &inv0)
	V3Copy(&result.Col1, &inv1)
	V3Copy(&result.Col2, &inv2)
	V3ScalarMul(&tmpV3_0, &inv0, tfrm.Col3.X)
	V3ScalarMul(&tmpV3_1, &inv1, tfrm.Col3.Y)
	V3ScalarMul(&tmpV3_2, &inv2, tfrm.Col3.Z)
	V3Add(&tmpV3_3, &tmpV3_1, &tmpV3_2)
	V3Add(&tmpV3_4, &tmpV3_0, &tmpV3_3)
	V3Neg(&tmpV3_5, &tmpV3_4)
	V3Copy(&result.Col3, &tmpV3_5)
}

func (t *Transform3) Invert() {
	T3Inverse(t, t)
}

func T3OrthoInverse(result, tfrm *Transform3) {
	var inv0, inv1, inv2, tmpV3_0, tmpV3_1, tmpV3_2, tmpV3_3, tmpV3_4, tmpV3_5 Vector3
	V3MakeFromElems(&inv0, tfrm.Col0.X, tfrm.Col1.X, tfrm.Col2.X)
	V3MakeFromElems(&inv1, tfrm.Col0.Y, tfrm.Col1.Y, tfrm.Col2.Y)
	V3MakeFromElems(&inv2, tfrm.Col0.Z, tfrm.Col1.Z, tfrm.Col2.Z)
	V3Copy(&result.Col0, &inv0)
	V3Copy(&result.Col1, &inv1)
	V3Copy(&result.Col2, &inv2)
	V3ScalarMul(&tmpV3_0, &inv0, tfrm.Col3.X)
	V3ScalarMul(&tmpV3_1, &inv1, tfrm.Col3.Y)
	V3ScalarMul(&tmpV3_2, &inv2, tfrm.Col3.Z)
	V3Add(&tmpV3_3, &tmpV3_1, &tmpV3_2)
	V3Add(&tmpV3_4, &tmpV3_0, &tmpV3_3)
	V3Neg(&tmpV3_5, &tmpV3_4)
	V3Copy(&result.Col3, &tmpV3_5)
}

func (t *Transform3) OrthoInvert() {
	T3OrthoInverse(t, t)
}

func T3AbsPerElem(result, tfrm *Transform3) {
	V3AbsPerElem(&result.Col0, &tfrm.Col0)
	V3AbsPerElem(&result.Col1, &tfrm.Col1)
	V3AbsPerElem(&result.Col2, &tfrm.Col2)
	V3AbsPerElem(&result.Col3, &tfrm.Col3)
}

func T3MulV3(result *Vector3, tfrm *Transform3, vec *Vector3) {
	tmpX := ((tfrm.Col0.X * vec.X) + (tfrm.Col1.X * vec.Y)) + (tfrm.Col2.X * vec.Z)
	tmpY := ((tfrm.Col0.Y * vec.X) + (tfrm.Col1.Y * vec.Y)) + (tfrm.Col2.Y * vec.Z)
	tmpZ := ((tfrm.Col0.Z * vec.X) + (tfrm.Col1.Z * vec.Y)) + (tfrm.Col2.Z * vec.Z)
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func T3MulP3(result *Point3, tfrm *Transform3, pnt *Point3) {
	tmpX := ((((tfrm.Col0.X * pnt.X) + (tfrm.Col1.X * pnt.Y)) + (tfrm.Col2.X * pnt.Z)) + tfrm.Col3.X)
	tmpY := ((((tfrm.Col0.Y * pnt.X) + (tfrm.Col1.Y * pnt.Y)) + (tfrm.Col2.Y * pnt.Z)) + tfrm.Col3.Y)
	tmpZ := ((((tfrm.Col0.Z * pnt.X) + (tfrm.Col1.Z * pnt.Y)) + (tfrm.Col2.Z * pnt.Z)) + tfrm.Col3.Z)
	P3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func T3Mul(result, tfrm0, tfrm1 *Transform3) {
	var tmpResult Transform3
	var tmpP3_0, tmpP3_1 Point3
	T3MulV3(&tmpResult.Col0, tfrm0, &tfrm1.Col0)
	T3MulV3(&tmpResult.Col1, tfrm0, &tfrm1.Col1)
	T3MulV3(&tmpResult.Col2, tfrm0, &tfrm1.Col2)
	P3MakeFromV3(&tmpP3_0, &tfrm1.Col3)
	T3MulP3(&tmpP3_1, tfrm0, &tmpP3_0)
	V3MakeFromP3(&tmpResult.Col3, &tmpP3_1)
	T3Copy(result, &tmpResult)
}

func T3MulPerElem(result, tfrm0, tfrm1 *Transform3) {
	V3MulPerElem(&result.Col0, &tfrm0.Col0, &tfrm1.Col0)
	V3MulPerElem(&result.Col1, &tfrm0.Col1, &tfrm1.Col1)
	V3MulPerElem(&result.Col2, &tfrm0.Col2, &tfrm1.Col2)
	V3MulPerElem(&result.Col3, &tfrm0.Col3, &tfrm1.Col3)
}

func T3MakeIdentity(result *Transform3) {
	V3MakeXAxis(&result.Col0)
	V3MakeYAxis(&result.Col1)
	V3MakeZAxis(&result.Col2)
	V3MakeFromScalar(&result.Col3, 0.0)
}

func (m *Transform3) SetUpper3x3(tfrm *Matrix3) {
	V3Copy(&m.Col0, &tfrm.Col0)
	V3Copy(&m.Col1, &tfrm.Col1)
	V3Copy(&m.Col2, &tfrm.Col2)
}

func T3GetUpper3x3(result *Matrix3, tfrm *Transform3) {
	M3MakeFromCols(result, &tfrm.Col0, &tfrm.Col1, &tfrm.Col2)
}

func (t *Transform3) Upper3x3(result *Matrix3) {
	T3GetUpper3x3(result, t)
}

func (t *Transform3) SetTranslation(translateVec *Vector3) {
	V3Copy(&t.Col3, translateVec)
}

func T3GetTranslation(result *Vector3, tfrm *Transform3) {
	V3Copy(result, &tfrm.Col3)
}

func (t *Transform3) Translation(result *Vector3) {
	T3GetTranslation(result, t)
}

func T3MakeRotationX(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeXAxis(&result.Col0)
	V3MakeFromElems(&result.Col1, 0.0, c, s)
	V3MakeFromElems(&result.Col2, 0.0, -s, c)
	V3MakeFromScalar(&result.Col3, 0.0)
}

func T3MakeRotationY(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.Col0, c, 0.0, -s)
	V3MakeYAxis(&result.Col1)
	V3MakeFromElems(&result.Col2, s, 0.0, c)
	V3MakeFromScalar(&result.Col3, 0.0)
}

func T3MakeRotationZ(result *Transform3, radians float32) {
	s := sin(radians)
	c := cos(radians)
	V3MakeFromElems(&result.Col0, c, s, 0.0)
	V3MakeFromElems(&result.Col1, -s, c, 0.0)
	V3MakeZAxis(&result.Col2)
	V3MakeFromScalar(&result.Col3, 0.0)
}

func T3MakeRotationZYX(result *Transform3, radiansXYZ *Vector3) {
	sX := sin(radiansXYZ.X)
	cX := cos(radiansXYZ.X)
	sY := sin(radiansXYZ.Y)
	cY := cos(radiansXYZ.Y)
	sZ := sin(radiansXYZ.Z)
	cZ := cos(radiansXYZ.Z)
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)
	V3MakeFromElems(&result.Col0, (cZ * cY), (sZ * cY), -sY)
	V3MakeFromElems(&result.Col1, ((tmp0 * sX) - (sZ * cX)), ((tmp1 * sX) + (cZ * cX)), (cY * sX))
	V3MakeFromElems(&result.Col2, ((tmp0 * cX) + (sZ * sX)), ((tmp1 * cX) - (cZ * sX)), (cY * cX))
	V3MakeFromScalar(&result.Col3, 0.0)
}

func T3MakeRotationAxis(result *Transform3, radians float32, unitVec *Vector3) {
	var tmpM3_0 Matrix3
	var tmpV3_0 Vector3
	M3MakeRotationAxis(&tmpM3_0, radians, unitVec)
	V3MakeFromScalar(&tmpV3_0, 0.0)
	T3MakeFromM3V3(result, &tmpM3_0, &tmpV3_0)
}

func T3MakeRotationQ(result *Transform3, unitQuat *Quaternion) {
	var tmpM3_0 Matrix3
	var tmpV3_0 Vector3
	M3MakeFromQ(&tmpM3_0, unitQuat)
	V3MakeFromScalar(&tmpV3_0, 0.0)
	T3MakeFromM3V3(result, &tmpM3_0, &tmpV3_0)
}

func T3MakeScale(result *Transform3, scaleVec *Vector3) {
	V3MakeFromElems(&result.Col0, scaleVec.X, 0.0, 0.0)
	V3MakeFromElems(&result.Col1, 0.0, scaleVec.Y, 0.0)
	V3MakeFromElems(&result.Col2, 0.0, 0.0, scaleVec.Z)
	V3MakeFromScalar(&result.Col3, 0.0)
}

func T3AppendScale(result, tfrm *Transform3, scaleVec *Vector3) {
	V3ScalarMul(&result.Col0, &tfrm.Col0, scaleVec.X)
	V3ScalarMul(&result.Col1, &tfrm.Col1, scaleVec.Y)
	V3ScalarMul(&result.Col2, &tfrm.Col2, scaleVec.Z)
	V3Copy(&result.Col3, &tfrm.Col3)
}

func T3PrependScale(result *Transform3, scaleVec *Vector3, tfrm *Transform3) {
	V3MulPerElem(&result.Col0, &tfrm.Col0, scaleVec)
	V3MulPerElem(&result.Col1, &tfrm.Col1, scaleVec)
	V3MulPerElem(&result.Col2, &tfrm.Col2, scaleVec)
	V3MulPerElem(&result.Col3, &tfrm.Col3, scaleVec)
}

func T3MakeTranslation(result *Transform3, translateVec *Vector3) {
	V3MakeXAxis(&result.Col0)
	V3MakeYAxis(&result.Col1)
	V3MakeZAxis(&result.Col2)
	V3Copy(&result.Col3, translateVec)
}

func T3Select(result, tfrm0, tfrm1 *Transform3, select1 int) {
	V3Select(&result.Col0, &tfrm0.Col0, &tfrm1.Col0, select1)
	V3Select(&result.Col1, &tfrm0.Col1, &tfrm1.Col1, select1)
	V3Select(&result.Col2, &tfrm0.Col2, &tfrm1.Col2, select1)
	V3Select(&result.Col3, &tfrm0.Col3, &tfrm1.Col3, select1)
}

func (t *Transform3) String() string {
	var tmpV4_0, tmpV4_1, tmpV4_2 Vector4
	T3GetRow(&tmpV4_0, t, 0)
	T3GetRow(&tmpV4_1, t, 1)
	T3GetRow(&tmpV4_2, t, 2)
	return tmpV4_0.String() + tmpV4_1.String() + tmpV4_2.String()
}


func QMakeFromM3(result *Quaternion, tfrm *Matrix3) {
	xx := tfrm.Col0.X
	yx := tfrm.Col0.Y
	zx := tfrm.Col0.Z
	xy := tfrm.Col1.X
	yy := tfrm.Col1.Y
	zy := tfrm.Col1.Z
	xz := tfrm.Col2.X
	yz := tfrm.Col2.Y
	zz := tfrm.Col2.Z

	trace := ((xx + yy) + zz)

	negTrace := (trace < 0.0)
	ZgtX := zz > xx
	ZgtY := zz > yy
	YgtX := yy > xx
	largestXorY := (!ZgtX || !ZgtY) && negTrace
	largestYorZ := (YgtX || ZgtX) && negTrace
	largestZorX := (ZgtY || !YgtX) && negTrace

	if largestXorY {
		zz = -zz
		xy = -xy
	}
	if largestYorZ {
		xx = -xx
		yz = -yz
	}
	if largestZorX {
		yy = -yy
		zx = -zx
	}

	radicand := (((xx + yy) + zz) + 1.0)
	scale := (0.5 * (1.0 / sqrt(radicand)))

	tmpx := ((zy - yz) * scale)
	tmpy := ((xz - zx) * scale)
	tmpz := ((yx - xy) * scale)
	tmpw := (radicand * scale)
	qx := tmpx
	qy := tmpy
	qz := tmpz
	qw := tmpw

	if largestXorY {
		qx = tmpw
		qy = tmpz
		qz = tmpy
		qw = tmpx
	}
	if largestYorZ {
		tmpx = qx
		tmpz = qz
		qx = qy
		qy = tmpx
		qz = qw
		qw = tmpz
	}

	result.X = qx
	result.Y = qy
	result.Z = qz
	result.W = qw
}

func V3Outer(result *Matrix3, tfrm0, tfrm1 *Vector3) {
	V3ScalarMul(&result.Col0, tfrm0, tfrm1.X)
	V3ScalarMul(&result.Col1, tfrm0, tfrm1.Y)
	V3ScalarMul(&result.Col2, tfrm0, tfrm1.Z)
}

func V4Outer(result *Matrix4, tfrm0, tfrm1 *Vector4) {
	V4ScalarMul(&result.Col0, tfrm0, tfrm1.X)
	V4ScalarMul(&result.Col1, tfrm0, tfrm1.Y)
	V4ScalarMul(&result.Col2, tfrm0, tfrm1.Z)
	V4ScalarMul(&result.Col3, tfrm0, tfrm1.W)
}

func V3RowMul(result *Vector3, vec *Vector3, mat *Matrix3) {
	tmpX := (((vec.X * mat.Col0.X) + (vec.Y * mat.Col0.Y)) + (vec.Z * mat.Col0.Z))
	tmpY := (((vec.X * mat.Col1.X) + (vec.Y * mat.Col1.Y)) + (vec.Z * mat.Col1.Z))
	tmpZ := (((vec.X * mat.Col2.X) + (vec.Y * mat.Col2.Y)) + (vec.Z * mat.Col2.Z))
	V3MakeFromElems(result, tmpX, tmpY, tmpZ)
}

func V3CrossMatrix(result *Matrix3, vec *Vector3) {
	V3MakeFromElems(&result.Col0, 0.0, vec.Z, -vec.Y)
	V3MakeFromElems(&result.Col1, -vec.Z, 0.0, vec.X)
	V3MakeFromElems(&result.Col2, vec.Y, -vec.X, 0.0)
}

func V3CrossMatrixMul(result *Matrix3, vec *Vector3, mat *Matrix3) {
	var tmpV3_0, tmpV3_1, tmpV3_2 Vector3
	V3Cross(&tmpV3_0, vec, &mat.Col0)
	V3Cross(&tmpV3_1, vec, &mat.Col1)
	V3Cross(&tmpV3_2, vec, &mat.Col2)
	M3MakeFromCols(result, &tmpV3_0, &tmpV3_1, &tmpV3_2)
}

*/
