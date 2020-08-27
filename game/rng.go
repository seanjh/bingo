package game

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	mrand "math/rand"
)

func safeSeed() int64 {
	var result int64
	err := binary.Read(crand.Reader, binary.LittleEndian, &result)
	if err != nil {
		fmt.Println("failed to read crypto/rand.Reader")
		return 0
	}
	return result
}

// NewRng returns a math.rand.Rand seeded with a safe random value.
func NewRng() *mrand.Rand {
	seed := safeSeed()
	source := mrand.NewSource(seed)
	return mrand.New(source)
}
