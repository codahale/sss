package sss

import (
	"fmt"
)

func Example() {
	secret := "well hello there!" // our secret
	n := 30                       // create 30 shares
	k := 2                        // require 2 of them to combine

	shares, err := Split(n, k, []byte(secret)) // split into 30 shares
	if err != nil {
		fmt.Println(err)
		return
	}

	// select a random subset of the total shares
	subset := make(map[int][]byte, k)
	for x, y := range shares { // just iterate since maps are randomized
		subset[x] = y
		if len(subset) == k {
			break
		}
	}

	// combine two shares and recover the secret
	recovered := string(Combine(subset))
	fmt.Println(recovered)

	// Output: well hello there!
}
