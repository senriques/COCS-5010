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
	"strings"
)

type Animal int

type ConfigData struct	{
	Name	string
	Year	string
	Value	string	"json:'Year'"
}

const (
	Unknown Animal = iota
	Gopher
	Zebra
)


func main() {

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

func IndentJSON(v interface{}) string	{
	s, err := json.MarshalIndent(v, "", "\t")
	if err != nil	{
		return fmt.Sprintf("Error: %s", err)
	} else {
		return string(s)
	}
}

func (a *Animal) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

