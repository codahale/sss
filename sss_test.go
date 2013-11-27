package sss

import (
	"fmt"
)

func Example() {
	// split into 30 shares, of which only 2 are required to combine
	n := 30
	k := 2

	secret := "well hello there!"
	shares, err := Split(n, k, []byte(secret))
	if err != nil {
		fmt.Println(err)
	}

	// select a random subset of the total shares
	subset := make(map[int][]byte, k)
	for x, y := range shares {
		subset[x] = y
		if len(subset) == k {
			break
		}
	}

	fmt.Println(string(Combine(subset)))
	// Output: well hello there!
}
