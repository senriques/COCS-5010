package main

// Shawn Enriques
// downloaded CustomMarshalJSON

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

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

func main() {

	unc main() {
	var Cfg = flag.String("cfg", "cfg.json", "config file for this call")

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Fprintf(os.Stderr, "Usage: ./echo2 [-cfg cfg.json] [arg1 ...]\n")
		os.Exit(1)
	}

	if Cfg == nil {
		fmt.Printf("--cfg is a required parameter\n")
		os.Exit(1)
	}

	gCfg, err := ReadConfig(*Cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read confguration: %s error %s\n", *Cfg, err)
		os.Exit(1)
	}

	fmt.Printf("Congiguration: %+v\n", gCfg)

	for ii, ag := range fns {
		if ii < len(fns) {
			fmt.Printf("%s ", ag)
		} else {
			fmt.Printf("%s", ag)
		}
	}
	fmt.Printf("\n")
}


	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}

