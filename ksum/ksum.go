// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.9 Hash and a simple test in Go
// base code downloaded from miguelmota/sha3.go

package main

import (
  "encoding/hex"
  "io/ioutil"
  "os"
  "flag"
  "fmt"
  "encoding/json"
  "sha3"
)

type TransactionType struct	{
	TxHash		string
	TxIn			int
	TxOut			int
}

func main() {

// 	var buf []byte

// Parse CLI arguments to this program, --cfg <name>.json for input and output file names
	flag.Parse()
	fns := flag.Args()
	if len(fns) == 0	{
		fmt.Fprintf(os.Stderr, "Usage: .files to be hashed\n")
		os.Exit(1)
	}
	var Cfg = flag.String("cfg", fns[0], "config file for this call")
	if Cfg == nil	{
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

// read in data using ReadFile from in file argument fns[ii] as binary slice
	fmt.Printf("\n")
	for ii := 0; ii < len(fns); ii++	{
		readin, err := ioutil.ReadFile(fns[ii])
  	if err != nil	{
			fmt.Printf("error from ReadFile\n")
			os.Exit(1)
		}
		fmt.Printf("%2d %10s: %x\n", ii+1, fns[ii], Keccak256(readin))
	}
	fmt.Printf("\n")
  // with sample files file1, file2.txt. file3 hashes should be:
	// file1 ecd67ca5a72802084fcea4883b6877ecfba7f95c0aece07ea504359d54eb4610
	// file2.txt 0695253b82a83d557392ab196ff309a1fedc6cbab0d7d4186d2664dcec92b5ff
	// file3 fb15d651aaf994584aa6da109b5dba096de83bf2f44da6a224cf41d8d5e92f14
	}

func IndentJSON(v interface{}) string	{
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil	{
		return fmt.Sprintf("Error: %s", err)
	} else {
    fmt.Printf("%s\n", s)
		return string(s)
	}
}

// decode string of hex to bytes
func decodeHex(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	d := sha3.NewKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}
