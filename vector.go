//Copyright (C) 2006, 2007 Sony Computer Entertainment Inc.
//  All rights reserved.
// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

const g_SLERP_TOL = 0.999

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

func V3Add(result, vec0, vec1 *Vector3) {
	result[x] = vec0[x] + vec1[x]
	result[y] = vec0[y] + vec1[y]
	result[z] = vec0[z] + vec1[z]
}

func V3Sub(result, vec0, vec1 *Vector3) {
	result[x] = vec0[x] - vec1[x]
	result[y] = vec0[y] - vec1[y]
	result[z] = vec0[z] - vec1[z]
}

func V3AddP3(result, vec0 *Vector3, pnt1 *Point3) {
	result[x] = vec0[x] + pnt1[x]
	result[y] = vec0[y] + pnt1[y]
	result[z] = vec0[z] + pnt1[z]
}

func V3ScalarMul(result, vec *Vector3, scalar float32) {
	result[x] = vec[x] * scalar
	result[y] = vec[y] * scalar
	result[z] = vec[z] * scalar
}

func V3ScalarDiv(result, vec *Vector3, scalar float32) {
	result[x] = vec[x] / scalar
	result[y] = vec[y] / scalar
	result[z] = vec[z] / scalar
}

func V3Neg(result, vec *Vector3) {
	result[x] = -vec[x]
	result[y] = -vec[y]
	result[z] = -vec[z]
}

func (v *Vector3) Neg() {
	V3Neg(v, v)
}

func V3MulPerElem(result, vec0, vec1 *Vector3) {
	result[x] = vec0[x] * vec1[x]
	result[y] = vec0[y] * vec1[y]
	result[z] = vec0[z] * vec1[z]
}

func V3DivPerElem(result, vec0, vec1 *Vector3) {
	result[x] = vec0[x] / vec1[x]
	result[y] = vec0[y] / vec1[y]
	result[z] = vec0[z] / vec1[z]
}

func V3RecipPerElem(result, vec *Vector3) {
	result[x] = 1.0 / vec[x]
	result[y] = 1.0 / vec[y]
	result[z] = 1.0 / vec[z]
}

func V3SqrtPerElem(result, vec *Vector3) {
	result[x] = sqrt(vec[x])
	result[y] = sqrt(vec[y])
	result[z] = sqrt(vec[z])
}

func V3RsqrtPerElem(result, vec *Vector3) {
	result[x] = 1.0 / sqrt(vec[x])
	result[y] = 1.0 / sqrt(vec[y])
	result[z] = 1.0 / sqrt(vec[z])
}

func V3AbsPerElem(result, vec *Vector3) {
	result[x] = abs(vec[x])
	result[y] = abs(vec[y])
	result[z] = abs(vec[z])
}

func V3CopySignPerElem(result, vec0, vec1 *Vector3) {
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

func V3MaxPerElem(result, vec0, vec1 *Vector3) {
	result[x] = max(vec0[x], vec1[x])
	result[y] = max(vec0[y], vec1[y])
	result[z] = max(vec0[z], vec1[z])
}

func (v *Vector3) MaxElem() float32 {
	var result float32
	result = max(v[x], v[y])
	result = max(v[z], result)
	return result
}

func V3MinPerElem(result, vec0, vec1 *Vector3) {
	result[x] = min(vec0[x], vec1[x])
	result[y] = min(vec0[y], vec1[y])
	result[z] = min(vec0[z], vec1[z])
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

func V3Dot(vec0, vec1 *Vector3) float32 {
	result := vec0[x] * vec1[x]
	result += vec0[y] * vec1[y]
	result += vec0[z] * vec1[z]
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

func V3Normalize(result, vec *Vector3) {
	lenSqr := vec.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result[x] = vec[x] * lenInv
	result[y] = vec[y] * lenInv
	result[z] = vec[z] * lenInv
}

func (v *Vector3) Normalize() {
	V3Normalize(v, v)
}

func V3Cross(result, vec0, vec1 *Vector3) {
	result[x] = vec0[y]*vec1[z] - vec0[z]*vec1[y]
	result[y] = vec0[z]*vec1[x] - vec0[x]*vec1[z]
	result[z] = vec0[x]*vec1[y] - vec0[y]*vec1[x]
}

func V3Select(result, vec0, vec1 *Vector3, select1 int) {
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

func V3Velocity(result, start, end *Vector3, elapsedTime float32) {
	//change in position / elapsedTime
	V3Sub(result, start, end)
	result[x] = result[x] / elapsedTime
	result[y] = result[y] / elapsedTime
	result[z] = result[z] / elapsedTime
}
func V3Lerp(result *Vector3, t float32, vec0, vec1 *Vector3) {
	V3Sub(result, vec1, vec0)
	V3ScalarMul(result, result, t)
	V3Add(result, vec0, result)
}

func (v *Vector3) Lerp(t float32, vecTo *Vector3) {
	//tmp := &Vector3{}
	//copy(tmp, *v)
	//V3Lerp(v, t, tmp, vecTo)
}

/*
func V3Slerp(result *Vector3, t float32, unitVec0, unitVec1 *Vector3) {
	var tmpV3_0, tmpV3_1 Vector3
	var scale0, scale1 float32
	cosAngle := V3Dot(unitVec0, unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := 1.0 / sin(angle)
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = 1.0 - t
		scale1 = t
	}
	V3ScalarMul(&tmpV3_0, unitVec0, scale0)
	V3ScalarMul(&tmpV3_1, unitVec1, scale1)
	V3Add(result, &tmpV3_0, &tmpV3_1)
}

func (v *Vector3) Slerp(t float32, vecTo *Vector3) {
	V3Slerp(v, t, v, vecTo)
}

func (v *Vector3) SetElem(index int, value float32) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	}
}





func V4Copy(result, vec *Vector4) {
	result.X = vec.X
	result.Y = vec.Y
	result.Z = vec.Z
	result.W = vec.W
}

func V4MakeFromElems(result *Vector4, x, y, z, w float32) {
	result.X = x
	result.Y = y
	result.Z = z
	result.W = w
}

func V4MakeFromV3Scalar(result *Vector4, xyz *Vector3, w float32) {
	result.SetXYZ(xyz)
	result.W = w
}

func V4MakeFromV3(result *Vector4, vec *Vector3) {
	result.X = vec.X
	result.Y = vec.Y
	result.Z = vec.Z
	result.W = 0.0
}

func V4MakeFromP3(result *Vector4, pnt *Point3) {
	result.X = pnt.X
	result.Y = pnt.Y
	result.Z = pnt.Z
	result.W = 1.0
}

func V4MakeFromQ(result *Vector4, quat *Quat) {
	result.X = quat.X
	result.Y = quat.Y
	result.Z = quat.Z
	result.W = quat.W
}

func V4MakeFromScalar(result *Vector4, scalar float32) {
	result.X = scalar
	result.Y = scalar
	result.Z = scalar
	result.W = scalar
}

func V4MakeXAxis(result *Vector4) {
	V4MakeFromElems(result, 1.0, 0.0, 0.0, 0.0)
}

func V4MakeYAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 1.0, 0.0, 0.0)
}

func V4MakeZAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 0.0, 1.0, 0.0)
}

func V4MakeWAxis(result *Vector4) {
	V4MakeFromElems(result, 0.0, 0.0, 0.0, 1.0)
}

func V4Lerp(result *Vector4, t float32, vec0, vec1 *Vector4) {
	var tmpV4_0, tmpV4_1 Vector4
	V4Sub(&tmpV4_0, vec1, vec0)
	V4ScalarMul(&tmpV4_1, &tmpV4_0, t)
	V4Add(result, vec0, &tmpV4_1)
}

func (v *Vector4) Lerp(t float32, vecTo *Vector4) {
	V4Lerp(v, t, v, vecTo)
}

func (v *Vector4) Slerp(t float32, vecTo *Vector4) {
	V4Slerp(v, t, v, vecTo)
}

func V4Slerp(result *Vector4, t float32, unitVec0, unitVec1 *Vector4) {
	var tmpV4_0, tmpV4_1 Vector4
	var scale0, scale1 float32
	cosAngle := V4Dot(unitVec0, unitVec1)
	if cosAngle < g_SLERP_TOL {
		angle := acos(cosAngle)
		recipSinAngle := (1.0 / sin(angle))
		scale0 = (sin(((1.0 - t) * angle)) * recipSinAngle)
		scale1 = (sin((t * angle)) * recipSinAngle)
	} else {
		scale0 = (1.0 - t)
		scale1 = t
	}
	V4ScalarMul(&tmpV4_0, unitVec0, scale0)
	V4ScalarMul(&tmpV4_1, unitVec1, scale1)
	V4Add(result, &tmpV4_0, &tmpV4_1)
}

func (v *Vector4) SetXYZ(vec *Vector3) {
	v.X = vec.X
	v.Y = vec.Y
	v.Z = vec.Z
}

func V4GetXYZ(result *Vector3, vec *Vector4) {
	V3MakeFromElems(result, vec.X, vec.Y, vec.Z)
}

func (v *Vector4) SetElem(index int, value float32) {
	switch index {
	case 0:
		v.X = value
	case 1:
		v.Y = value
	case 2:
		v.Z = value
	case 3:
		v.W = value
	}
}

func (v *Vector4) GetElem(index int) float32 {
	switch index {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	case 3:
		return v.W
	}
	return 0
}

func V4Add(result, vec0, vec1 *Vector4) {
	result.X = vec0.X + vec1.X
	result.Y = vec0.Y + vec1.Y
	result.Z = vec0.Z + vec1.Z
	result.W = vec0.W + vec1.W
}

func V4Sub(result, vec0, vec1 *Vector4) {
	result.X = vec0.X - vec1.X
	result.Y = vec0.Y - vec1.Y
	result.Z = vec0.Z - vec1.Z
	result.W = vec0.W - vec1.W
}

func V4ScalarMul(result, vec *Vector4, scalar float32) {
	result.X = vec.X * scalar
	result.Y = vec.Y * scalar
	result.Z = vec.Z * scalar
	result.W = vec.W * scalar
}

func V4ScalarDiv(result, vec *Vector4, scalar float32) {
	result.X = vec.X / scalar
	result.Y = vec.Y / scalar
	result.Z = vec.Z / scalar
	result.W = vec.W / scalar
}

func V4Neg(result, vec *Vector4) {
	result.X = -vec.X
	result.Y = -vec.Y
	result.Z = -vec.Z
	result.W = -vec.W
}

func (v *Vector4) Neg() {
	V4Neg(v, v)
}

func V4MulPerElem(result, vec0, vec1 *Vector4) {
	result.X = vec0.X * vec1.X
	result.Y = vec0.Y * vec1.Y
	result.Z = vec0.Z * vec1.Z
	result.W = vec0.W * vec1.W
}

func V4DivPerElem(result, vec0, vec1 *Vector4) {
	result.X = vec0.X / vec1.X
	result.Y = vec0.Y / vec1.Y
	result.Z = vec0.Z / vec1.Z
	result.W = vec0.W / vec1.W
}

func V4RecipPerElem(result, vec *Vector4) {
	result.X = 1.0 / vec.X
	result.Y = 1.0 / vec.Y
	result.Z = 1.0 / vec.Z
	result.W = 1.0 / vec.W
}

func V4SqrtPerElem(result, vec *Vector4) {
	result.X = sqrt(vec.X)
	result.Y = sqrt(vec.Y)
	result.Z = sqrt(vec.Z)
	result.W = sqrt(vec.W)
}

func V4RsqrtPerElem(result, vec *Vector4) {
	result.X = 1.0 / sqrt(vec.X)
	result.Y = 1.0 / sqrt(vec.Y)
	result.Z = 1.0 / sqrt(vec.Z)
	result.W = 1.0 / sqrt(vec.W)
}

func V4AbsPerElem(result, vec *Vector4) {
	result.X = abs(vec.X)
	result.Y = abs(vec.Y)
	result.Z = abs(vec.Z)
	result.W = abs(vec.W)
}

func V4CopySignPerElem(result, vec0, vec1 *Vector4) {
	if vec1.X < 0.0 {
		result.X = -abs(vec0.X)
	} else {
		result.X = abs(vec0.X)
	}
	if vec1.Y < 0.0 {
		result.Y = -abs(vec0.Y)
	} else {
		result.Y = abs(vec0.Y)
	}
	if vec1.Z < 0.0 {
		result.Z = -abs(vec0.Z)
	} else {
		result.Z = abs(vec0.Z)
	}
	if vec1.W < 0.0 {
		result.W = -abs(vec0.W)
	} else {
		result.W = abs(vec0.W)
	}
}

func V4MaxPerElem(result, vec0, vec1 *Vector4) {
	result.X = max(vec0.X, vec1.X)
	result.Y = max(vec0.Y, vec1.Y)
	result.Z = max(vec0.Z, vec1.Z)
	result.W = max(vec0.W, vec1.W)
}

func (v *Vector4) MaxElem() float32 {
	var result float32
	result = max(v.X, v.Y)
	result = max(v.Z, result)
	result = max(v.W, result)
	return result
}

func V4MinPerElem(result, vec0, vec1 *Vector4) {
	result.X = min(vec0.X, vec1.X)
	result.Y = min(vec0.Y, vec1.Y)
	result.Z = min(vec0.Z, vec1.Z)
	result.W = min(vec0.W, vec1.W)
}

func (v *Vector4) MinElem() float32 {
	var result float32
	result = min(v.X, v.Y)
	result = min(v.Z, result)
	result = min(v.W, result)
	return result
}

func (v *Vector4) Sum() float32 {
	var result float32
	result = v.X + v.Y + v.Z + v.W
	return result
}

func V4Dot(vec0, vec1 *Vector4) float32 {
	result := vec0.X * vec1.X
	result += vec0.Y * vec1.Y
	result += vec0.Z * vec1.Z
	result += vec0.W * vec1.W
	return result
}

func (v *Vector4) Dot(vec1 *Vector4) float32 {
	result := v.X * vec1.X
	result += v.Y * vec1.Y
	result += v.Z * vec1.Z
	result += v.W * vec1.W
	return result
}

func (v *Vector4) LengthSqr() float32 {
	result := v.X * v.X
	result += v.Y * v.Y
	result += v.Z * v.Z
	result += v.W * v.W
	return result
}

func (v *Vector4) Length() float32 {
	return sqrt(v.LengthSqr())
}

func V4Normalize(result, vec *Vector4) {
	lenSqr := vec.LengthSqr()
	lenInv := 1.0 / sqrt(lenSqr)
	result.X = vec.X * lenInv
	result.Y = vec.Y * lenInv
	result.Z = vec.Z * lenInv
	result.W = vec.W * lenInv
}

func (v *Vector4) Normalize() {
	V4Normalize(v, v)
}

func V4Select(result, vec0, vec1 *Vector4, select1 int) {
	if select1 != 0 {
		result.X = vec1.X
		result.Y = vec1.Y
		result.Z = vec1.Z
		result.W = vec1.W
	} else {
		result.X = vec0.X
		result.Y = vec0.Y
		result.Z = vec0.Z
		result.W = vec0.W
	}
}

func (v *Vector4) String() string {
	return fmt.Sprintf("( %f %f %f %f )", v.X, v.Y, v.Z, v.W)
}


func P3Copy(result, pnt *Point3) {
	result.X = pnt.X
	result.Y = pnt.Y
	result.Z = pnt.Z
}

func P3MakeFromElems(result *Point3, x, y, z float32) {
	result.X = x
	result.Y = y
	result.Z = z
}

func P3MakeFromV3(result *Point3, vec *Vector3) {
	result.X = vec.X
	result.Y = vec.Y
	result.Z = vec.Z
}

func P3MakeFromScalar(result *Point3, scalar float32) {
	result.X = scalar
	result.Y = scalar
	result.Z = scalar
}

func P3Lerp(result *Point3, t float32, pnt0, pnt1 *Point3) {
	var tmpV3_0, tmpV3_1 Vector3
	P3Sub(&tmpV3_0, pnt1, pnt0)
	V3ScalarMul(&tmpV3_1, &tmpV3_0, t)
	P3AddV3(result, pnt0, &tmpV3_1)
}

func (p *Point3) Lerp(t float32, pointTo *Point3) {
	P3Lerp(p, t, p, pointTo)
}

func (p *Point3) SetElem(index int, value float32) {
	switch index {
	case 0:
		p.X = value
	case 1:
		p.Y = value
	case 2:
		p.Z = value
	}
}

func (p *Point3) GetElem(index int) float32 {
	switch index {
	case 0:
		return p.X
	case 1:
		return p.Y
	case 2:
		return p.Z
	}
	return 0
}

func P3Sub(result *Vector3, pnt0, pnt1 *Point3) {
	result.X = pnt0.X - pnt1.X
	result.Y = pnt0.Y - pnt1.Y
	result.Z = pnt0.Z - pnt1.Z
}

func P3AddV3(result, pnt0 *Point3, vec1 *Vector3) {
	result.X = pnt0.X + vec1.X
	result.Y = pnt0.Y + vec1.Y
	result.Z = pnt0.Z + vec1.Z
}

func P3SubV3(result, pnt0 *Point3, vec1 *Vector3) {
	result.X = pnt0.X - vec1.X
	result.Y = pnt0.Y - vec1.Y
	result.Z = pnt0.Z - vec1.Z
}

func P3MulPerElem(result, pnt0, pnt1 *Point3) {
	result.X = pnt0.X * pnt1.X
	result.Y = pnt0.Y * pnt1.Y
	result.Z = pnt0.Z * pnt1.Z
}

func P3DivPerElem(result, pnt0, pnt1 *Point3) {
	result.X = pnt0.X / pnt1.X
	result.Y = pnt0.Y / pnt1.Y
	result.Z = pnt0.Z / pnt1.Z
}

func P3RecipPerElem(result, pnt *Point3) {
	result.X = 1.0 / pnt.X
	result.Y = 1.0 / pnt.Y
	result.Z = 1.0 / pnt.Z
}

func P3SqrtPerElem(result, pnt *Point3) {
	result.X = sqrt(pnt.X)
	result.Y = sqrt(pnt.Y)
	result.Z = sqrt(pnt.Z)
}

func P3RsqrtPerElem(result, pnt *Point3) {
	result.X = 1.0 / sqrt(pnt.X)
	result.Y = 1.0 / sqrt(pnt.Y)
	result.Z = 1.0 / sqrt(pnt.Z)
}

func P3AbsPerElem(result, pnt *Point3) {
	result.X = abs(pnt.X)
	result.Y = abs(pnt.Y)
	result.Z = abs(pnt.Z)
}

func P3CopySignPerElem(result, pnt0, pnt1 *Point3) {
	if pnt1.X < 0.0 {
		result.X = -abs(pnt0.X)
	} else {
		result.X = abs(pnt0.X)
	}
	if pnt1.Y < 0.0 {
		result.Y = -abs(pnt0.Y)
	} else {
		result.Y = abs(pnt0.Y)
	}
	if pnt1.Z < 0.0 {
		result.Z = -abs(pnt0.Z)
	} else {
		result.Z = abs(pnt0.Z)
	}
}

func P3MaxPerElem(result, pnt0, pnt1 *Point3) {
	result.X = max(pnt0.X, pnt1.X)
	result.Y = max(pnt0.Y, pnt1.Y)
	result.Z = max(pnt0.Z, pnt1.Z)
}

func (p *Point3) MaxElem() float32 {
	var result float32
	result = max(p.X, p.Y)
	result = max(p.Z, result)
	return result
}

func P3MinPerElem(result, pnt0, pnt1 *Point3) {
	result.X = min(pnt0.X, pnt1.X)
	result.Y = min(pnt0.Y, pnt1.Y)
	result.Z = min(pnt0.Z, pnt1.Z)
}

func (p *Point3) MinElem() float32 {
	var result float32
	result = min(p.X, p.Y)
	result = min(p.Z, result)
	return result
}

func (p *Point3) Sum() float32 {
	var result float32
	result = p.X + p.Y + p.Z
	return result
}

func P3Scale(result, pnt *Point3, scaleVal float32) {
	var tmpP3_0 Point3
	P3MakeFromScalar(&tmpP3_0, scaleVal)
	P3MulPerElem(result, pnt, &tmpP3_0)
}

func P3NonUniformScale(result, pnt *Point3, scaleVec *Vector3) {
	var tmpP3_0 Point3
	P3MakeFromV3(&tmpP3_0, scaleVec)
	P3MulPerElem(result, pnt, &tmpP3_0)
}

func (p *Point3) Projection(unitVec *Vector3) float32 {
	result := p.X * unitVec.X
	result += p.Y * unitVec.Y
	result += p.Z * unitVec.Z
	return result
}

func (p *Point3) DistSqrFromOrigin() float32 {
	var tmpV3_0 Vector3
	V3MakeFromP3(&tmpV3_0, p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) DistFromOrigin() float32 {
	var tmpV3_0 Vector3
	V3MakeFromP3(&tmpV3_0, p)
	return tmpV3_0.Length()
}

func (p *Point3) DistSqr(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	P3Sub(&tmpV3_0, pnt1, p)
	return tmpV3_0.LengthSqr()
}

func (p *Point3) Dist(pnt1 *Point3) float32 {
	var tmpV3_0 Vector3
	P3Sub(&tmpV3_0, pnt1, p)
	return tmpV3_0.Length()
}

func P3Select(result, pnt0, pnt1 *Point3, select1 int) {
	if select1 != 0 {
		result.X = pnt1.X
		result.Y = pnt1.Y
		result.Z = pnt1.Z
	} else {
		result.X = pnt0.X
		result.Y = pnt0.Y
		result.Z = pnt0.Z
	}
}

*/
