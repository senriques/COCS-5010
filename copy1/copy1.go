// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.8 copy a file

// This code was taken from sample at https://shapeshed.com/copy-a-file-in-go/

package main

import (
  "io"
  "log"
  "os"
  "flag"
  "fmt"
)

func main() {

	flag.Parse()	// Parse CLI arguments to this program, --cfg <name>.json

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

  from, err := os.Open(fns[0])
  if err != nil {
    log.Fatal(err)
  }
  defer from.Close()

  to, err := os.OpenFile(fns[1], os.O_RDWR|os.O_CREATE, 0666)
  if err != nil {
    log.Fatal(err)
  }
  defer to.Close()

//  _, err = io.Copy(to, from)  original code
// modify code per instructions

	test1, err := io.Copy(to, from)

  if err != nil {
    log.Fatal(err)
  }
}
