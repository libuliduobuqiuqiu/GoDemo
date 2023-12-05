package testingdemo

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)

	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rand.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)

	r := rand.New(rand.NewSource(seed))

	for i := 0; i < 100; i++ {
		p := randomPalindrome(r)
		fmt.Println(p)
	}
}
