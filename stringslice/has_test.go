package stringslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHas(t *testing.T) {
	assert.True(t, Has([]string{"foo", "bar"}, "foo"))
	assert.True(t, Has([]string{"foo", "bar"}, "bar"))
	assert.False(t, Has([]string{"foo", "bar"}, "baz"))
	assert.False(t, Has([]string{"foo", "bar"}, "baR"))
}

func TestHasI(t *testing.T) {
	assert.True(t, HasI([]string{"foO", "bAr"}, "foo"))
	assert.True(t, HasI([]string{"foo", "baR"}, "bar"))
	assert.False(t, HasI([]string{"foo", "bar"}, "baz"))
}

func BenchmarkHasI(b *testing.B) {
	var heystack = []string{
		"apple",
		"orange",
		"banana",
		"strawberry",
		"cherry",
		"pear",
		"pineapple",
	}
	for i := 0; i < b.N; i++ {
		HasI(heystack, "cherry")
	}
}
