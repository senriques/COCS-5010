// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.5 Read and write JSON file

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type TransactionType struct	{
	TxHash		string
	TxIn			int
	TxOut			int
}

func main() {

// read in arguments for in and out files
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

//read in .JSON file, format and print it to the screen
	gCfg, err := ReadConfig(*Cfg)
	if err != nil	{
		fmt.Fprintf(os.Stderr, "Unable to read: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}
	fmt.Printf("Configuration: %+v\n", gCfg)
	fmt.Printf("JSON: %+v\n", IndentJSON(gCfg))
	for ii, ag := range fns	{
		if ii < len(fns)	{
			fmt.Printf("%s ", ag)
		} else {
			fmt.Printf("%s%", ag)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("JSON: %t %b\n", gCfg, gCfg)

// format the data to .JSON file and write output file named in argument 2
	bb, err := json.Marshal(gCfg)
	if err != nil	{
		fmt.Printf("json.Marshal failed to convert gCfg\n")
	}
	errout := ioutil.WriteFile(fns[1], bb, 0644)
	if errout != nil	{
	}
	return
}

//read the .JSON file and format the data
func ReadConfig(filename string) (rv TransactionType, err error)	{
	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil	{
		return TransactionType{}, err
	}
	err = json.Unmarshal(buf, &rv)
	if err != nil	{
		return TransactionType{}, err
	}
	return
}

//format the .JSON file tabulated for eas of screen reading
func IndentJSON(v interface{}) string	{
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil	{
		return fmt.Sprintf("Error: %s", err)
	} else {
		return string(s)
	}
}
