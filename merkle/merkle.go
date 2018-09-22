// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 3 Merkle Trees

package merkle

import (
	"github.com/Univ-Wyo-Education/Blockchain-4010-Fall-2018/Assignments/A-03/hash"
)

func MerkleHash(data [][]byte) (bt []byte)	{

	hTmp   		:= len(data)
	fflag  		:= 0
	loopflag	:= hTmp
	oddflag   := 0
	ii				:= 0

	// intial hash encoding of data slice
	for ii = 0; ii < hTmp; ii++ 	{
		data[ii] = hash.HashOf(data[ii])
	}

	// merkle tree if there is only one input
	if hTmp == 1	{
		data[0] = hash.Keccak256(data[0])
		return (data[0])
	}

	//loop through each level of the merkle tree until you have a single hash
	for loopflag > 1		{
		// determine if the input is odd, handle last item of array separately if odd
		if loopflag % 2 != 0	{
			loopflag = (loopflag/2)*2
			oddflag  = 1
		}
		// loop through creating merkle tree for all pairs
		for ii = 0; ii < loopflag; ii+= 2 	{
			data[fflag] = hash.Keccak256(data[ii], data[ii+1])
			fflag++
			}
		// if input is odd in lenght hash last odd item if applicable
		if oddflag == 1	{
			data[fflag] = hash.Keccak256(data[ii])
			fflag++
		}
		loopflag = fflag
		oddflag  = 0
		fflag    = 0
	}
	return (data[0])
}
