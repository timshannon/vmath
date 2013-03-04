//Copyright (C) 2006, 2007 Sony Computer Entertainment Inc.
//  All rights reserved.
// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

const g_SLERP_TOL = 0.999

//Vector3
func (v *Vector3) MakeFromP3(pnt *Point3) {
	v[x] = pnt[x]
	v[y] = pnt[y]
	v[z] = pnt[z]
}

func (v *Vector3) MakeFromScalar(scalar float32) {
	v[x] = scalar
	v[y] = scalar
	v[z] = scalar
}

func (v *Vector3) MakeXAxis() {
	v[x] = 1.0
	v[y] = 0.0
	v[z] = 0.0
}

func (v *Vector3) MakeYAxis() {
	v[x] = 0.0
	v[y] = 1.0
	v[z] = 0.0

}

func (v *Vector3) MakeZAxis() {
	v[x] = 0.0
	v[y] = 0.0
	v[z] = 1.0

}

func (result *Vector3) Add(vec0, vec1 *Vector3) {
	result[x] = vec0[x] + vec1[x]
	result[y] = vec0[y] + vec1[y]
	result[z] = vec0[z] + vec1[z]
}

func (result *Vector3) AddToSelf(vec *Vector3) {
	result.Add(result, vec)
}

func (result *Vector3) Sub(vec0, vec1 *Vector3) {
	result[x] = vec0[x] - vec1[x]
	result[y] = vec0[y] - vec1[y]
	result[z] = vec0[z] - vec1[z]
}

func (result *Vector3) SubFromSelf(vec *Vector3) {
	result.Sub(result, vec)
}

func (result *Vector3) AddP3(vec0 *Vector3, pnt1 *Point3) {
	result[x] = vec0[x] + pnt1[x]
	result[y] = vec0[y] + pnt1[y]
	result[z] = vec0[z] + pnt1[z]
}

func (result *Vector3) AddP3ToSelf(pnt1 *Point3) {
	result.AddP3(result, pnt1)
}

func (result *Vector3) ScalarMul(vec *Vector3, scalar float32) {
	result[x] = vec[x] * scalar
	result[y] = vec[y] * scalar
	result[z] = vec[z] * scalar
}

func (result *Vector3) ScalarMulSelf(scalar float32) {
	result.ScalarMul(result, scalar)
}

func (result *Vector3) ScalarDiv(vec *Vector3, scalar float32) {
	result[x] = vec[x] / scalar
	result[y] = vec[y] / scalar
	result[z] = vec[z] / scalar
}

func (result *Vector3) ScalarDivSelf(scalar float32) {
	result.ScalarDiv(result, scalar)
}

func (result *Vector3) Neg(vec *Vector3) {
	result[x] = -vec[x]
	result[y] = -vec[y]
	result[z] = -vec[z]
}

func (result *Vector3) NegSelf() {
	result.Neg(result)
}

func (result *Vector3) MulPerElem(vec0, vec1 *Vector3) {
	result[x] = vec0[x] * vec1[x]
	result[y] = vec0[y] * vec1[y]
	result[z] = vec0[z] * vec1[z]
}

func (result *Vector3) MulPerElemSelf(vec *Vector3) {
	result.MulPerElem(result, vec)
}

func (result *Vector3) DivPerElem(vec0, vec1 *Vector3) {
	result[x] = vec0[x] / vec1[x]
	result[y] = vec0[y] / vec1[y]
	result[z] = vec0[z] / vec1[z]
}

func (result *Vector3) DivPerElemSelf(vec *Vector3) {
	result.DivPerElem(result, vec)
}

func (result *Vector3) RecipPerElem(vec *Vector3) {
	result[x] = 1.0 / vec[x]
	result[y] = 1.0 / vec[y]
	result[z] = 1.0 / vec[z]
}

func (result *Vector3) RecipPerElemSelf() {
	result.RecipPerElem(result)
}

func (result *Vector3) SqrtPerElem(vec *Vector3) {
	result[x] = sqrt(vec[x])
	result[y] = sqrt(vec[y])
	result[z] = sqrt(vec[z])
}

func (result *Vector3) SqrtPerElemSelf() {
	result.SqrtPerElem(result)
}

func (result *Vector3) RsqrtPerElem(vec *Vector3) {
	result[x] = 1.0 / sqrt(vec[x])
	result[y] = 1.0 / sqrt(vec[y])
	result[z] = 1.0 / sqrt(vec[z])
}

func (result *Vector3) RsqrtPerElemSelf() {
	result.RsqrtPerElem(result)
}

func (result *Vector3) AbsPerElem(vec *Vector3) {
	result[x] = abs(vec[x])
	result[y] = abs(vec[y])
	result[z] = abs(vec[z])
}

func (result *Vector3) AbsPerElemSelf() {
	result.AbsPerElem(result)
}

func (result *Vector3) CopySignPerElem(vec0, vec1 *Vector3) {
	if vec1[x] < 0.0 {
		result[x] = -abs(vec0[x])
	} else {
		result[x] = abs(vec0[x])
	}
	if vec1[y] < 0.0 {
		result[y] = -abs(vec0[y])
	} else {
		result[y] = abs(vec0[y])
	}
	if vec1[z] < 0.0 {
		result[z] = -abs(vec0[z])
	} else {
		result[z] = abs(vec0[z])
	}
}

func (result *Vector3) CopySignPerElemSelf(vec *Vector3) {
	result.CopySignPerElem(result, vec)
}

func (result *Vector3) MaxPerElem(vec0, vec1 *Vector3) {
	result[x] = max(vec0[x], vec1[x])
	result[y] = max(vec0[y], vec1[y])
	result[z] = max(vec0[z], vec1[z])
}

func (result *Vector3) MaxPerElemSelf(vec *Vector3) {
	result.MaxPerElem(result, vec)
}

func (v *Vector3) MaxElem() float32 {
	var result float32
	result = max(v[x], v[y])
	result = max(v[z], result)
	return result
}

func (result *Vector3) MinPerElem(vec0, vec1 *Vector3) {
	result[x] = min(vec0[x], vec1[x])
	result[y] = min(vec0[y], vec1[y])
	result[z] = min(vec0[z], vec1[z])
}

func (result *Vector3) MinPerElemSelf(vec *Vector3) {
	result.MinPerElem(result, vec)
}

func (v *Vector3) MinElem() float32 {
	var result float32
	result = min(v[x], v[y])
	result = min(v[z], result)
	return result
}

func (v *Vector3) Sum() float32 {
	var result float32
	result = v[x] + v[y] + v[z]
	return result
}

func (v *Vector3) Dot(vec1 *Vector3) float32 {
	result := v[x] * vec1[x]
	result += v[y] * vec1[y]
	result += v[z] * vec1[z]
	return result
}

func (v *Vector3) LengthSqr() float32 {
	result := v[x] * v[x]
	result += v[y] * v[y]
	result += v[z] * v[z]
	return result
}

func (v *Vector3) Length() float32 {
	return sqrt(v.LengthSqr())
}

func (result *Vector3) Normalize(v *Vector3) {
	lenSqr := v.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result[x] = v[x] * lenInv
	result[y] = v[y] * lenInv
	result[z] = v[z] * lenInv
}

func (result *Vector3) NormalizeSelf() {
	result.Normalize(result)
}

func (result *Vector3) Cross(vec0, vec1 *Vector3) {
	result[x] = vec0[y]*vec1[z] - vec0[z]*vec1[y]
	result[y] = vec0[z]*vec1[x] - vec0[x]*vec1[z]
	result[z] = vec0[x]*vec1[y] - vec0[y]*vec1[x]
}

func (result *Vector3) Select(vec0, vec1 *Vector3, select1 int) {
	if select1 != 0 {
		result[x] = vec1[x]
		result[y] = vec1[y]
		result[z] = vec1[z]
	} else {
		result[x] = vec0[x]
		result[y] = vec0[y]
		result[z] = vec0[z]
	}
}

func (result *Vector3) Velocity(start, end *Vector3, elapsedTime float32) {
	//change in position / elapsedTime
	result.Sub(start, end)
	result[x] = result[x] / elapsedTime
	result[y] = result[y] / elapsedTime
	result[z] = result[z] / elapsedTime
}

func (result *Vector3) Lerp(t float32, vec0, vec1 *Vector3) {
	result.Sub(vec1, vec0)
	result.ScalarMulSelf(t)
	result.Add(vec0, result)
}

func (result *Vector3) LerpSelf(t float32, vecTo *Vector3) {
	tmp := *result
	result.Lerp(t, &tmp, vecTo)
}

func (result *Vector3) Slerp(t float32, unitVec0, unitVec1 *Vector3) {
	var tmpV3 Vector3
	var scale0, scale1 float32
	cosAngle := unitVec0.Dot(unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := 1.0 / sin(angle)
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = 1.0 - t
		scale1 = t
	}

	tmpV3.ScalarMul(unitVec0, scale0)
	result.ScalarMul(unitVec1, scale1)
	result.AddToSelf(&tmpV3)
}

func (result *Vector3) SlerpSelf(t float32, vecTo *Vector3) {
	result.Slerp(t, result, vecTo)
}

//Vector4

func (result *Vector4) MakeFromV3(vec *Vector3) {
	result[x] = vec[x]
	result[y] = vec[y]
	result[z] = vec[z]
	result[w] = 0.0
}

func (result *Vector4) MakeFromP3(pnt *Point3) {
	result[x] = pnt[x]
	result[y] = pnt[y]
	result[z] = pnt[z]
	result[w] = 1.0
}

func (result *Vector4) MakeFromQ(quat *Quaternion) {
	result[x] = quat[x]
	result[y] = quat[y]
	result[z] = quat[z]
	result[w] = quat[w]
}

func (result *Vector4) MakeFromScalar(scalar float32) {
	result[x] = scalar
	result[y] = scalar
	result[z] = scalar
	result[w] = scalar
}

func (v *Vector4) MakeXAxis() {
	v[x] = 1.0
	v[y] = 0.0
	v[z] = 0.0
	v[w] = 0.0
}

func (v *Vector4) MakeYAxis() {
	v[x] = 0.0
	v[y] = 1.0
	v[z] = 0.0
	v[w] = 0.0
}

func (v *Vector4) MakeZAxis() {
	v[x] = 0.0
	v[y] = 0.0
	v[z] = 1.0
	v[w] = 0.0
}

func (v *Vector4) MakeWAxis() {
	v[x] = 0.0
	v[y] = 0.0
	v[z] = 0.0
	v[w] = 1.0
}

func (result *Vector4) Lerp(t float32, vec0, vec1 *Vector4) {
	var tmpV4_0, tmpV4_1 Vector4
	tmpV4_0.Sub(vec1, vec0)
	tmpV4_1.ScalarMul(&tmpV4_0, t)
	result.Add(vec0, &tmpV4_1)
}

func (v *Vector4) LerpSelf(t float32, vecTo *Vector4) {
	v.Lerp(t, v, vecTo)
}

func (result *Vector4) Slerp(t float32, unitVec0, unitVec1 *Vector4) {
	var tmp_0, tmp_1 Vector4
	var scale0, scale1 float32
	cosAngle := unitVec0.Dot(unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := (1.0 / sin(angle))
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = (1.0 - t)
		scale1 = t
	}
	tmp_0.ScalarMul(unitVec0, scale0)
	tmp_1.ScalarMul(unitVec1, scale1)
	result.Add(&tmp_0, &tmp_1)
}

func (v *Vector4) SlerpSelf(t float32, vecTo *Vector4) {
	v.Slerp(t, v, vecTo)
}

func (v *Vector4) SetXYZ(vec *Vector3) {
	v[x] = vec[x]
	v[y] = vec[y]
	v[z] = vec[z]
}

func (vec *Vector4) XYZ(result *Vector3) {
	result[x] = vec[x]
	result[y] = vec[y]
	result[z] = vec[z]
}

func (result *Vector4) Add(vec0, vec1 *Vector4) {
	result[x] = vec0[x] + vec1[x]
	result[y] = vec0[y] + vec1[y]
	result[z] = vec0[z] + vec1[z]
	result[w] = vec0[w] + vec1[w]
}

func (result *Vector4) AddToSelf(vec *Vector4) {
	result.Add(result, vec)
}

func (result *Vector4) Sub(vec0, vec1 *Vector4) {
	result[x] = vec0[x] - vec1[x]
	result[y] = vec0[y] - vec1[y]
	result[z] = vec0[z] - vec1[z]
	result[w] = vec0[w] - vec1[w]
}

func (result *Vector4) SubFromSelf(vec *Vector4) {
	result.Sub(result, vec)
}

func (result *Vector4) ScalarMul(vec *Vector4, scalar float32) {
	result[x] = vec[x] * scalar
	result[y] = vec[y] * scalar
	result[z] = vec[z] * scalar
	result[w] = vec[w] * scalar
}

func (result *Vector4) ScalarMulSelf(scalar float32) {
	result.ScalarMul(result, scalar)
}

func (result *Vector4) ScalarDiv(vec *Vector4, scalar float32) {
	result[x] = vec[x] / scalar
	result[y] = vec[y] / scalar
	result[z] = vec[z] / scalar
	result[w] = vec[w] / scalar
}

func (result *Vector4) ScalarDivSelf(scalar float32) {
	result.ScalarDiv(result, scalar)
}

func (result *Vector4) Neg(vec *Vector4) {
	result[x] = -vec[x]
	result[y] = -vec[y]
	result[z] = -vec[z]
	result[w] = -vec[w]
}

func (v *Vector4) NegSelf() {
	v.Neg(v)
}

func (result *Vector4) MulPerElem(vec0, vec1 *Vector4) {
	result[x] = vec0[x] * vec1[x]
	result[y] = vec0[y] * vec1[y]
	result[z] = vec0[z] * vec1[z]
	result[w] = vec0[w] * vec1[w]
}

func (result *Vector4) MulPerElemSelf(vec *Vector4) {
	result.MulPerElem(result, vec)
}

func (result *Vector4) DivPerElem(vec0, vec1 *Vector4) {
	result[x] = vec0[x] / vec1[x]
	result[y] = vec0[y] / vec1[y]
	result[z] = vec0[z] / vec1[z]
	result[w] = vec0[w] / vec1[w]
}

func (result *Vector4) DivPerElemSelf(vec *Vector4) {
	result.DivPerElem(result, vec)
}

func (result *Vector4) RecipPerElem(vec *Vector4) {
	result[x] = 1.0 / vec[x]
	result[y] = 1.0 / vec[y]
	result[z] = 1.0 / vec[z]
	result[w] = 1.0 / vec[w]
}

func (result *Vector4) RecipPerElemSelf() {
	result.RecipPerElem(result)
}

func (result *Vector4) SqrtPerElem(vec *Vector4) {
	result[x] = sqrt(vec[x])
	result[y] = sqrt(vec[y])
	result[z] = sqrt(vec[z])
	result[w] = sqrt(vec[w])
}

func (result *Vector4) SqrtPerElemSelf() {
	result.SqrtPerElem(result)
}

func (result *Vector4) RsqrtPerElem(vec *Vector4) {
	result[x] = 1.0 / sqrt(vec[x])
	result[y] = 1.0 / sqrt(vec[y])
	result[z] = 1.0 / sqrt(vec[z])
	result[w] = 1.0 / sqrt(vec[w])
}

func (result *Vector4) RsqrtPerElemSelf() {
	result.RsqrtPerElem(result)
}

func (result *Vector4) AbsPerElem(vec *Vector4) {
	result[x] = abs(vec[x])
	result[y] = abs(vec[y])
	result[z] = abs(vec[z])
	result[w] = abs(vec[w])
}

func (result *Vector4) AbsPerElemSelf() {
	result.AbsPerElem(result)
}

func (result *Vector4) CopySignPerElem(vec0, vec1 *Vector4) {
	if vec1[x] < 0.0 {
		result[x] = -abs(vec0[x])
	} else {
		result[x] = abs(vec0[x])
	}
	if vec1[y] < 0.0 {
		result[y] = -abs(vec0[y])
	} else {
		result[y] = abs(vec0[y])
	}
	if vec1[z] < 0.0 {
		result[z] = -abs(vec0[z])
	} else {
		result[z] = abs(vec0[z])
	}
	if vec1[w] < 0.0 {
		result[w] = -abs(vec0[w])
	} else {
		result[w] = abs(vec0[w])
	}
}

func (result *Vector4) CopySignPerElemSelf(vec *Vector4) {
	result.CopySignPerElem(result, vec)
}

func (result *Vector4) MaxPerElem(vec0, vec1 *Vector4) {
	result[x] = max(vec0[x], vec1[x])
	result[y] = max(vec0[y], vec1[y])
	result[z] = max(vec0[z], vec1[z])
	result[w] = max(vec0[w], vec1[w])
}

func (result *Vector4) MaxPerElemSelf(vec *Vector4) {
	result.MaxPerElem(result, vec)
}

func (v *Vector4) MaxElem() float32 {
	var result float32
	result = max(v[x], v[y])
	result = max(v[z], result)
	result = max(v[w], result)
	return result
}

func (result *Vector4) MinPerElem(vec0, vec1 *Vector4) {
	result[x] = min(vec0[x], vec1[x])
	result[y] = min(vec0[y], vec1[y])
	result[z] = min(vec0[z], vec1[z])
	result[w] = min(vec0[w], vec1[w])
}

func (result *Vector4) MinPerElemSelf(vec *Vector4) {
	result.MinPerElem(result, vec)
}

func (v *Vector4) MinElem() float32 {
	var result float32
	result = min(v[x], v[y])
	result = min(v[z], result)
	result = min(v[w], result)
	return result
}

func (v *Vector4) Sum() float32 {
	var result float32
	result = v[x] + v[y] + v[z] + v[w]
	return result
}

func (v *Vector4) Dot(vec *Vector4) float32 {
	result := v[x] * vec[x]
	result += v[y] * vec[y]
	result += v[z] * vec[z]
	result += v[w] * vec[w]
	return result
}

func (v *Vector4) LengthSqr() float32 {
	result := v[x] * v[x]
	result += v[y] * v[y]
	result += v[z] * v[z]
	result += v[w] * v[w]
	return result
}

func (v *Vector4) Length() float32 {
	return sqrt(v.LengthSqr())
}

func (result *Vector4) Normalize(vec *Vector4) {
	lenSqr := vec.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result[x] = vec[x] * lenInv
	result[y] = vec[y] * lenInv
	result[z] = vec[z] * lenInv
	result[w] = vec[w] * lenInv
}

func (v *Vector4) NormalizeSelf() {
	v.Normalize(v)
}

func (result *Vector4) Select(vec0, vec1 *Vector4, select1 int) {
	if select1 != 0 {
		result[x] = vec1[x]
		result[y] = vec1[y]
		result[z] = vec1[z]
		result[w] = vec1[w]
	} else {
		result[x] = vec0[x]
		result[y] = vec0[y]
		result[z] = vec0[z]
		result[w] = vec0[w]
	}
}

//Point3

func (result *Point3) MakeFromV3(vec *Vector3) {
	result[x] = vec[x]
	result[y] = vec[y]
	result[z] = vec[z]
}

func (result *Point3) MakeFromScalar(scalar float32) {
	result[x] = scalar
	result[y] = scalar
	result[z] = scalar
}

func (result *Point3) Lerp(t float32, pnt0, pnt1 *Point3) {
	var tmpV3_0, tmpV3_1 Vector3
	tmpV3_0.P3Sub(pnt1, pnt0)
	tmpV3_1.ScalarMul(&tmpV3_0, t)

	result.AddV3(pnt0, &tmpV3_1)
}

func (p *Point3) LerpSelf(t float32, pointTo *Point3) {
	p.Lerp(t, p, pointTo)
}

func (result *Vector3) P3Sub(pnt0, pnt1 *Point3) {
	result[x] = pnt0[x] - pnt1[x]
	result[y] = pnt0[y] - pnt1[y]
	result[z] = pnt0[z] - pnt1[z]
}

func (result *Point3) AddV3(pnt0 *Point3, vec1 *Vector3) {
	result[x] = pnt0[x] + vec1[x]
	result[y] = pnt0[y] + vec1[y]
	result[z] = pnt0[z] + vec1[z]
}

func (result *Point3) AddV3ToSelf(vec1 *Vector3) {
	result.AddV3(result, vec1)
}

func (result *Point3) SubV3(pnt0 *Point3, vec1 *Vector3) {
	result[x] = pnt0[x] - vec1[x]
	result[y] = pnt0[y] - vec1[y]
	result[z] = pnt0[z] - vec1[z]
}

func (result *Point3) SubV3FromSelf(vec1 *Vector3) {
	result.SubV3(result, vec1)
}

func (result *Point3) MulPerElem(pnt0, pnt1 *Point3) {
	result[x] = pnt0[x] * pnt1[x]
	result[y] = pnt0[y] * pnt1[y]
	result[z] = pnt0[z] * pnt1[z]
}

func (result *Point3) MulPerElemSelf(pnt *Point3) {
	result.MulPerElem(result, pnt)
}

func (result *Point3) DivPerElem(pnt0, pnt1 *Point3) {
	result[x] = pnt0[x] / pnt1[x]
	result[y] = pnt0[y] / pnt1[y]
	result[z] = pnt0[z] / pnt1[z]
}

func (result *Point3) DivPerElemSelf(pnt *Point3) {
	result.DivPerElem(result, pnt)
}

func (result *Point3) RecipPerElem(pnt *Point3) {
	result[x] = 1.0 / pnt[x]
	result[y] = 1.0 / pnt[y]
	result[z] = 1.0 / pnt[z]
}

func (result *Point3) RecipPerElemSelf() {
	result.RecipPerElem(result)
}

func (result *Point3) SqrtPerElem(pnt *Point3) {
	result[x] = sqrt(pnt[x])
	result[y] = sqrt(pnt[y])
	result[z] = sqrt(pnt[z])
}

func (result *Point3) SqrtPerElemSelf() {
	result.SqrtPerElem(result)
}

func (result *Point3) RsqrtPerElem(pnt *Point3) {
	result[x] = 1.0 / sqrt(pnt[x])
	result[y] = 1.0 / sqrt(pnt[y])
	result[z] = 1.0 / sqrt(pnt[z])
}

func (result *Point3) RsqrtPerElemSelf() {
	result.RsqrtPerElem(result)
}

func (result *Point3) AbsPerElem(pnt *Point3) {
	result[x] = abs(pnt[x])
	result[y] = abs(pnt[y])
	result[z] = abs(pnt[z])
}

func (result *Point3) AbsPerElemSelf() {
	result.AbsPerElem(result)
}

func (result *Point3) CopySignPerElem(pnt0, pnt1 *Point3) {
	if pnt1[x] < 0.0 {
		result[x] = -abs(pnt0[x])
	} else {
		result[x] = abs(pnt0[x])
	}
	if pnt1[y] < 0.0 {
		result[y] = -abs(pnt0[y])
	} else {
		result[y] = abs(pnt0[y])
	}
	if pnt1[z] < 0.0 {
		result[z] = -abs(pnt0[z])
	} else {
		result[z] = abs(pnt0[z])
	}
}

func (result *Point3) CopySignPerElemSelf(pnt *Point3) {
	result.CopySignPerElem(result, pnt)
}

func (result *Point3) MaxPerElem(pnt0, pnt1 *Point3) {
	result[x] = max(pnt0[x], pnt1[x])
	result[y] = max(pnt0[y], pnt1[y])
	result[z] = max(pnt0[z], pnt1[z])
}

func (result *Point3) MaxPerElemSelf(pnt *Point3) {
	result.MaxPerElem(result, pnt)
}

func (p *Point3) MaxElem() float32 {
	var result float32
	result = max(p[x], p[y])
	result = max(p[z], result)
	return result
}

func (result *Point3) MinPerElem(pnt0, pnt1 *Point3) {
	result[x] = min(pnt0[x], pnt1[x])
	result[y] = min(pnt0[y], pnt1[y])
	result[z] = min(pnt0[z], pnt1[z])
}

func (result *Point3) MinPerElemSelf(pnt *Point3) {
	result.MinPerElem(result, pnt)
}

func (p *Point3) MinElem() float32 {
	var result float32
	result = min(p[x], p[y])
	result = min(p[z], result)
	return result
}

func (p *Point3) Sum() float32 {
	var result float32
	result = p[x] + p[y] + p[z]
	return result
}

func (result *Point3) Scale(pnt *Point3, scaleVal float32) {
	var tmp_0 Point3
	tmp_0.MakeFromScalar(scaleVal)
	result.MulPerElem(pnt, &tmp_0)
}

func (result *Point3) ScaleSelf(scaleVal float32) {
	result.Scale(result, scaleVal)
}

func (result *Point3) NonUniformScale(pnt *Point3, scaleVec *Vector3) {
	var tmp_0 Point3
	tmp_0.MakeFromV3(scaleVec)
	result.MulPerElem(pnt, &tmp_0)
}

func (result *Point3) NonUniformScaleSelf(scaleVec *Vector3) {
	result.NonUniformScale(result, scaleVec)
}

func (p *Point3) Projection(unitVec *Vector3) float32 {
	result := p[x] * unitVec[x]
	result += p[y] * unitVec[y]
	result += p[z] * unitVec[z]
	return result
}

func (p *Point3) DistSqrFromOrigin() float32 {
	var tmpV3_0 Vector3
	tmpV3_0.MakeFromP3(p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) DistFromOrigin() float32 {
	var tmpV3_0 Vector3
	tmpV3_0.MakeFromP3(p)
	return tmpV3_0.Length()
}

func (p *Point3) DistSqr(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	tmpV3_0.P3Sub(pnt1, p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) Dist(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	tmpV3_0.P3Sub(pnt1, p)
	return tmpV3_0.Length()
}

func (result *Point3) Select(pnt0, pnt1 *Point3, select1 int) {
	if select1 != 0 {
		result[x] = pnt1[x]
		result[y] = pnt1[y]
		result[z] = pnt1[z]
	} else {
		result[x] = pnt0[x]
		result[y] = pnt0[y]
		result[z] = pnt0[z]
	}
}
