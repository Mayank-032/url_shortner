package hashfunction

import (
	"context"
	"fmt"
	"short-url/pkg/domain"
)

type RollingHash struct {
}

// hash[input_str] = s[0] + s[1]*p + s[2]*p*p + s[3]*p*p*p + ... + s[n]*p*p*p*..*p(n-times) % m
// p = nearest prime number (greater-one) to total distinct characters, an input string can have
// m = required for modulo to reduce number of collisions
func (rh RollingHash) HashFunction(ctx context.Context, inputToHash string) (string, bool, error) {
	var (
		p             = domain.NearestPrimeToTotalDistinctAlphabet
		m             = domain.Modulo
		pInitialValue = 1
		hashValue     = 0
		chunkSize     = 10
	)

	startPointer := 0
	endPointer := 0
	for endPointer < len(inputToHash) {
		if (endPointer + chunkSize) >= len(inputToHash) {
			startPointer = endPointer
			endPointer = len(inputToHash)
		} else {
			startPointer = endPointer
			endPointer = endPointer + chunkSize
		}

		chunk := inputToHash[startPointer:endPointer]

		for i := range chunk {
			asciiValue := (int(chunk[i]) - 65) + 1
			hashValue = (hashValue + (asciiValue * pInitialValue)) % m
			pInitialValue = (pInitialValue * p) % m
		}
	}

	isSignedHashValue := false
	if hashValue < 0 {
		hashValue = -1 * hashValue
		isSignedHashValue = true
	}
	return fmt.Sprintf("%v", hashValue), isSignedHashValue, nil
}
