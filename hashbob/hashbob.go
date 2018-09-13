// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.9a Hashbob example
// Keccak256 From: https://github.com/ethereum/go-ethereum/blob/master/crypto/crypto.go lines 44...51.

package main

	import (
		"fmt"
    "sha3"
	)

	func main() {
    // type cast from string, "bob", to slice of byte.
		fmt.Printf("%x\n", Keccak256([]byte("bob"))) 	}

	// Keccak256 calculates and returns the Keccak256 hash of the input data.
	func Keccak256(data ...[]byte) []byte {
		d := sha3.NewKeccak256()
		for _, b := range data {
			d.Write(b)
		}
		return d.Sum(nil)
	}
