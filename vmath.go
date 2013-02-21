//Copyright (C) 2006, 2007 Sony Computer Entertainment Inc.
//  All rights reserved.
// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

const (
	x = iota
	y
	z
	w
)

type Vector3 [3]float32

func (v *Vector3) Array() *[3]float32 {
	return (*[3]float32)(v)
}

type Vector4 [4]float32

func (v *Vector4) Array() *[4]float32 {
	return (*[4]float32)(v)
}

type Point3 [3]float32

func (p *Point3) Array() *[3]float32 {
	return (*[3]float32)(p)
}

type Quaternion [4]float32

func (q *Quaternion) Array() *[4]float32 {
	return (*[4]float32)(q)
}

type Matrix3 [3 * 3]float32

func (m *Matrix3) Array() *[3 * 3]float32 {
	return (*[3 * 3]float32)(m)
}

type Matrix4 [4 * 4]float32

func (m *Matrix4) Array() *[4 * 4]float32 {
	return (*[4 * 4]float32)(m)
}

type Transform3 [3 * 4]float32

func (t *Transform3) Array() *[3 * 4]float32 {
	return (*[3 * 4]float32)(t)
}
