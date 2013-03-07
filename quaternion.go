//Copyright (C) 2006, 2007 Sony Computer Entertainment Inc.
//  All rights reserved.
// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

import (
	"unsafe"
)

func (result *Quaternion) MakeFromM3(tfrm *Matrix3) {
	xx := tfrm[t3col0+x]
	yx := tfrm[t3col0+y]
	zx := tfrm[t3col0+z]
	xy := tfrm[t3col1+x]
	yy := tfrm[t3col1+y]
	zy := tfrm[t3col1+z]
	xz := tfrm[t3col2+x]
	yz := tfrm[t3col2+y]
	zz := tfrm[t3col2+z]

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

	result[x] = qx
	result[y] = qy
	result[z] = qz
	result[w] = qw
}

func (result *Quaternion) MakeFromV3Scalar(xyz *Vector3, W float32) {
	result[x] = xyz[x]
	result[y] = xyz[y]
	result[z] = xyz[z]
	result[w] = W
}

func (result *Quaternion) MakeFromV4(vec *Vector4) {
	result[x] = vec[x]
	result[y] = vec[y]
	result[z] = vec[z]
	result[w] = vec[w]
}

func (result *Quaternion) MakeFromScalar(scalar float32) {
	result[x] = scalar
	result[y] = scalar
	result[z] = scalar
	result[w] = scalar
}

func (result *Quaternion) MakeIdentity() {
	result[x] = 0.0
	result[y] = 0.0
	result[z] = 0.0
	result[w] = 1.0
}

func (result *Quaternion) Lerp(t float32, quat0, quat1 *Quaternion) {
	var tmpQ_0, tmpQ_1 Quaternion

	tmpQ_0.Sub(quat1, quat0)
	tmpQ_1.ScalarMul(&tmpQ_0, t)
	result.Add(quat0, &tmpQ_1)
}

func (result *Quaternion) LerpTo(t float32, quatTo *Quaternion) {
	tmp := *result
	result.Lerp(t, &tmp, quatTo)
}

func (result *Quaternion) Slerp(t float32, unitQuat0, unitQuat1 *Quaternion) {
	if unsafe.Pointer(result) == unsafe.Pointer(unitQuat0) {
		result.SlerpSelf(t, unitQuat1)
		return
	}
	var start, tmpQ_0, tmpQ_1 Quaternion
	var scale0, scale1 float32

	cosAngle := unitQuat0.Dot(unitQuat1)
	if cosAngle < 0.0 {
		cosAngle = -cosAngle
		start.Neg(unitQuat0)
	} else {
		copy(start[:], unitQuat0[:])
	}
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := (1.0 / sin(angle))
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = (1.0 - t)
		scale1 = t
	}
	tmpQ_0.ScalarMul(&start, scale0)
	tmpQ_1.ScalarMul(unitQuat1, scale1)
	result.Add(&tmpQ_0, &tmpQ_1)
}

func (result *Quaternion) SlerpSelf(t float32, unitQuatTo *Quaternion) {
	tmp := *result

	result.Slerp(t, &tmp, unitQuatTo)
}

func (result *Quaternion) Squad(t float32, unitQuat0, unitQuat1, unitQuat2, unitQuat3 *Quaternion) {
	var tmp0, tmp1 Quaternion
	tmp0.Slerp(t, unitQuat0, unitQuat3)
	tmp1.Slerp(t, unitQuat1, unitQuat2)
	result.Slerp((2.0*t)*(1.0-t), &tmp0, &tmp1)
}

func (q *Quaternion) SetXYZ(vec *Vector3) {
	q[x] = vec[x]
	q[y] = vec[y]
	q[z] = vec[z]
}

func (result *Quaternion) Add(quat0, quat1 *Quaternion) {
	result[x] = quat0[x] + quat1[x]
	result[y] = quat0[y] + quat1[y]
	result[z] = quat0[z] + quat1[z]
	result[w] = quat0[w] + quat1[w]
}

func (result *Quaternion) AddToSelf(quat *Quaternion) {
	result.Add(result, quat)
}

func (result *Quaternion) Sub(quat0, quat1 *Quaternion) {
	result[x] = quat0[x] - quat1[x]
	result[y] = quat0[y] - quat1[y]
	result[z] = quat0[z] - quat1[z]
	result[w] = quat0[w] - quat1[w]
}

func (result *Quaternion) SubFromSelf(quat *Quaternion) {
	result.Sub(result, quat)
}

func (result *Quaternion) ScalarMul(quat *Quaternion, scalar float32) {
	result[x] = quat[x] * scalar
	result[y] = quat[y] * scalar
	result[z] = quat[z] * scalar
	result[w] = quat[w] * scalar
}

func (result *Quaternion) ScalarMulSelf(scalar float32) {
	result.ScalarMul(result, scalar)
}

func (result *Quaternion) ScalarDiv(quat *Quaternion, scalar float32) {
	result[x] = quat[x] / scalar
	result[y] = quat[y] / scalar
	result[z] = quat[z] / scalar
	result[w] = quat[w] / scalar
}

func (result *Quaternion) ScalarDivSelf(scalar float32) {
	result.ScalarDiv(result, scalar)
}

func (result *Quaternion) Neg(quat *Quaternion) {
	result[x] = -quat[x]
	result[y] = -quat[y]
	result[z] = -quat[z]
	result[w] = -quat[w]
}

func (result *Quaternion) NegSelf() {
	result.Neg(result)
}

func (q *Quaternion) Dot(quat *Quaternion) float32 {
	result := q[x] * quat[x]
	result += q[y] * quat[y]
	result += q[z] * quat[z]
	result += q[w] * quat[w]
	return result
}

func (q *Quaternion) Norm() float32 {
	result := q[x] * q[x]
	result += q[y] * q[y]
	result += q[z] * q[z]
	result += q[w] * q[w]
	return result
}

func (q *Quaternion) Length() float32 {
	return sqrt(q.Norm())
}

func (result *Quaternion) Normalize(quat *Quaternion) {
	lenSqr := quat.Norm()
	lenInv := 1.0 / sqrt(lenSqr)
	result[x] = quat[x] * lenInv
	result[y] = quat[y] * lenInv
	result[z] = quat[z] * lenInv
	result[w] = quat[w] * lenInv
}

func (result *Quaternion) NormalizeSelf() {
	result.Normalize(result)

}

func (result *Quaternion) MakeRotationArc(unitVec0, unitVec1 *Vector3) {
	var tmpV3_0, tmpV3_1 Vector3
	cosHalfAngleX2 := sqrt((2.0 * (1.0 + unitVec0.Dot(unitVec1))))
	recipCosHalfAngleX2 := (1.0 / cosHalfAngleX2)
	tmpV3_0.Cross(unitVec0, unitVec1)

	tmpV3_1.ScalarMul(&tmpV3_0, recipCosHalfAngleX2)
	result.MakeFromV3Scalar(&tmpV3_1, (cosHalfAngleX2 * 0.5))
}

func (result *Quaternion) MakeRotationAxis(radians float32, unitVec *Vector3) {
	var tmpV3_0 Vector3
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	tmpV3_0.ScalarMul(unitVec, s)
	result.MakeFromV3Scalar(&tmpV3_0, c)
}

func (result *Quaternion) MakeRotationX(radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)
	result[x] = s
	result[y] = 0.0
	result[z] = 0.0
	result[w] = c
}

func (result *Quaternion) MakeRotationY(radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)

	result[x] = 0.0
	result[y] = s
	result[z] = 0.0
	result[w] = c

}

func (result *Quaternion) MakeRotationZ(radians float32) {
	angle := radians * 0.5
	s := sin(angle)
	c := cos(angle)

	result[x] = 0.0
	result[y] = 0.0
	result[z] = s
	result[w] = c
}

func (result *Quaternion) Mul(quat0, quat1 *Quaternion) {
	if unsafe.Pointer(result) == unsafe.Pointer(quat0) {
		result.MulSelf(quat1)
		return
	}

	if unsafe.Pointer(result) == unsafe.Pointer(quat1) {
		result.MulSelf(quat0)
		return
	}
	result[x] = (quat0[w] * quat1[x]) + (quat0[x] * quat1[w]) + (quat0[y] * quat1[z]) - (quat0[z] * quat1[y])
	result[y] = (quat0[w] * quat1[y]) + (quat0[y] * quat1[w]) + (quat0[z] * quat1[x]) - (quat0[x] * quat1[z])
	result[z] = (quat0[w] * quat1[z]) + (quat0[z] * quat1[w]) + (quat0[x] * quat1[y]) - (quat0[y] * quat1[x])
	result[w] = (quat0[w] * quat1[w]) - (quat0[x] * quat1[x]) - (quat0[y] * quat1[y]) - (quat0[z] * quat1[z])
}

func (result *Quaternion) MulSelf(quat *Quaternion) {
	tmp := *result
	result.Mul(&tmp, quat)

}

func (result *Vector3) Rotate(quat *Quaternion, vec *Vector3) {
	tmpX := (quat[w] * vec[x]) + (quat[y] * vec[z]) - (quat[z] * vec[y])
	tmpY := (quat[w] * vec[y]) + (quat[z] * vec[x]) - (quat[x] * vec[z])
	tmpZ := (quat[w] * vec[z]) + (quat[x] * vec[y]) - (quat[y] * vec[x])
	tmpW := (quat[x] * vec[x]) + (quat[y] * vec[y]) + (quat[z] * vec[z])
	result[x] = (tmpW * quat[x]) + (tmpX * quat[w]) - (tmpY * quat[z]) + (tmpZ * quat[y])
	result[y] = (tmpW * quat[y]) + (tmpY * quat[w]) - (tmpZ * quat[x]) + (tmpX * quat[z])
	result[z] = (tmpW * quat[z]) + (tmpZ * quat[w]) - (tmpX * quat[y]) + (tmpY * quat[x])
}

func (result *Vector3) RotateSelf(quat *Quaternion) {
	result.Rotate(quat, result)
}

func (result *Quaternion) Conj(quat *Quaternion) {
	result[x] = -quat[x]
	result[y] = -quat[y]
	result[z] = -quat[z]
	result[w] = quat[w]
}

func (result *Quaternion) ConjSelf() {
	result.Conj(result)
}

func (result *Quaternion) Select(quat0, quat1 *Quaternion, select1 int) {
	if select1 != 0 {
		result[x] = quat1[x]
		result[y] = quat1[y]
		result[z] = quat1[z]
		result[w] = quat1[w]
	} else {
		result[x] = quat0[x]
		result[y] = quat0[y]
		result[z] = quat0[z]
		result[w] = quat0[w]
	}
}
