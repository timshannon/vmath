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

func (m *Matrix3) Elem(col, row int) float32 {
	return m[col*3+row]
}

func (m *Matrix3) Col(result *Vector3, col int) {
	switch col {
	case 0:
		copy(result[:], m[m3col0:m3col1-1])
	case 1:
		copy(result[:], m[m3col1:m3col2-1])
	case 2:
		copy(result[:], m[m3col2:])
	}
}

func (mat *Matrix3) Row(result *Vector3, row int) {
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

	mat.Col(&col0, 0)
	mat.Col(&col1, 1)
	mat.Col(&col2, 2)

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

func (m *Matrix3) InverseSelf() {
	m.Inverse(m)
}

func (m *Matrix3) Determinant() float32 {
	var col0, col1, col2, tmp Vector3
	m.Col(&col0, 0)
	m.Col(&col1, 0)
	m.Col(&col2, 0)

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

func (result *Matrix3) AddToSelf(mat *Matrix3) {
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

func (result *Matrix3) SubFromSelf(mat *Matrix3) {
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

//Matrix 4
const (
	m4col0 = 0
	m4col1 = 4
	m4col2 = 8
	m4col3 = 12
)

func (result *Matrix4) MakeFromScalar(scalar float32) {
	result[m4col0+x] = scalar
	result[m4col0+y] = scalar
	result[m4col0+z] = scalar
	result[m4col0+w] = scalar

	result[m4col1+x] = scalar
	result[m4col1+y] = scalar
	result[m4col1+z] = scalar
	result[m4col1+w] = scalar

	result[m4col2+x] = scalar
	result[m4col2+y] = scalar
	result[m4col2+z] = scalar
	result[m4col2+w] = scalar

	result[m4col3+x] = scalar
	result[m4col3+y] = scalar
	result[m4col3+z] = scalar
	result[m4col3+w] = scalar

}

func (result *Matrix4) MakeFromT3(trns *Transform3) {
	result[m4col0+x] = trns[t3col0+x]
	result[m4col0+y] = trns[t3col0+y]
	result[m4col0+z] = trns[t3col0+z]
	result[m4col0+w] = 0.0

	result[m4col1+x] = trns[t3col1+x]
	result[m4col1+y] = trns[t3col1+y]
	result[m4col1+z] = trns[t3col1+z]
	result[m4col1+w] = 0.0

	result[m4col2+x] = trns[t3col2+x]
	result[m4col2+y] = trns[t3col2+y]
	result[m4col2+z] = trns[t3col2+z]
	result[m4col2+w] = 0.0

	result[m4col3+x] = trns[t3col3+x]
	result[m4col3+y] = trns[t3col3+y]
	result[m4col3+z] = trns[t3col3+z]
	result[m4col3+w] = 1.0

}

func (m *Matrix4) SetCol(col int, vec *Vector4) {
	switch col {
	case 0:
		m[m4col0+x] = vec[x]
		m[m4col0+y] = vec[y]
		m[m4col0+z] = vec[z]
		m[m4col0+w] = vec[w]
	case 1:
		m[m4col1+x] = vec[x]
		m[m4col1+y] = vec[y]
		m[m4col1+z] = vec[z]
		m[m4col1+w] = vec[w]
	case 2:
		m[m4col2+x] = vec[x]
		m[m4col2+y] = vec[y]
		m[m4col2+z] = vec[z]
		m[m4col2+w] = vec[w]
	case 3:
		m[m4col3+x] = vec[x]
		m[m4col3+y] = vec[y]
		m[m4col3+z] = vec[z]
		m[m4col3+w] = vec[w]
	}
}

func (result *Matrix4) MakeFromCols(col0, col1, col2, col3 *Vector4) {
	result.SetCol(0, col0)
	result.SetCol(1, col1)
	result.SetCol(2, col2)
	result.SetCol(3, col3)
}

func (result *Matrix4) MakeFromM3V3(mat *Matrix3, translateVec *Vector3) {
	result[m4col0+x] = mat[m3col0+x]
	result[m4col0+y] = mat[m3col0+y]
	result[m4col0+z] = mat[m3col0+z]
	result[m4col0+w] = 0.0

	result[m4col1+x] = mat[m3col1+x]
	result[m4col1+y] = mat[m3col1+y]
	result[m4col1+z] = mat[m3col1+z]
	result[m4col1+w] = 0.0

	result[m4col2+x] = mat[m3col2+x]
	result[m4col2+y] = mat[m3col2+y]
	result[m4col2+z] = mat[m3col2+z]
	result[m4col2+w] = 0.0

	result[m4col3+x] = translateVec[x]
	result[m4col3+y] = translateVec[y]
	result[m4col3+z] = translateVec[z]
	result[m4col3+w] = 1.0

}

func (result *Matrix4) MakeFromQV3(unitQuat *Quaternion, translateVec *Vector3) {
	var mat Matrix3
	mat.MakeFromQ(unitQuat)
	result.MakeFromM3V3(&mat, translateVec)
}

func (m *Matrix4) SetRow(row int, vec *Vector4) {
	m[m4col0+row] = vec[x]
	m[m4col1+row] = vec[y]
	m[m4col2+row] = vec[z]
	m[m4col3+row] = vec[w]
}

func (m *Matrix4) SetElem(col, row int, val float32) {
	m[col*4+row] = val
}

func (m *Matrix4) Elem(col, row int) float32 {
	return m[col*4+row]
}

func (m *Matrix4) Col(result *Vector4, col int) {
	switch col {
	case 0:
		copy(result[:], m[m4col0:m4col1-1])
	case 1:
		copy(result[:], m[m4col1:m4col2-1])
	case 2:
		copy(result[:], m[m4col2:m4col3-1])
	case 3:
		copy(result[:], m[m4col3:])

	}
}

func (mat *Matrix4) Row(result *Vector4, row int) {
	result[x] = mat[m4col0+row]
	result[y] = mat[m4col1+row]
	result[z] = mat[m4col2+row]
	result[w] = mat[m4col3+row]
}

func (result *Matrix4) Transpose(mat *Matrix4) {
	result[m4col0+x] = mat[m4col0+x]
	result[m4col0+y] = mat[m4col1+x]
	result[m4col0+z] = mat[m4col2+x]
	result[m4col0+w] = mat[m4col3+x]

	result[m4col1+x] = mat[m4col0+y]
	result[m4col1+y] = mat[m4col1+y]
	result[m4col1+z] = mat[m4col2+y]
	result[m4col1+w] = mat[m4col3+y]

	result[m4col2+x] = mat[m4col0+z]
	result[m4col2+y] = mat[m4col1+z]
	result[m4col2+z] = mat[m4col2+z]
	result[m4col2+w] = mat[m4col3+z]

	result[m4col3+x] = mat[m4col0+w]
	result[m4col3+y] = mat[m4col1+w]
	result[m4col3+z] = mat[m4col2+w]
	result[m4col3+w] = mat[m4col3+w]

}

func (m *Matrix4) TransposeSelf() {
	tmp := *m
	m.Transpose(&tmp)
}

func (result *Matrix4) Inverse(mat *Matrix4) {
	var res0, res1, res2, res3 Vector4
	mA := mat[m4col0+x]
	mB := mat[m4col0+y]
	mC := mat[m4col0+z]
	mD := mat[m4col0+w]
	mE := mat[m4col1+x]
	mF := mat[m4col1+y]
	mG := mat[m4col1+z]
	mH := mat[m4col1+w]
	mI := mat[m4col2+x]
	mJ := mat[m4col2+y]
	mK := mat[m4col2+z]
	mL := mat[m4col2+w]
	mM := mat[m4col3+x]
	mN := mat[m4col3+y]
	mO := mat[m4col3+z]
	mP := mat[m4col3+w]
	tmp0 := ((mK * mD) - (mC * mL))
	tmp1 := ((mO * mH) - (mG * mP))
	tmp2 := ((mB * mK) - (mJ * mC))
	tmp3 := ((mF * mO) - (mN * mG))
	tmp4 := ((mJ * mD) - (mB * mL))
	tmp5 := ((mN * mH) - (mF * mP))
	res0[x] = (((mJ * tmp1) - (mL * tmp3)) - (mK * tmp5))
	res0[y] = (((mN * tmp0) - (mP * tmp2)) - (mO * tmp4))
	res0[z] = (((mD * tmp3) + (mC * tmp5)) - (mB * tmp1))
	res0[w] = (((mH * tmp2) + (mG * tmp4)) - (mF * tmp0))
	detInv := (1.0 / ((((mA * res0[x]) + (mE * res0[y])) + (mI * res0[z])) + (mM * res0[w])))
	res1[x] = (mI * tmp1)
	res1[y] = (mM * tmp0)
	res1[z] = (mA * tmp1)
	res1[w] = (mE * tmp0)
	res3[x] = (mI * tmp3)
	res3[y] = (mM * tmp2)
	res3[z] = (mA * tmp3)
	res3[w] = (mE * tmp2)
	res2[x] = (mI * tmp5)
	res2[y] = (mM * tmp4)
	res2[z] = (mA * tmp5)
	res2[w] = (mE * tmp4)
	tmp0 = ((mI * mB) - (mA * mJ))
	tmp1 = ((mM * mF) - (mE * mN))
	tmp2 = ((mI * mD) - (mA * mL))
	tmp3 = ((mM * mH) - (mE * mP))
	tmp4 = ((mI * mC) - (mA * mK))
	tmp5 = ((mM * mG) - (mE * mO))
	res2[x] = (((mL * tmp1) - (mJ * tmp3)) + res2[x])
	res2[y] = (((mP * tmp0) - (mN * tmp2)) + res2[y])
	res2[z] = (((mB * tmp3) - (mD * tmp1)) - res2[z])
	res2[w] = (((mF * tmp2) - (mH * tmp0)) - res2[w])
	res3[x] = (((mJ * tmp5) - (mK * tmp1)) + res3[x])
	res3[y] = (((mN * tmp4) - (mO * tmp0)) + res3[y])
	res3[z] = (((mC * tmp1) - (mB * tmp5)) - res3[z])
	res3[w] = (((mG * tmp0) - (mF * tmp4)) - res3[w])
	res1[x] = (((mK * tmp3) - (mL * tmp5)) - res1[x])
	res1[y] = (((mO * tmp2) - (mP * tmp4)) - res1[y])
	res1[z] = (((mD * tmp5) - (mC * tmp3)) + res1[z])
	res1[w] = (((mH * tmp4) - (mG * tmp2)) + res1[w])

	res0.ScalarMulSelf(detInv)
	result.SetCol(0, &res0)

	res1.ScalarMulSelf(detInv)
	result.SetCol(1, &res1)

	res2.ScalarMulSelf(detInv)
	result.SetCol(2, &res2)

	res3.ScalarMulSelf(detInv)
	result.SetCol(3, &res3)

}

func (result *Matrix4) InverseSelf() {
	result.Inverse(result)
}

func (result *Matrix4) AffineInverse(mat *Matrix4) {
	var affineMat Transform3

	affineMat[t3col0+x] = mat[m4col0+x]
	affineMat[t3col0+y] = mat[m4col0+y]
	affineMat[t3col0+z] = mat[m4col0+z]

	affineMat[t3col1+x] = mat[m4col1+x]
	affineMat[t3col1+y] = mat[m4col1+y]
	affineMat[t3col1+z] = mat[m4col1+z]

	affineMat[t3col2+x] = mat[m4col2+x]
	affineMat[t3col2+y] = mat[m4col2+y]
	affineMat[t3col2+z] = mat[m4col2+z]

	affineMat[t3col3+x] = mat[m4col3+x]
	affineMat[t3col3+y] = mat[m4col3+y]
	affineMat[t3col3+z] = mat[m4col3+z]

	affineMat.InverseSelf()

	result.MakeFromT3(&affineMat)
}

func (result *Matrix4) AffineInverseSelf() {
	result.AffineInverse(result)
}

func (result *Matrix4) OrthoInverse(mat *Matrix4) {
	var affineMat Transform3

	affineMat[t3col0+x] = mat[m4col0+x]
	affineMat[t3col0+y] = mat[m4col0+y]
	affineMat[t3col0+z] = mat[m4col0+z]

	affineMat[t3col1+x] = mat[m4col1+x]
	affineMat[t3col1+y] = mat[m4col1+y]
	affineMat[t3col1+z] = mat[m4col1+z]

	affineMat[t3col2+x] = mat[m4col2+x]
	affineMat[t3col2+y] = mat[m4col2+y]
	affineMat[t3col2+z] = mat[m4col2+z]

	affineMat[t3col3+x] = mat[m4col3+x]
	affineMat[t3col3+y] = mat[m4col3+y]
	affineMat[t3col3+z] = mat[m4col3+z]

	affineMat.OrthoInverseSelf()

	result.MakeFromT3(&affineMat)
}

func (result *Matrix4) OrthoInverseSelf() {
	result.OrthoInverse(result)
}

func (m *Matrix4) Determinant() float32 {
	mA := m[m4col0+x]
	mB := m[m4col0+y]
	mC := m[m4col0+z]
	mD := m[m4col0+w]
	mE := m[m4col1+x]
	mF := m[m4col1+y]
	mG := m[m4col1+z]
	mH := m[m4col1+w]
	mI := m[m4col2+x]
	mJ := m[m4col2+y]
	mK := m[m4col2+z]
	mL := m[m4col2+w]
	mM := m[m4col3+x]
	mN := m[m4col3+y]
	mO := m[m4col3+z]
	mP := m[m4col3+w]
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

func (result *Matrix4) Add(mat0, mat1 *Matrix4) {
	result[m4col0+x] = mat0[m4col0+x] + mat1[m4col0+x]
	result[m4col0+y] = mat0[m4col0+y] + mat1[m4col0+y]
	result[m4col0+z] = mat0[m4col0+z] + mat1[m4col0+z]
	result[m4col0+w] = mat0[m4col0+w] + mat1[m4col0+w]

	result[m4col1+x] = mat0[m4col1+x] + mat1[m4col1+x]
	result[m4col1+y] = mat0[m4col1+y] + mat1[m4col1+y]
	result[m4col1+z] = mat0[m4col1+z] + mat1[m4col1+z]
	result[m4col1+w] = mat0[m4col1+w] + mat1[m4col1+w]

	result[m4col2+x] = mat0[m4col2+x] + mat1[m4col2+x]
	result[m4col2+y] = mat0[m4col2+y] + mat1[m4col2+y]
	result[m4col2+z] = mat0[m4col2+z] + mat1[m4col2+z]
	result[m4col2+w] = mat0[m4col2+w] + mat1[m4col2+w]

	result[m4col3+x] = mat0[m4col3+x] + mat1[m4col3+x]
	result[m4col3+y] = mat0[m4col3+y] + mat1[m4col3+y]
	result[m4col3+z] = mat0[m4col3+z] + mat1[m4col3+z]
	result[m4col3+w] = mat0[m4col3+w] + mat1[m4col3+w]
}

func (result *Matrix4) AddToSelf(mat *Matrix4) {
	result.Add(result, mat)
}

func (result *Matrix4) Sub(mat0, mat1 *Matrix4) {
	result[m4col0+x] = mat0[m4col0+x] - mat1[m4col0+x]
	result[m4col0+y] = mat0[m4col0+y] - mat1[m4col0+y]
	result[m4col0+z] = mat0[m4col0+z] - mat1[m4col0+z]
	result[m4col0+w] = mat0[m4col0+w] - mat1[m4col0+w]

	result[m4col1+x] = mat0[m4col1+x] - mat1[m4col1+x]
	result[m4col1+y] = mat0[m4col1+y] - mat1[m4col1+y]
	result[m4col1+z] = mat0[m4col1+z] - mat1[m4col1+z]
	result[m4col1+w] = mat0[m4col1+w] - mat1[m4col1+w]

	result[m4col2+x] = mat0[m4col2+x] - mat1[m4col2+x]
	result[m4col2+y] = mat0[m4col2+y] - mat1[m4col2+y]
	result[m4col2+z] = mat0[m4col2+z] - mat1[m4col2+z]
	result[m4col2+w] = mat0[m4col2+w] - mat1[m4col2+w]

	result[m4col3+x] = mat0[m4col3+x] - mat1[m4col3+x]
	result[m4col3+y] = mat0[m4col3+y] - mat1[m4col3+y]
	result[m4col3+z] = mat0[m4col3+z] - mat1[m4col3+z]
	result[m4col3+w] = mat0[m4col3+w] - mat1[m4col3+w]
}

func (result *Matrix4) SubFromSelf(mat *Matrix4) {
	result.Sub(result, mat)
}

func (result *Matrix4) Neg(mat *Matrix4) {
	result[m4col0+x] = -mat[m4col0+x]
	result[m4col0+y] = -mat[m4col0+y]
	result[m4col0+z] = -mat[m4col0+z]
	result[m4col0+w] = -mat[m4col0+w]

	result[m4col1+x] = -mat[m4col1+x]
	result[m4col1+y] = -mat[m4col1+y]
	result[m4col1+z] = -mat[m4col1+z]
	result[m4col1+w] = -mat[m4col1+w]

	result[m4col2+x] = -mat[m4col2+x]
	result[m4col2+y] = -mat[m4col2+y]
	result[m4col2+z] = -mat[m4col2+z]
	result[m4col2+w] = -mat[m4col2+w]

	result[m4col3+x] = -mat[m4col3+x]
	result[m4col3+y] = -mat[m4col3+y]
	result[m4col3+z] = -mat[m4col3+z]
	result[m4col3+w] = -mat[m4col3+w]

}

func (m *Matrix4) NegSelf() {
	m.Neg(m)
}

func (result *Matrix4) AbsPerElem(mat *Matrix4) {
	result[m4col0+x] = abs(mat[m4col0+x])
	result[m4col0+y] = abs(mat[m4col0+y])
	result[m4col0+z] = abs(mat[m4col0+z])
	result[m4col0+w] = abs(mat[m4col0+w])

	result[m4col1+x] = abs(mat[m4col1+x])
	result[m4col1+y] = abs(mat[m4col1+y])
	result[m4col1+z] = abs(mat[m4col1+z])
	result[m4col1+w] = abs(mat[m4col1+w])

	result[m4col2+x] = abs(mat[m4col2+x])
	result[m4col2+y] = abs(mat[m4col2+y])
	result[m4col2+z] = abs(mat[m4col2+z])
	result[m4col2+w] = abs(mat[m4col2+w])

	result[m4col3+x] = abs(mat[m4col3+x])
	result[m4col3+y] = abs(mat[m4col3+y])
	result[m4col3+z] = abs(mat[m4col3+z])
	result[m4col3+w] = abs(mat[m4col3+w])
}

func (result *Matrix4) AbsPerElemSelf() {
	result.AbsPerElem(result)
}

func (result *Matrix4) ScalarMul(mat *Matrix4, scalar float32) {
	result[m4col0+x] = mat[m4col0+x] * scalar
	result[m4col0+y] = mat[m4col0+y] * scalar
	result[m4col0+z] = mat[m4col0+z] * scalar
	result[m4col0+w] = mat[m4col0+w] * scalar

	result[m4col1+x] = mat[m4col1+x] * scalar
	result[m4col1+y] = mat[m4col1+y] * scalar
	result[m4col1+z] = mat[m4col1+z] * scalar
	result[m4col1+w] = mat[m4col1+w] * scalar

	result[m4col2+x] = mat[m4col2+x] * scalar
	result[m4col2+y] = mat[m4col2+y] * scalar
	result[m4col2+z] = mat[m4col2+z] * scalar
	result[m4col2+w] = mat[m4col2+w] * scalar

	result[m4col3+x] = mat[m4col3+x] * scalar
	result[m4col3+y] = mat[m4col3+y] * scalar
	result[m4col3+z] = mat[m4col3+z] * scalar
	result[m4col3+w] = mat[m4col3+w] * scalar
}

func (result *Matrix4) ScalarMulSelf(scalar float32) {
	result.ScalarMul(result, scalar)
}

func (result *Vector4) MulM4(vec *Vector4, mat *Matrix4) {
	result[x] = (((mat[m4col0+x] * vec[x]) + (mat[m4col1+x] * vec[y])) + (mat[m4col2+x] * vec[z])) + (mat[m4col3+x] * vec[w])
	result[y] = (((mat[m4col0+y] * vec[x]) + (mat[m4col1+y] * vec[y])) + (mat[m4col2+y] * vec[z])) + (mat[m4col3+y] * vec[w])
	result[z] = (((mat[m4col0+z] * vec[x]) + (mat[m4col1+z] * vec[y])) + (mat[m4col2+z] * vec[z])) + (mat[m4col3+z] * vec[w])
	result[w] = (((mat[m4col0+w] * vec[x]) + (mat[m4col1+w] * vec[y])) + (mat[m4col2+w] * vec[z])) + (mat[m4col3+w] * vec[w])
}

func (result *Vector4) MulM4Self(mat *Matrix4) {
	tmp := *result
	result.MulM4(&tmp, mat)

}

func (result *Vector4) MulM4V3(mat *Matrix4, vec *Vector3) {
	result[x] = ((mat[m4col0+x] * vec[x]) + (mat[m4col1+x] * vec[y])) + (mat[m4col2+x] * vec[z])
	result[y] = ((mat[m4col0+y] * vec[x]) + (mat[m4col1+y] * vec[y])) + (mat[m4col2+y] * vec[z])
	result[z] = ((mat[m4col0+z] * vec[x]) + (mat[m4col1+z] * vec[y])) + (mat[m4col2+z] * vec[z])
	result[w] = ((mat[m4col0+w] * vec[x]) + (mat[m4col1+w] * vec[y])) + (mat[m4col2+w] * vec[z])
}

func (result *Vector4) MulM4P3(mat *Matrix4, pnt *Point3) {
	result[x] = (((mat[m4col0+x] * pnt[x]) + (mat[m4col1+x] * pnt[y])) + (mat[m4col2+x] * pnt[z])) + mat[m4col3+x]
	result[y] = (((mat[m4col0+y] * pnt[x]) + (mat[m4col1+y] * pnt[y])) + (mat[m4col2+y] * pnt[z])) + mat[m4col3+y]
	result[z] = (((mat[m4col0+z] * pnt[x]) + (mat[m4col1+z] * pnt[y])) + (mat[m4col2+z] * pnt[z])) + mat[m4col3+z]
	result[w] = (((mat[m4col0+w] * pnt[x]) + (mat[m4col1+w] * pnt[y])) + (mat[m4col2+w] * pnt[z])) + mat[m4col3+w]
}

func (result *Matrix4) Mul(mat0, mat1 *Matrix4) {
	//M4MulV4(&tmpResult.Col0, mat0, &mat1.Col0)
	result[m4col0+x] = (((mat0[m4col0+x] * mat1[m4col0+x]) + (mat0[m4col1+x] * mat1[m4col0+y])) + (mat0[m4col2+x] * mat1[m4col0+z])) + (mat0[m4col3+x] * mat1[m4col0+w])
	result[m4col0+y] = (((mat0[m4col0+y] * mat1[m4col0+x]) + (mat0[m4col1+y] * mat1[m4col0+y])) + (mat0[m4col2+y] * mat1[m4col0+z])) + (mat0[m4col3+y] * mat1[m4col0+w])
	result[m4col0+z] = (((mat0[m4col0+z] * mat1[m4col0+x]) + (mat0[m4col1+z] * mat1[m4col0+y])) + (mat0[m4col2+z] * mat1[m4col0+z])) + (mat0[m4col3+z] * mat1[m4col0+w])
	result[m4col0+w] = (((mat0[m4col0+w] * mat1[m4col0+x]) + (mat0[m4col1+w] * mat1[m4col0+y])) + (mat0[m4col2+w] * mat1[m4col0+z])) + (mat0[m4col3+w] * mat1[m4col0+w])

	//M4MulV4(&tmpResult.Col1, mat0, &mat1.Col1)
	result[m4col1+x] = (((mat0[m4col0+x] * mat1[m4col1+x]) + (mat0[m4col1+x] * mat1[m4col1+y])) + (mat0[m4col2+x] * mat1[m4col1+z])) + (mat0[m4col3+x] * mat1[m4col1+w])
	result[m4col1+y] = (((mat0[m4col0+y] * mat1[m4col1+x]) + (mat0[m4col1+y] * mat1[m4col1+y])) + (mat0[m4col2+y] * mat1[m4col1+z])) + (mat0[m4col3+y] * mat1[m4col1+w])
	result[m4col1+z] = (((mat0[m4col0+z] * mat1[m4col1+x]) + (mat0[m4col1+z] * mat1[m4col1+y])) + (mat0[m4col2+z] * mat1[m4col1+z])) + (mat0[m4col3+z] * mat1[m4col1+w])
	result[m4col1+w] = (((mat0[m4col0+w] * mat1[m4col1+x]) + (mat0[m4col1+w] * mat1[m4col1+y])) + (mat0[m4col2+w] * mat1[m4col1+z])) + (mat0[m4col3+w] * mat1[m4col1+w])

	//M4MulV4(&tmpResult.Col2, mat0, &mat1.Col2)
	result[m4col2+x] = (((mat0[m4col0+x] * mat1[m4col2+x]) + (mat0[m4col1+x] * mat1[m4col2+y])) + (mat0[m4col2+x] * mat1[m4col2+z])) + (mat0[m4col3+x] * mat1[m4col2+w])
	result[m4col2+y] = (((mat0[m4col0+y] * mat1[m4col2+x]) + (mat0[m4col1+y] * mat1[m4col2+y])) + (mat0[m4col2+y] * mat1[m4col2+z])) + (mat0[m4col3+y] * mat1[m4col2+w])
	result[m4col2+z] = (((mat0[m4col0+z] * mat1[m4col2+x]) + (mat0[m4col1+z] * mat1[m4col2+y])) + (mat0[m4col2+z] * mat1[m4col2+z])) + (mat0[m4col3+z] * mat1[m4col2+w])
	result[m4col2+w] = (((mat0[m4col0+w] * mat1[m4col2+x]) + (mat0[m4col1+w] * mat1[m4col2+y])) + (mat0[m4col2+w] * mat1[m4col2+z])) + (mat0[m4col3+w] * mat1[m4col2+w])

	//M4MulV4(&tmpResult.Col3, mat0, &mat1.Col3)
	result[m4col3+x] = (((mat0[m4col0+x] * mat1[m4col3+x]) + (mat0[m4col1+x] * mat1[m4col3+y])) + (mat0[m4col2+x] * mat1[m4col3+z])) + (mat0[m4col3+x] * mat1[m4col3+w])
	result[m4col3+y] = (((mat0[m4col0+y] * mat1[m4col3+x]) + (mat0[m4col1+y] * mat1[m4col3+y])) + (mat0[m4col2+y] * mat1[m4col3+z])) + (mat0[m4col3+y] * mat1[m4col3+w])
	result[m4col3+z] = (((mat0[m4col0+z] * mat1[m4col3+x]) + (mat0[m4col1+z] * mat1[m4col3+y])) + (mat0[m4col2+z] * mat1[m4col3+z])) + (mat0[m4col3+z] * mat1[m4col3+w])
	result[m4col3+w] = (((mat0[m4col0+w] * mat1[m4col3+x]) + (mat0[m4col1+w] * mat1[m4col3+y])) + (mat0[m4col2+w] * mat1[m4col3+z])) + (mat0[m4col3+w] * mat1[m4col3+w])

}

func (result *Matrix4) MulSelf(mat *Matrix4) {
	tmp := *result
	result.Mul(result, &tmp)
}

func (result *Matrix4) MulT3(mat *Matrix4, tfrm *Transform3) {
	//M4MulV3(&tmpResult.Col0, mat, &tfrm1.Col0)
	result[m4col0+x] = ((mat[m4col0+x] * tfrm[t3col0+x]) + (mat[m4col1+x] * tfrm[t3col0+y])) + (mat[m4col2+x] * tfrm[t3col0+z])
	result[m4col0+y] = ((mat[m4col0+y] * tfrm[t3col0+x]) + (mat[m4col1+y] * tfrm[t3col0+y])) + (mat[m4col2+y] * tfrm[t3col0+z])
	result[m4col0+z] = ((mat[m4col0+z] * tfrm[t3col0+x]) + (mat[m4col1+z] * tfrm[t3col0+y])) + (mat[m4col2+z] * tfrm[t3col0+z])
	result[m4col0+w] = ((mat[m4col0+w] * tfrm[t3col0+x]) + (mat[m4col1+w] * tfrm[t3col0+y])) + (mat[m4col2+w] * tfrm[t3col0+z])

	//M4MulV3(&tmpResult.Col1, mat, &tfrm1.Col1)
	result[m4col1+x] = ((mat[m4col0+x] * tfrm[t3col1+x]) + (mat[m4col1+x] * tfrm[t3col1+y])) + (mat[m4col2+x] * tfrm[t3col1+z])
	result[m4col1+y] = ((mat[m4col0+y] * tfrm[t3col1+x]) + (mat[m4col1+y] * tfrm[t3col1+y])) + (mat[m4col2+y] * tfrm[t3col1+z])
	result[m4col1+z] = ((mat[m4col0+z] * tfrm[t3col1+x]) + (mat[m4col1+z] * tfrm[t3col1+y])) + (mat[m4col2+z] * tfrm[t3col1+z])
	result[m4col1+w] = ((mat[m4col0+w] * tfrm[t3col1+x]) + (mat[m4col1+w] * tfrm[t3col1+y])) + (mat[m4col2+w] * tfrm[t3col1+z])

	//M4MulV3(&tmpResult.Col2, mat, &tfrm1.Col2)
	result[m4col2+x] = ((mat[m4col0+x] * tfrm[t3col2+x]) + (mat[m4col1+x] * tfrm[t3col2+y])) + (mat[m4col2+x] * tfrm[t3col2+z])
	result[m4col2+y] = ((mat[m4col0+y] * tfrm[t3col2+x]) + (mat[m4col1+y] * tfrm[t3col2+y])) + (mat[m4col2+y] * tfrm[t3col2+z])
	result[m4col2+z] = ((mat[m4col0+z] * tfrm[t3col2+x]) + (mat[m4col1+z] * tfrm[t3col2+y])) + (mat[m4col2+z] * tfrm[t3col2+z])
	result[m4col2+w] = ((mat[m4col0+w] * tfrm[t3col2+x]) + (mat[m4col1+w] * tfrm[t3col2+y])) + (mat[m4col2+w] * tfrm[t3col2+z])

	result[m4col3+x] = (((mat[m4col0+x] * tfrm[t3col3+x]) + (mat[m4col1+x] * tfrm[t3col3+y])) + (mat[m4col2+x] * tfrm[t3col3+z])) + mat[m4col3+x]
	result[m4col3+y] = (((mat[m4col0+y] * tfrm[t3col3+x]) + (mat[m4col1+y] * tfrm[t3col3+y])) + (mat[m4col2+y] * tfrm[t3col3+z])) + mat[m4col3+y]
	result[m4col3+z] = (((mat[m4col0+z] * tfrm[t3col3+x]) + (mat[m4col1+z] * tfrm[t3col3+y])) + (mat[m4col2+z] * tfrm[t3col3+z])) + mat[m4col3+z]
	result[m4col3+w] = (((mat[m4col0+w] * tfrm[t3col3+x]) + (mat[m4col1+w] * tfrm[t3col3+y])) + (mat[m4col2+w] * tfrm[t3col3+z])) + mat[m4col3+w]

}

func (result *Matrix4) MulT3Self(tfrm *Transform3) {
	tmp := *result
	result.MulT3(&tmp, tfrm)
}

func (result *Matrix4) MulPerElem(mat0, mat1 *Matrix4) {
	result[m4col0+x] = mat0[m4col0+x] * mat1[m4col0+x]
	result[m4col0+y] = mat0[m4col0+y] * mat1[m4col0+y]
	result[m4col0+z] = mat0[m4col0+z] * mat1[m4col0+z]
	result[m4col0+w] = mat0[m4col0+w] * mat1[m4col0+w]

	result[m4col1+x] = mat0[m4col1+x] * mat1[m4col1+x]
	result[m4col1+y] = mat0[m4col1+y] * mat1[m4col1+y]
	result[m4col1+z] = mat0[m4col1+z] * mat1[m4col1+z]
	result[m4col1+w] = mat0[m4col1+w] * mat1[m4col1+w]

	result[m4col2+x] = mat0[m4col2+x] * mat1[m4col2+x]
	result[m4col2+y] = mat0[m4col2+y] * mat1[m4col2+y]
	result[m4col2+z] = mat0[m4col2+z] * mat1[m4col2+z]
	result[m4col2+w] = mat0[m4col2+w] * mat1[m4col2+w]

	result[m4col3+x] = mat0[m4col3+x] * mat1[m4col3+x]
	result[m4col3+y] = mat0[m4col3+y] * mat1[m4col3+y]
	result[m4col3+z] = mat0[m4col3+z] * mat1[m4col3+z]
	result[m4col3+w] = mat0[m4col3+w] * mat1[m4col3+w]
}

func (result *Matrix4) MulPerElemSelf(mat *Matrix4) {
	result.MulPerElem(result, mat)

}

func (result *Matrix4) MakeIdentity() {
	//x-axis
	result[m4col0+x] = 1.0
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0
	//y-axis
	result[m4col1+x] = 0.0
	result[m4col1+y] = 1.0
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0
	//z-axis
	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = 1.0
	result[m4col2+w] = 0.0
	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (m *Matrix4) SetUpper3x3(mat3 *Matrix3) {
	m[m4col0+x] = mat3[m3col0+x]
	m[m4col0+y] = mat3[m3col0+y]
	m[m4col0+z] = mat3[m3col0+z]

	m[m4col1+x] = mat3[m3col1+x]
	m[m4col1+y] = mat3[m3col1+y]
	m[m4col1+z] = mat3[m3col1+z]

	m[m4col2+x] = mat3[m3col2+x]
	m[m4col2+y] = mat3[m3col2+y]
	m[m4col2+z] = mat3[m3col2+z]
}

func (m *Matrix4) Upper3x3(result *Matrix3) {
	result[m3col0+x] = m[m4col0+x]
	result[m3col0+y] = m[m4col0+y]
	result[m3col0+z] = m[m4col0+z]

	result[m3col1+x] = m[m4col1+x]
	result[m3col1+y] = m[m4col1+y]
	result[m3col1+z] = m[m4col1+z]

	result[m3col2+x] = m[m4col2+x]
	result[m3col2+y] = m[m4col2+y]
	result[m3col2+z] = m[m4col2+z]
}

func (m *Matrix4) SetTranslation(translateVec *Vector3) {
	m[m4col3+x] = translateVec[x]
	m[m4col3+y] = translateVec[y]
	m[m4col3+z] = translateVec[z]
}

func (m *Matrix4) Translation(result *Vector3) {
	result[x] = m[m4col3+x]
	result[y] = m[m4col3+y]
	result[z] = m[m4col3+z]
}

func (result *Matrix4) MakeRotationX(radians float32) {
	s := sin(radians)
	c := cos(radians)

	//x-axis
	result[m4col0+x] = 1.0
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	result[m4col1+x] = 0.0
	result[m4col1+y] = c
	result[m4col1+z] = s
	result[m4col1+w] = 0.0

	result[m4col2+x] = 0.0
	result[m4col2+y] = -s
	result[m4col2+z] = c
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (result *Matrix4) MakeRotationY(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[m4col0+x] = c
	result[m4col0+y] = 0.0
	result[m4col0+z] = -s
	result[m4col0+w] = 0.0

	//y-axis
	result[m4col1+x] = 0.0
	result[m4col1+y] = 1.0
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	result[m4col2+x] = s
	result[m4col2+y] = 0.0
	result[m4col2+z] = c
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (result *Matrix4) MakeRotationZ(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[m4col0+x] = c
	result[m4col0+y] = s
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	result[m4col1+x] = -s
	result[m4col1+y] = c
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	//z-axis
	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = 1.0
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (result *Matrix4) MakeRotationZYX(radiansXYZ *Vector3) {
	sX := sin(radiansXYZ[x])
	cX := cos(radiansXYZ[x])
	sY := sin(radiansXYZ[y])
	cY := cos(radiansXYZ[y])
	sZ := sin(radiansXYZ[z])
	cZ := cos(radiansXYZ[z])
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)

	result[m4col0+x] = (cZ * cY)
	result[m4col0+y] = (sZ * cY)
	result[m4col0+z] = -sY
	result[m4col0+w] = 0.0

	result[m4col1+x] = ((tmp0 * sX) - (sZ * cX))
	result[m4col1+y] = ((tmp1 * sX) + (cZ * cX))
	result[m4col1+z] = (cY * sX)
	result[m4col1+w] = 0.0

	result[m4col2+x] = ((tmp0 * cX) + (sZ * sX))
	result[m4col2+y] = ((tmp1 * cX) - (cZ * sX))
	result[m4col2+z] = (cY * cX)
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (result *Matrix4) MakeRotationAxis(radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	X := unitVec[x]
	Y := unitVec[y]
	Z := unitVec[z]
	xy := X * Y
	yz := Y * Z
	zx := Z * X
	oneMinusC := 1.0 - c

	result[m4col0+x] = (((X * X) * oneMinusC) + c)
	result[m4col0+y] = ((xy * oneMinusC) + (Z * s))
	result[m4col0+z] = ((zx * oneMinusC) - (Y * s))
	result[m4col0+w] = 0.0

	result[m4col1+x] = ((xy * oneMinusC) - (Z * s))
	result[m4col1+y] = (((Y * Y) * oneMinusC) + c)
	result[m4col1+z] = ((yz * oneMinusC) + (X * s))
	result[m4col1+w] = 0.0

	result[m4col2+x] = ((zx * oneMinusC) + (Y * s))
	result[m4col2+y] = ((yz * oneMinusC) - (X * s))
	result[m4col2+z] = (((Z * Z) * oneMinusC) + c)
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0

}

func (result *Matrix4) MakeRotationQ(unitQuat *Quaternion) {
	var tmpT3 Transform3

	tmpT3.MakeRotationQ(unitQuat)
	result.MakeFromT3(&tmpT3)
}

func (result *Matrix4) MakeScale(scaleVec *Vector3) {
	result[m4col0+x] = scaleVec[x]
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	result[m4col1+x] = 0.0
	result[m4col1+y] = scaleVec[y]
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = scaleVec[z]
	result[m4col2+w] = 0.0

	//w-axis
	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = 0.0
	result[m4col3+w] = 1.0
}

func (result *Matrix4) AppendScale(mat *Matrix4, scaleVec *Vector3) {
	result[m4col0+x] = mat[m4col0+x] * scaleVec[x]
	result[m4col0+y] = mat[m4col0+y] * scaleVec[x]
	result[m4col0+z] = mat[m4col0+z] * scaleVec[x]
	result[m4col0+w] = mat[m4col0+w] * scaleVec[x]

	result[m4col1+x] = mat[m4col1+x] * scaleVec[y]
	result[m4col1+y] = mat[m4col1+y] * scaleVec[y]
	result[m4col1+z] = mat[m4col1+z] * scaleVec[y]
	result[m4col1+w] = mat[m4col1+w] * scaleVec[y]

	result[m4col2+x] = mat[m4col2+x] * scaleVec[z]
	result[m4col2+y] = mat[m4col2+y] * scaleVec[z]
	result[m4col2+z] = mat[m4col2+z] * scaleVec[z]
	result[m4col2+w] = mat[m4col2+w] * scaleVec[z]

	result[m4col3+x] = mat[m4col3+x]
	result[m4col3+y] = mat[m4col3+y]
	result[m4col3+z] = mat[m4col3+z]
	result[m4col3+w] = mat[m4col3+w]

}

func (result *Matrix4) AppendScaleSelf(scaleVec *Vector3) {
	result.AppendScale(result, scaleVec)
}

func (result *Matrix4) PrependScale(scaleVec *Vector3, mat *Matrix4) {
	result[m4col0+x] = mat[m4col0+x] * scaleVec[x]
	result[m4col0+y] = mat[m4col0+y] * scaleVec[y]
	result[m4col0+z] = mat[m4col0+z] * scaleVec[z]
	result[m4col0+w] = mat[m4col0+w] * 1.0

	result[m4col1+x] = mat[m4col1+x] * scaleVec[x]
	result[m4col1+y] = mat[m4col1+y] * scaleVec[y]
	result[m4col1+z] = mat[m4col1+z] * scaleVec[z]
	result[m4col1+w] = mat[m4col1+w] * 1.0

	result[m4col2+x] = mat[m4col2+x] * scaleVec[x]
	result[m4col2+y] = mat[m4col2+y] * scaleVec[y]
	result[m4col2+z] = mat[m4col2+z] * scaleVec[z]
	result[m4col2+w] = mat[m4col2+w] * 1.0

	result[m4col3+x] = mat[m4col3+x] * scaleVec[x]
	result[m4col3+y] = mat[m4col3+y] * scaleVec[y]
	result[m4col3+z] = mat[m4col3+z] * scaleVec[z]
	result[m4col3+w] = mat[m4col3+w] * 1.0

}

func (result *Matrix4) PrependScaleSelf(scaleVec *Vector3) {
	result.PrependScale(scaleVec, result)
}

func (result *Matrix4) MakeTranslation(translateVec *Vector3) {
	//x-axis
	result[m4col0+x] = 1.0
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0
	//y-axis
	result[m4col1+x] = 0.0
	result[m4col1+y] = 1.0
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0
	//z-axis
	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = 1.0
	result[m4col2+w] = 0.0

	result[m4col3+x] = translateVec[x]
	result[m4col3+y] = translateVec[y]
	result[m4col3+z] = translateVec[z]
	result[m4col3+w] = 1.0
}

func (result *Matrix4) MakeLookAt(eyePos, lookAtPos *Point3, upVec *Vector3) {
	var m4EyeFrame Matrix4
	var v3X, v3Y, v3Z, tmpV3_0, tmpV3_1 Vector3
	var tmpV4_0, tmpV4_1, tmpV4_2, tmpV4_3 Vector4

	v3Y.Normalize(upVec)
	tmpV3_0.P3Sub(eyePos, lookAtPos)
	v3Z.Normalize(&tmpV3_0)
	tmpV3_1.Cross(&v3Y, &v3Z)
	v3X.Normalize(&tmpV3_1)
	v3Y.Cross(&v3Z, &v3X)
	tmpV4_0.MakeFromV3(&v3X)
	tmpV4_1.MakeFromV3(&v3Y)
	tmpV4_2.MakeFromV3(&v3Z)
	tmpV4_3.MakeFromP3(eyePos)
	m4EyeFrame.MakeFromCols(&tmpV4_0, &tmpV4_1, &tmpV4_2, &tmpV4_3)
	result.OrthoInverse(&m4EyeFrame)
}

func (result *Matrix4) MakePerspective(fovyRadians, aspect, zNear, zFar float32) {
	f := tan(g_PI_OVER_2 - (0.5 * fovyRadians))
	rangeInv := 1.0 / (zNear - zFar)

	result[m4col0+x] = (f / aspect)
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	result[m4col1+x] = 0.0
	result[m4col1+y] = f
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = ((zNear + zFar) * rangeInv)
	result[m4col2+w] = -1.0

	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = (((zNear * zFar) * rangeInv) * 2.0)
	result[m4col3+w] = 0.0
}

func (result *Matrix4) MakeFrustum(left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))
	n2 := (zNear + zNear)

	result[m4col0+x] = (n2 * inv_rl)
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	result[m4col1+x] = 0.0
	result[m4col1+y] = (n2 * inv_tb)
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	result[m4col2+x] = (sum_rl * inv_rl)
	result[m4col2+y] = (sum_tb * inv_tb)
	result[m4col2+z] = (sum_nf * inv_nf)
	result[m4col2+w] = -1.0

	result[m4col3+x] = 0.0
	result[m4col3+y] = 0.0
	result[m4col3+z] = ((n2 * inv_nf) * zFar)
	result[m4col3+w] = 0.0
}

func (result *Matrix4) MakeOrthographic(left, right, bottom, top, zNear, zFar float32) {
	sum_rl := (right + left)
	sum_tb := (top + bottom)
	sum_nf := (zNear + zFar)
	inv_rl := (1.0 / (right - left))
	inv_tb := (1.0 / (top - bottom))
	inv_nf := (1.0 / (zNear - zFar))

	//V4MakeFromElems(&result.Col0, (inv_rl + inv_rl), 0.0, 0.0, 0.0)
	result[m4col0+x] = (inv_rl + inv_rl)
	result[m4col0+y] = 0.0
	result[m4col0+z] = 0.0
	result[m4col0+w] = 0.0

	//V4MakeFromElems(&result.Col1, 0.0, (inv_tb + inv_tb), 0.0, 0.0)
	result[m4col1+x] = 0.0
	result[m4col1+y] = (inv_tb + inv_tb)
	result[m4col1+z] = 0.0
	result[m4col1+w] = 0.0

	//V4MakeFromElems(&result.Col2, 0.0, 0.0, (inv_nf + inv_nf), 0.0)
	result[m4col2+x] = 0.0
	result[m4col2+y] = 0.0
	result[m4col2+z] = (inv_nf + inv_nf)
	result[m4col2+w] = 0.0

	//V4MakeFromElems(&result.Col3, (-sum_rl * inv_rl), (-sum_tb * inv_tb), (sum_nf * inv_nf), 1.0)
	result[m4col3+x] = (-sum_rl * inv_rl)
	result[m4col3+y] = (-sum_tb * inv_tb)
	result[m4col3+z] = (sum_nf * inv_nf)
	result[m4col3+w] = 1.0
}

func (result *Matrix4) Select(mat0, mat1 *Matrix4, select1 int) {
	if select1 != 0 {
		result[m4col0+x] = mat1[m4col0+x]
		result[m4col0+y] = mat1[m4col0+y]
		result[m4col0+z] = mat1[m4col0+z]
		result[m4col0+w] = mat1[m4col0+w]

		result[m4col1+x] = mat1[m4col1+x]
		result[m4col1+y] = mat1[m4col1+y]
		result[m4col1+z] = mat1[m4col1+z]
		result[m4col1+w] = mat1[m4col1+w]

		result[m4col2+x] = mat1[m4col2+x]
		result[m4col2+y] = mat1[m4col2+y]
		result[m4col2+z] = mat1[m4col2+z]
		result[m4col2+w] = mat1[m4col2+w]

		result[m4col3+x] = mat1[m4col3+x]
		result[m4col3+y] = mat1[m4col3+y]
		result[m4col3+z] = mat1[m4col3+z]
		result[m4col3+w] = mat1[m4col3+w]

	} else {
		result[m4col0+x] = mat0[m4col0+x]
		result[m4col0+y] = mat0[m4col0+y]
		result[m4col0+z] = mat0[m4col0+z]
		result[m4col0+w] = mat0[m4col0+w]

		result[m4col1+x] = mat0[m4col1+x]
		result[m4col1+y] = mat0[m4col1+y]
		result[m4col1+z] = mat0[m4col1+z]
		result[m4col1+w] = mat0[m4col1+w]

		result[m4col2+x] = mat0[m4col2+x]
		result[m4col2+y] = mat0[m4col2+y]
		result[m4col2+z] = mat0[m4col2+z]
		result[m4col2+w] = mat0[m4col2+w]

		result[m4col3+x] = mat0[m4col3+x]
		result[m4col3+y] = mat0[m4col3+y]
		result[m4col3+z] = mat0[m4col3+z]
		result[m4col3+w] = mat0[m4col3+w]

	}
}

//Transform3
const (
	t3col0 = 0
	t3col1 = 3
	t3col2 = 6
	t3col3 = 9
)

func (result *Transform3) MakeFromScalar(scalar float32) {
	result[t3col0+x] = scalar
	result[t3col0+y] = scalar
	result[t3col0+z] = scalar

	result[t3col1+x] = scalar
	result[t3col1+y] = scalar
	result[t3col1+z] = scalar

	result[t3col2+x] = scalar
	result[t3col2+y] = scalar
	result[t3col2+z] = scalar

	result[t3col3+x] = scalar
	result[t3col3+y] = scalar
	result[t3col3+z] = scalar
}

func (result *Transform3) MakeFromCols(col0, col1, col2, col3 *Vector3) {
	result.SetCol(0, col0)
	result.SetCol(1, col1)
	result.SetCol(2, col2)
	result.SetCol(3, col3)
}

func (result *Transform3) MakeFromM3V3(tfrm *Matrix3, translateVec *Vector3) {
	result.SetUpper3x3(tfrm)
	result.SetTranslation(translateVec)
}

func (result *Transform3) MakeFromQV3(unitQuat *Quaternion, translateVec *Vector3) {
	var tmpM3_0 Matrix3
	tmpM3_0.MakeFromQ(unitQuat)
	result.SetUpper3x3(&tmpM3_0)
	result.SetTranslation(translateVec)
}

func (t *Transform3) SetCol(col int, vec *Vector3) {
	switch col {
	case 0:
		t[t3col0+x] = vec[x]
		t[t3col0+y] = vec[y]
		t[t3col0+z] = vec[z]
	case 1:
		t[t3col1+x] = vec[x]
		t[t3col1+y] = vec[y]
		t[t3col1+z] = vec[z]
	case 2:
		t[t3col2+x] = vec[x]
		t[t3col2+y] = vec[y]
		t[t3col2+z] = vec[z]
	case 3:
		t[t3col3+x] = vec[x]
		t[t3col3+y] = vec[y]
		t[t3col3+z] = vec[z]
	}
}

func (t *Transform3) SetRow(row int, vec *Vector4) {
	t[t3col0+row] = vec[x]
	t[t3col1+row] = vec[y]
	t[t3col2+row] = vec[z]
}

func (t *Transform3) SetElem(col, row int, val float32) {
	t[col*4+row] = val
}

func (t *Transform3) Elem(col, row int) float32 {
	return t[col*4+row]
}

func (t *Transform3) Col(result *Vector3, col int) {
	switch col {
	case 0:
		copy(result[:], t[t3col0:t3col1-1])
	case 1:
		copy(result[:], t[t3col1:t3col2-1])
	case 2:
		copy(result[:], t[t3col2:t3col3-1])
	case 3:
		copy(result[:], t[t3col3:])

	}
}

func (t *Transform3) Row(result *Vector4, row int) {
	result[x] = t[t3col0+row]
	result[y] = t[t3col1+row]
	result[z] = t[t3col2+row]
	result[w] = t[t3col3+row]
}

func (result *Transform3) Inverse(tfrm *Transform3) {
	var tmp0, tmp1, tmp2, tmpV3_3, tmpV3_4, tmpV3_5 Vector3
	var tfrmCol2 Vector3

	tmp0[x] = tfrm[t3col1+y]*tfrm[t3col2+z] - tfrm[t3col1+z]*tfrm[t3col2+y]
	tmp0[y] = tfrm[t3col1+z]*tfrm[t3col2+x] - tfrm[t3col1+x]*tfrm[t3col2+z]
	tmp0[z] = tfrm[t3col1+x]*tfrm[t3col2+y] - tfrm[t3col1+y]*tfrm[t3col2+x]

	tmp1[x] = tfrm[t3col2+y]*tfrm[t3col0+z] - tfrm[t3col2+z]*tfrm[t3col0+y]
	tmp1[y] = tfrm[t3col2+z]*tfrm[t3col0+x] - tfrm[t3col2+x]*tfrm[t3col0+z]
	tmp1[z] = tfrm[t3col2+x]*tfrm[t3col0+y] - tfrm[t3col2+y]*tfrm[t3col0+x]

	tmp2[x] = tfrm[t3col0+y]*tfrm[t3col1+z] - tfrm[t3col0+z]*tfrm[t3col1+y]
	tmp2[y] = tfrm[t3col0+z]*tfrm[t3col1+x] - tfrm[t3col0+x]*tfrm[t3col1+z]
	tmp2[z] = tfrm[t3col0+x]*tfrm[t3col1+y] - tfrm[t3col0+y]*tfrm[t3col1+x]

	tfrm.Col(&tfrmCol2, 2)

	detinv := (1.0 / tfrmCol2.Dot(&tmp2))

	result[t3col0+x] = (tmp0[x] * detinv)
	result[t3col0+y] = (tmp1[x] * detinv)
	result[t3col0+z] = (tmp2[x] * detinv)

	result[t3col1+x] = (tmp0[y] * detinv)
	result[t3col1+y] = (tmp1[y] * detinv)
	result[t3col1+z] = (tmp2[y] * detinv)

	result[t3col2+x] = (tmp0[z] * detinv)
	result[t3col2+y] = (tmp1[z] * detinv)
	result[t3col2+z] = (tmp2[z] * detinv)

	tmpV3_0 := Vector3{
		result[t3col0+x] * tfrm[t3col3+x],
		result[t3col0+y] * tfrm[t3col3+x],
		result[t3col0+z] * tfrm[t3col3+x]}

	tmpV3_1 := Vector3{
		result[t3col1+x] * tfrm[t3col3+y],
		result[t3col1+y] * tfrm[t3col3+y],
		result[t3col1+z] * tfrm[t3col3+y]}

	tmpV3_2 := Vector3{
		result[t3col2+x] * tfrm[t3col3+z],
		result[t3col2+y] * tfrm[t3col3+z],
		result[t3col2+z] * tfrm[t3col3+z]}

	tmpV3_3.Add(&tmpV3_1, &tmpV3_2)
	tmpV3_4.Add(&tmpV3_0, &tmpV3_3)
	tmpV3_5.Neg(&tmpV3_4)

	result[t3col3+x] = tmpV3_5[x]
	result[t3col3+y] = tmpV3_5[y]
	result[t3col3+z] = tmpV3_5[z]

}

func (t *Transform3) InverseSelf() {
	tmp := *t
	t.Inverse(&tmp)
}

func (result *Transform3) OrthoInverse(tfrm *Transform3) {
	var tmpV3_3, tmpV3_4, tmpV3_5 Vector3

	result[t3col0+x] = tfrm[t3col0+x]
	result[t3col0+y] = tfrm[t3col1+x]
	result[t3col0+z] = tfrm[t3col2+x]

	result[t3col1+x] = tfrm[t3col0+y]
	result[t3col1+y] = tfrm[t3col1+y]
	result[t3col1+z] = tfrm[t3col2+y]

	result[t3col2+x] = tfrm[t3col0+z]
	result[t3col2+y] = tfrm[t3col1+z]
	result[t3col2+z] = tfrm[t3col2+z]

	tmpV3_0 := Vector3{
		result[t3col0+x] * tfrm[t3col3+x],
		result[t3col0+y] * tfrm[t3col3+x],
		result[t3col0+z] * tfrm[t3col3+x]}

	tmpV3_1 := Vector3{
		result[t3col1+x] * tfrm[t3col3+y],
		result[t3col1+y] * tfrm[t3col3+y],
		result[t3col1+z] * tfrm[t3col3+y]}

	tmpV3_2 := Vector3{
		result[t3col2+x] * tfrm[t3col3+z],
		result[t3col2+y] * tfrm[t3col3+z],
		result[t3col2+z] * tfrm[t3col3+z]}

	tmpV3_3.Add(&tmpV3_1, &tmpV3_2)
	tmpV3_4.Add(&tmpV3_0, &tmpV3_3)
	tmpV3_5.Neg(&tmpV3_4)

	result[t3col3+x] = tmpV3_5[x]
	result[t3col3+y] = tmpV3_5[y]
	result[t3col3+z] = tmpV3_5[z]
}

func (result *Transform3) OrthoInverseSelf() {
	tmp := *result
	result.OrthoInverse(&tmp)
}

func (result *Transform3) AbsPerElem(tfrm *Transform3) {
	result[t3col0+x] = abs(tfrm[t3col0+x])
	result[t3col0+y] = abs(tfrm[t3col0+y])
	result[t3col0+z] = abs(tfrm[t3col0+z])

	result[t3col1+x] = abs(tfrm[t3col1+x])
	result[t3col1+y] = abs(tfrm[t3col1+y])
	result[t3col1+z] = abs(tfrm[t3col1+z])

	result[t3col2+x] = abs(tfrm[t3col2+x])
	result[t3col2+y] = abs(tfrm[t3col2+y])
	result[t3col2+z] = abs(tfrm[t3col2+z])

	result[t3col3+x] = abs(tfrm[t3col3+x])
	result[t3col3+y] = abs(tfrm[t3col3+y])
	result[t3col3+z] = abs(tfrm[t3col3+z])
}

func (result *Vector3) MulT3(tfrm *Transform3, vec *Vector3) {
	result[x] = ((tfrm[t3col0+x] * vec[x]) + (tfrm[t3col1+x] * vec[y])) + (tfrm[t3col2+x] * vec[z])
	result[y] = ((tfrm[t3col0+y] * vec[x]) + (tfrm[t3col1+y] * vec[y])) + (tfrm[t3col2+y] * vec[z])
	result[z] = ((tfrm[t3col0+z] * vec[x]) + (tfrm[t3col1+z] * vec[y])) + (tfrm[t3col2+z] * vec[z])
}

func (result *Vector3) MulT3Self(tfrm *Transform3) {
	tmp := *result
	result.MulT3(tfrm, &tmp)
}

func (result *Point3) MulT3(tfrm *Transform3, pnt *Point3) {
	result[x] = ((((tfrm[t3col0+x] * pnt[x]) + (tfrm[t3col1+x] * pnt[y])) + (tfrm[t3col2+x] * pnt[z])) + tfrm[t3col3+x])
	result[y] = ((((tfrm[t3col0+y] * pnt[x]) + (tfrm[t3col1+y] * pnt[y])) + (tfrm[t3col2+y] * pnt[z])) + tfrm[t3col3+y])
	result[z] = ((((tfrm[t3col0+z] * pnt[x]) + (tfrm[t3col1+z] * pnt[y])) + (tfrm[t3col2+z] * pnt[z])) + tfrm[t3col3+z])
}

func (result *Point3) MulT3Self(tfrm *Transform3) {
	tmp := *result

	result.MulT3(tfrm, &tmp)
}

func (result *Transform3) Mul(tfrm0, tfrm1 *Transform3) {

	result[t3col0+x] = ((tfrm0[t3col0+x] * tfrm1[t3col0+x]) + (tfrm0[t3col1+x] * tfrm1[t3col0+y])) + (tfrm0[t3col2+x] * tfrm1[t3col0+z])
	result[t3col0+y] = ((tfrm0[t3col0+y] * tfrm1[t3col0+x]) + (tfrm0[t3col1+y] * tfrm1[t3col0+y])) + (tfrm0[t3col2+y] * tfrm1[t3col0+z])
	result[t3col0+z] = ((tfrm0[t3col0+z] * tfrm1[t3col0+x]) + (tfrm0[t3col1+z] * tfrm1[t3col0+y])) + (tfrm0[t3col2+z] * tfrm1[t3col0+z])

	result[t3col1+x] = ((tfrm0[t3col0+x] * tfrm1[t3col1+x]) + (tfrm0[t3col1+x] * tfrm1[t3col1+y])) + (tfrm0[t3col2+x] * tfrm1[t3col1+z])
	result[t3col1+y] = ((tfrm0[t3col0+y] * tfrm1[t3col1+x]) + (tfrm0[t3col1+y] * tfrm1[t3col1+y])) + (tfrm0[t3col2+y] * tfrm1[t3col1+z])
	result[t3col1+z] = ((tfrm0[t3col0+z] * tfrm1[t3col1+x]) + (tfrm0[t3col1+z] * tfrm1[t3col1+y])) + (tfrm0[t3col2+z] * tfrm1[t3col1+z])

	result[t3col2+x] = ((tfrm0[t3col0+x] * tfrm1[t3col2+x]) + (tfrm0[t3col1+x] * tfrm1[t3col2+y])) + (tfrm0[t3col2+x] * tfrm1[t3col2+z])
	result[t3col2+y] = ((tfrm0[t3col0+y] * tfrm1[t3col2+x]) + (tfrm0[t3col1+y] * tfrm1[t3col2+y])) + (tfrm0[t3col2+y] * tfrm1[t3col2+z])
	result[t3col2+z] = ((tfrm0[t3col0+z] * tfrm1[t3col2+x]) + (tfrm0[t3col1+z] * tfrm1[t3col2+y])) + (tfrm0[t3col2+z] * tfrm1[t3col2+z])

	result[t3col3+x] = ((((tfrm0[t3col0+x] * tfrm1[t3col3+x]) + (tfrm0[t3col1+x] * tfrm1[t3col3+y])) + (tfrm0[t3col2+x] * tfrm1[t3col3+z])) + tfrm0[t3col3+x])
	result[t3col3+y] = ((((tfrm0[t3col0+y] * tfrm1[t3col3+x]) + (tfrm0[t3col1+y] * tfrm1[t3col3+y])) + (tfrm0[t3col2+y] * tfrm1[t3col3+z])) + tfrm0[t3col3+y])
	result[t3col3+z] = ((((tfrm0[t3col0+z] * tfrm1[t3col3+x]) + (tfrm0[t3col1+z] * tfrm1[t3col3+y])) + (tfrm0[t3col2+z] * tfrm1[t3col3+z])) + tfrm0[t3col3+z])

}

func (result *Transform3) MulSelf(tfrm *Transform3) {
	tmp := *result
	result.Mul(&tmp, tfrm)
}

func (result *Transform3) MulPerElem(tfrm0, tfrm1 *Transform3) {
	result[t3col0+x] = tfrm0[t3col0+x] * tfrm1[t3col0+x]
	result[t3col0+y] = tfrm0[t3col0+y] * tfrm1[t3col0+y]
	result[t3col0+z] = tfrm0[t3col0+z] * tfrm1[t3col0+z]

	result[t3col1+x] = tfrm0[t3col1+x] * tfrm1[t3col1+x]
	result[t3col1+y] = tfrm0[t3col1+y] * tfrm1[t3col1+y]
	result[t3col1+z] = tfrm0[t3col1+z] * tfrm1[t3col1+z]

	result[t3col2+x] = tfrm0[t3col2+x] * tfrm1[t3col2+x]
	result[t3col2+y] = tfrm0[t3col2+y] * tfrm1[t3col2+y]
	result[t3col2+z] = tfrm0[t3col2+z] * tfrm1[t3col2+z]

	result[t3col3+x] = tfrm0[t3col3+x] * tfrm1[t3col3+x]
	result[t3col3+y] = tfrm0[t3col3+y] * tfrm1[t3col3+y]
	result[t3col3+z] = tfrm0[t3col3+z] * tfrm1[t3col3+z]
}

func (result *Transform3) MulPerElemSelf(tfrm *Transform3) {
	result.MulPerElem(result, tfrm)
}

func (result *Transform3) MakeIdentity() {
	//x-axis
	result[t3col0+x] = 1.0
	result[t3col0+y] = 0.0
	result[t3col0+z] = 0.0

	//y-axis
	result[t3col1+x] = 0.0
	result[t3col1+y] = 1.0
	result[t3col1+z] = 0.0

	//z-axis
	result[t3col2+x] = 0.0
	result[t3col2+y] = 0.0
	result[t3col2+z] = 1.0

	//w-axis
	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0

}

func (t *Transform3) SetUpper3x3(m *Matrix3) {
	t[t3col0+x] = m[m3col0+x]
	t[t3col0+y] = m[m3col0+y]
	t[t3col0+z] = m[m3col0+z]

	t[t3col1+x] = m[m3col1+x]
	t[t3col1+y] = m[m3col1+y]
	t[t3col1+z] = m[m3col1+z]

	t[t3col2+x] = m[m3col2+x]
	t[t3col2+y] = m[m3col2+y]
	t[t3col2+z] = m[m3col2+z]
}

func (t *Transform3) Upper3x3(result *Matrix3) {
	result[m3col0+x] = t[t3col0+x]
	result[m3col0+y] = t[t3col0+y]
	result[m3col0+z] = t[t3col0+z]

	result[m3col1+x] = t[t3col1+x]
	result[m3col1+y] = t[t3col1+y]
	result[m3col1+z] = t[t3col1+z]

	result[m3col2+x] = t[t3col2+x]
	result[m3col2+y] = t[t3col2+y]
	result[m3col2+z] = t[t3col2+z]
}

func (t *Transform3) SetTranslation(translateVec *Vector3) {
	t[t3col3+x] = translateVec[x]
	t[t3col3+y] = translateVec[y]
	t[t3col3+z] = translateVec[z]
}

func (tfrm *Transform3) Translation(result *Vector3) {
	result[x] = tfrm[t3col3+x]
	result[y] = tfrm[t3col3+y]
	result[z] = tfrm[t3col3+z]
}

func (result *Transform3) MakeRotationX(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[t3col0+x] = 1.0
	result[t3col0+y] = 0.0
	result[t3col0+z] = 0.0

	result[t3col1+x] = 0.0
	result[t3col1+y] = c
	result[t3col1+z] = s

	result[t3col1+x] = 0.0
	result[t3col1+y] = -s
	result[t3col1+z] = c

	result[t3col2+x] = 0
	result[t3col2+y] = 0
	result[t3col2+z] = 0

}

func (result *Transform3) MakeRotationY(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[t3col0+x] = c
	result[t3col0+y] = 0.0
	result[t3col0+z] = -s

	//y-axis
	result[t3col1+x] = 0.0
	result[t3col1+y] = 1.0
	result[t3col1+z] = 0.0

	result[t3col2+x] = s
	result[t3col2+y] = 0.0
	result[t3col2+z] = c

	//w-axis
	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0

}

func (result *Transform3) MakeRotationZ(radians float32) {
	s := sin(radians)
	c := cos(radians)

	result[t3col0+x] = c
	result[t3col0+y] = s
	result[t3col0+z] = 0.0

	result[t3col1+x] = -s
	result[t3col1+y] = c
	result[t3col1+z] = 0.0

	//z-axis
	result[t3col2+x] = 0.0
	result[t3col2+y] = 0.0
	result[t3col2+z] = 1.0

	//w-axis
	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0

}

func (result *Transform3) MakeRotationZYX(radiansXYZ *Vector3) {
	sX := sin(radiansXYZ[x])
	cX := cos(radiansXYZ[x])
	sY := sin(radiansXYZ[y])
	cY := cos(radiansXYZ[y])
	sZ := sin(radiansXYZ[z])
	cZ := cos(radiansXYZ[z])
	tmp0 := (cZ * sY)
	tmp1 := (sZ * sY)

	result[t3col0+x] = (cZ * cY)
	result[t3col0+y] = (sZ * cY)
	result[t3col0+z] = -sY

	result[t3col1+x] = ((tmp0 * sX) - (sZ * cX))
	result[t3col1+y] = ((tmp1 * sX) + (cZ * cX))
	result[t3col1+z] = (cY * sX)

	result[t3col2+x] = ((tmp0 * cX) + (sZ * sX))
	result[t3col2+y] = ((tmp1 * cX) - (cZ * sX))
	result[t3col2+z] = (cY * cX)

	//w-axis
	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0

}

func (result *Transform3) MakeRotationAxis(radians float32, unitVec *Vector3) {
	s := sin(radians)
	c := cos(radians)
	X := unitVec[x]
	Y := unitVec[y]
	Z := unitVec[z]
	xy := X * Y
	yz := Y * Z
	zx := Z * X
	oneMinusC := 1.0 - c

	result[t3col0+x] = (((X * X) * oneMinusC) + c)
	result[t3col0+y] = ((xy * oneMinusC) + (Z * s))
	result[t3col0+z] = ((zx * oneMinusC) - (Y * s))

	result[t3col1+x] = ((xy * oneMinusC) - (Z * s))
	result[t3col1+y] = (((Y * Y) * oneMinusC) + c)
	result[t3col1+z] = ((yz * oneMinusC) + (X * s))

	result[t3col2+x] = ((zx * oneMinusC) + (Y * s))
	result[t3col2+y] = ((yz * oneMinusC) - (X * s))
	result[t3col2+z] = (((Z * Z) * oneMinusC) + c)

	//w-axis
	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0

}

func (result *Transform3) MakeRotationQ(unitQuat *Quaternion) {
	var tmpM3 Matrix3

	tmpM3.MakeFromQ(unitQuat)
	result.MakeFromM3V3(&tmpM3, &Vector3{0, 0, 0})
}

func (result *Transform3) MakeScale(scaleVec *Vector3) {
	result[t3col0+x] = scaleVec[x]
	result[t3col0+y] = 0.0
	result[t3col0+z] = 0.0

	result[t3col1+x] = 0.0
	result[t3col1+y] = scaleVec[y]
	result[t3col1+z] = 0.0

	result[t3col2+x] = 0.0
	result[t3col2+y] = 0.0
	result[t3col2+z] = scaleVec[z]

	result[t3col3+x] = 0.0
	result[t3col3+y] = 0.0
	result[t3col3+z] = 0.0
}

func (result *Transform3) AppendScale(tfrm *Transform3, scaleVec *Vector3) {
	result[t3col0+x] = tfrm[t3col0+x] * scaleVec[x]
	result[t3col0+y] = tfrm[t3col0+y] * scaleVec[x]
	result[t3col0+z] = tfrm[t3col0+z] * scaleVec[x]

	result[t3col1+x] = tfrm[t3col1+x] * scaleVec[y]
	result[t3col1+y] = tfrm[t3col1+y] * scaleVec[y]
	result[t3col1+z] = tfrm[t3col1+z] * scaleVec[y]

	result[t3col2+x] = tfrm[t3col2+x] * scaleVec[z]
	result[t3col2+y] = tfrm[t3col2+y] * scaleVec[z]
	result[t3col2+z] = tfrm[t3col2+z] * scaleVec[z]

	result[t3col3+x] = tfrm[t3col3+x]
	result[t3col3+y] = tfrm[t3col3+y]
	result[t3col3+z] = tfrm[t3col3+z]

}

func (result *Transform3) AppendScaleSelf(scaleVec *Vector3) {
	result.AppendScale(result, scaleVec)
}

func (result *Transform3) PrependScale(scaleVec *Vector3, tfrm *Transform3) {
	result[t3col0+x] = tfrm[t3col0+x] * scaleVec[x]
	result[t3col0+y] = tfrm[t3col0+y] * scaleVec[y]
	result[t3col0+z] = tfrm[t3col0+z] * scaleVec[z]

	result[t3col1+x] = tfrm[t3col1+x] * scaleVec[x]
	result[t3col1+y] = tfrm[t3col1+y] * scaleVec[y]
	result[t3col1+z] = tfrm[t3col1+z] * scaleVec[z]

	result[t3col2+x] = tfrm[t3col2+x] * scaleVec[x]
	result[t3col2+y] = tfrm[t3col2+y] * scaleVec[y]
	result[t3col2+z] = tfrm[t3col2+z] * scaleVec[z]

	result[t3col3+x] = tfrm[t3col3+x] * scaleVec[x]
	result[t3col3+y] = tfrm[t3col3+y] * scaleVec[y]
	result[t3col3+z] = tfrm[t3col3+z] * scaleVec[z]
}

func (result *Transform3) PrependScaleSelf(scaleVec *Vector3) {
	result.PrependScale(scaleVec, result)
}

func (result *Transform3) MakeTranslation(translateVec *Vector3) {
	//x-axis
	result[t3col0+x] = 1.0
	result[t3col0+y] = 0.0
	result[t3col0+z] = 0.0
	//y-axis
	result[t3col1+x] = 0.0
	result[t3col1+y] = 1.0
	result[t3col1+z] = 0.0
	//z-axis
	result[t3col2+x] = 0.0
	result[t3col2+y] = 0.0
	result[t3col2+z] = 1.0

	result[t3col3+x] = translateVec[x]
	result[t3col3+y] = translateVec[y]
	result[t3col3+z] = translateVec[z]
}

func (result *Transform3) Select(tfrm0, tfrm1 *Transform3, select1 int) {
	if select1 != 0 {
		result[t3col0+x] = tfrm1[t3col0+x]
		result[t3col0+y] = tfrm1[t3col0+y]
		result[t3col0+z] = tfrm1[t3col0+z]

		result[t3col1+x] = tfrm1[t3col1+x]
		result[t3col1+y] = tfrm1[t3col1+y]
		result[t3col1+z] = tfrm1[t3col1+z]

		result[t3col2+x] = tfrm1[t3col2+x]
		result[t3col2+y] = tfrm1[t3col2+y]
		result[t3col2+z] = tfrm1[t3col2+z]

		result[t3col3+x] = tfrm1[t3col3+x]
		result[t3col3+y] = tfrm1[t3col3+y]
		result[t3col3+z] = tfrm1[t3col3+z]

	} else {
		result[t3col0+x] = tfrm0[t3col0+x]
		result[t3col0+y] = tfrm0[t3col0+y]
		result[t3col0+z] = tfrm0[t3col0+z]

		result[t3col1+x] = tfrm0[t3col1+x]
		result[t3col1+y] = tfrm0[t3col1+y]
		result[t3col1+z] = tfrm0[t3col1+z]

		result[t3col2+x] = tfrm0[t3col2+x]
		result[t3col2+y] = tfrm0[t3col2+y]
		result[t3col2+z] = tfrm0[t3col2+z]

		result[t3col3+x] = tfrm0[t3col3+x]
		result[t3col3+y] = tfrm0[t3col3+y]
		result[t3col3+z] = tfrm0[t3col3+z]

	}
}

func (result *Matrix3) V3Outer(tfrm0, tfrm1 *Vector3) {
	result[m3col0+x] = tfrm0[x] * tfrm1[x]
	result[m3col0+y] = tfrm0[y] * tfrm1[x]
	result[m3col0+z] = tfrm0[z] * tfrm1[x]

	result[m3col1+x] = tfrm0[x] * tfrm1[y]
	result[m3col1+y] = tfrm0[y] * tfrm1[y]
	result[m3col1+z] = tfrm0[z] * tfrm1[y]

	result[m3col2+x] = tfrm0[x] * tfrm1[z]
	result[m3col2+y] = tfrm0[y] * tfrm1[z]
	result[m3col2+z] = tfrm0[z] * tfrm1[z]

}

func (result *Matrix4) V4Outer(tfrm0, tfrm1 *Vector4) {
	result[m4col0+x] = tfrm0[x] * tfrm1[x]
	result[m4col0+y] = tfrm0[y] * tfrm1[x]
	result[m4col0+z] = tfrm0[z] * tfrm1[x]
	result[m4col0+w] = tfrm0[w] * tfrm1[x]

	result[m4col1+x] = tfrm0[x] * tfrm1[y]
	result[m4col1+y] = tfrm0[y] * tfrm1[y]
	result[m4col1+z] = tfrm0[z] * tfrm1[y]
	result[m4col1+w] = tfrm0[w] * tfrm1[y]

	result[m4col2+x] = tfrm0[x] * tfrm1[z]
	result[m4col2+y] = tfrm0[y] * tfrm1[z]
	result[m4col2+z] = tfrm0[z] * tfrm1[z]
	result[m4col2+w] = tfrm0[w] * tfrm1[z]

	result[m4col3+x] = tfrm0[x] * tfrm1[z]
	result[m4col3+y] = tfrm0[y] * tfrm1[z]
	result[m4col3+z] = tfrm0[z] * tfrm1[z]
	result[m4col3+w] = tfrm0[w] * tfrm1[z]
}

func (result *Vector3) RowMulMat3(vec *Vector3, mat *Matrix3) {
	result[x] = (((vec[x] * mat[m3col0+x]) + (vec[y] * mat[m3col0+y])) + (vec[z] * mat[m3col0+z]))
	result[y] = (((vec[x] * mat[m3col1+x]) + (vec[y] * mat[m3col1+y])) + (vec[z] * mat[m3col1+z]))
	result[z] = (((vec[x] * mat[m3col2+x]) + (vec[y] * mat[m3col2+y])) + (vec[z] * mat[m3col2+z]))
}

func (result *Vector3) RowMulMat3Self(mat *Matrix3) {
	tmp := *result
	result.RowMulMat3(&tmp, mat)
}

func (result *Matrix3) V3CrossMatrix(vec *Vector3) {
	result[m3col0+x] = 0.0
	result[m3col0+y] = vec[z]
	result[m3col0+z] = -vec[y]

	result[m3col1+x] = -vec[z]
	result[m3col1+y] = 0.0
	result[m3col1+z] = vec[x]

	result[m3col2+x] = vec[y]
	result[m3col2+y] = -vec[x]
	result[m3col2+z] = 0.0
}

func (result *Matrix3) V3CrossMatrixMul(vec *Vector3, mat *Matrix3) {
	result[m3col0+x] = vec[y]*mat[m3col0+z] - vec[z]*mat[m3col0+y]
	result[m3col0+y] = vec[z]*mat[m3col0+x] - vec[x]*mat[m3col0+z]
	result[m3col0+z] = vec[x]*mat[m3col0+y] - vec[y]*mat[m3col0+x]

	result[m3col1+x] = vec[y]*mat[m3col1+z] - vec[z]*mat[m3col1+y]
	result[m3col1+y] = vec[z]*mat[m3col1+x] - vec[x]*mat[m3col1+z]
	result[m3col1+z] = vec[x]*mat[m3col1+y] - vec[y]*mat[m3col1+x]

	result[m3col2+x] = vec[y]*mat[m3col2+z] - vec[z]*mat[m3col2+y]
	result[m3col2+y] = vec[z]*mat[m3col2+x] - vec[x]*mat[m3col2+z]
	result[m3col2+z] = vec[x]*mat[m3col2+y] - vec[y]*mat[m3col2+x]
}

func (result *Matrix3) V3CrossMatrixMulSelf(vec *Vector3) {
	tmp := *result
	result.V3CrossMatrixMul(vec, &tmp)
}
