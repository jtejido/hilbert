/*
Copyright 2014 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package hilbert

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/big"
	"math/bits"
	"testing"
)

func TestHilbert(t *testing.T) {
	hilbert, _ := New(5, 2) // testing at 2 dimensions
	h := hilbert.Encode(0, 0)
	res := hilbert.Decode(h)
	assert.Equal(t, uint64(0), h.Uint64())
	assert.Equal(t, uint64(0), res[0])
	assert.Equal(t, uint64(0), res[1])

	h = hilbert.Encode(1, 0)
	res = hilbert.Decode(h)
	assert.Equal(t, uint64(3), h.Uint64())
	assert.Equal(t, uint64(1), res[0])
	assert.Equal(t, uint64(0), res[1])

	h = hilbert.Encode(1, 1)
	res = hilbert.Decode(h)
	assert.Equal(t, uint64(2), h.Uint64())
	assert.Equal(t, uint64(1), res[0])
	assert.Equal(t, uint64(1), res[1])

	h = hilbert.Encode(0, 1)
	res = hilbert.Decode(h)
	assert.Equal(t, uint64(1), h.Uint64())
	assert.Equal(t, uint64(0), res[0])
	assert.Equal(t, uint64(1), res[1])
}

func TestHilbertAtMaxRange(t *testing.T) {
	hilbert, _ := New(uint32(bits.Len64(math.MaxInt64)), 3) // testing at 3 dimensions and 63 (bits.Len64(math.MaxInt64)) bits
	x, y, z := uint64(math.MaxInt64), uint64(math.MaxInt64), uint64(math.MaxInt64)
	h := hilbert.Encode(x, y, z)
	result := hilbert.Decode(h)
	assert.Equal(t, x, result[0])
	assert.Equal(t, y, result[1])
	assert.Equal(t, z, result[2])
}

func BenchmarkEncode(b *testing.B) {
	hilbert, _ := New(5, 3) // testing at 3-dimensions and 5-bits
	for i := 0; i < b.N; i++ {
		hilbert.Encode(uint64(i), uint64(i), uint64(i))
	}
}

func BenchmarkDecode(b *testing.B) {
	hilbert, _ := New(5, 3) // testing at 3-dimensions
	for i := 0; i < b.N; i++ {
		hilbert.Decode(big.NewInt(int64(i)))
	}
}

func BenchmarkEncodeAtMaxRange(b *testing.B) {
	hilbert, _ := New(uint32(bits.Len64(math.MaxInt64)), 3) // testing at 3-dimensions and max-bits
	for i := 0; i < b.N; i++ {
		hilbert.Encode(uint64(i), uint64(i), uint64(i))
	}
}

func BenchmarkDecodeAtMaxRange(b *testing.B) {
	hilbert, _ := New(uint32(bits.Len64(math.MaxInt64)), 3) // testing at 3-dimensions and max-bits
	for i := 0; i < b.N; i++ {
		hilbert.Decode(big.NewInt(int64(i)))
	}
}
