// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.4 echo command line arguments and parse arguments

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type ConfigData struct	{
	Name	string
	Year	string
	Value	string	"json:'Year'"
}

func main() {

// passe arguments to input file name
	flag.Parse()	// Parse CLI arguments to this program, --cfg <name>.json
	fns := flag.Args()
	if len(fns) == 0	{
		fmt.Fprintf(os.Stderr, "Usage: ./echo2 [-cfg cfg.json] arg1... \n")
		os.Exit(1)
	}
	var Cfg = flag.String("cfg", fns[0], "config file for this call")
	if Cfg == nil	{
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

// read in data from input file, format and print it to screen
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
}

// configures .JSON file data by unmarshaling into structure
func ReadConfig(filename string) (rv ConfigData, err error)	{
	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil	{
		return ConfigData{}, err
	}
	err = json.Unmarshal(buf, &rv)
	if err != nil	{
		return ConfigData{}, err
	}
	return
}

// formats data for ease of printing to screen
func IndentJSON(v interface{}) string	{
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil	{
		return fmt.Sprintf("Error: %s", err)
	} else {
		return string(s)
	}
}
