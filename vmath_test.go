// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

import (
	"fmt"
	"github.com/timshannon/vectormath"
	"testing"
)

func v3IsEqual(vector *Vector3, other *vectormath.Vector3) bool {
	if vector[0] != other.X {
		return false
	}

	if vector[1] != other.Y {
		return false
	}

	if vector[2] != other.Z {
		return false
	}

	return true

}

func m3IsEqual(t *testing.T, matrix *Matrix3, other *vectormath.Matrix3) {

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if matrix.Elem(c, r) != other.GetElem(c, r) {
				t.Error("M3 Not equal: ", matrix.Elem(c, r), other.GetElem(c, r))
			}
		}
	}

}

func TestEquality(t *testing.T) {
	var time float32 = 0.4
	sStart := new(vectormath.Vector3)
	sEnd := new(vectormath.Vector3)
	sResult := new(vectormath.Vector3)
	vectormath.V3MakeFromElems(sStart, 1, 1, 1)
	vectormath.V3MakeFromElems(sEnd, 3, 3, 3)

	vStart := &Vector3{1, 1, 1}
	vEnd := &Vector3{3, 3, 3}
	vResult := &Vector3{}

	vResult.Slerp(time, vStart, vEnd)
	vectormath.V3Slerp(sResult, time, sStart, sEnd)

	if !v3IsEqual(vResult, sResult) {
		t.Error("Slerp not equal", vStart, sResult)

	}

	vStart.SlerpSelf(time, vEnd)

	if !v3IsEqual(vStart, sResult) {
		t.Error("Slerp not equal", vStart, sResult)
	}

}

func TestM3Equality(t *testing.T) {
	othermat0 := new(vectormath.Matrix3)
	othermat1 := new(vectormath.Matrix3)
	mat0 := &Matrix3{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	mat1 := &Matrix3{23.41, 21.12, 0,
		214.23, 213.9821, -32.02,
		991.90, 75.123, -231.02}

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			//column major
			othermat0.SetElem(c, r, mat0.Elem(c, r))
			othermat1.SetElem(c, r, mat1.Elem(c, r))
		}
	}

	m3IsEqual(t, mat0, othermat0)
	m3IsEqual(t, mat1, othermat1)

	vectormath.M3Mul(othermat0, othermat0, othermat1)
	mat0.Mul(mat0, mat1)

	fmt.Println("After")
	m3IsEqual(t, mat0, othermat0)

}
