package vmath

import (
	"testing"
)

type AVector [4]float32

type SVector struct {
	X, Y, Z, W float32
}

type AMatrix [4 * 4]float32

type SMatrix struct {
	Col0, Col1, Col2 SVector
}
