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
)

type Vector3 [3]float32

type Vector4 [4]float32

type Point3 [3]float32

type Quaternion [4]float32

type Matrix3 [3 * 3]float32

type Matrix4 [4 * 4]float32

type Transform3 [3 * 4]float32
