// Copyright 2013 Tim Shannon. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package vmath

import (
	"github.com/spate/vectormath"
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

	V3Slerp(vResult, time, vStart, vEnd)
	vectormath.V3Slerp(sResult, time, sStart, sEnd)

	if !v3IsEqual(vResult, sResult) {
		t.Error("Slerp not equal", vStart, sResult)

	}

	vStart.Slerp(time, vEnd)

	if !v3IsEqual(vStart, sResult) {
		t.Error("Slerp not equal", vStart, sResult)
	}

}
