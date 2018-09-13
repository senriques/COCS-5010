// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.8 copy a file with ioutil.ReadFile and ioutil.WriteFile

// This code was taken from sample at https://shapeshed.com/copy-a-file-in-go/

package main

import (
  "io/ioutil"
  "os"
  "flag"
  "fmt"
  "encoding/json"
)

type TransactionType struct	{
	TxHash		string
	TxIn			int
	TxOut			int
}

func main() {

// Parse CLI arguments to this program, --cfg <name>.json for input and output file names
	flag.Parse()
	fns := flag.Args()
	if len(fns) == 0	{
		fmt.Fprintf(os.Stderr, "Usage: ./read-json1 [-in in.json] [-out out.json]\n")
		os.Exit(1)
	}
	var Cfg = flag.String("cfg", fns[0], "config file for this call")
	if Cfg == nil	{
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

// read in data using ReadFile from in file argument
  readin, err := ioutil.ReadFile(fns[0])
  if err != nil	{
		fmt.Printf("error from ReadFile\n")
		os.Exit(1)
	}
// print file read
  json1 := IndentJSON(readin)
  fmt.Printf("from IndentJSON: %b\n", json1)

  // write data using ReadFile from in file argument
    err = ioutil.WriteFile(fns[1], readin, 0644)
    if err != nil	{
  		fmt.Printf("error from WriteFile\n")
  		os.Exit(1)
  	}
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
