package hidden

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundTrip(t *testing.T) {
	str := "H"
	encoded := ToString([]byte(str))
	rt, err := FromString(encoded)
	if assert.NoError(t, err) {
		assert.Equal(t, str, string(rt))
	}
}

func TestExtract(t *testing.T) {
	a := []byte("Here is my string")
	d := make([]byte, 8)
	binary.BigEndian.PutUint64(d, 56)
	str := fmt.Sprintf("hidden%s \x11data%s is fun", ToString(a), ToString(d))
	t.Log(str)
	out, err := Extract(str)
	if assert.NoError(t, err) {
		if assert.Len(t, out, 2) {
			assert.Equal(t, out, [][]byte{a, d})
		}
	}
	assert.Equal(t, "hidden data is fun", Clean(str))
}

func BenchmarkClean(b *testing.B) {
	a := []byte("Here is my string")
	d := make([]byte, 8)
	binary.BigEndian.PutUint64(d, 56)
	str := fmt.Sprintf("hidden%s data%s is fun", ToString(a), ToString(d))
	b.ResetTimer()
	for i:=0; i<b.N; i++ {
		Clean(str)
	}
}