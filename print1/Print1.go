// Shawn Enriques
// senriqu1@uwyo.edu
// Assignment 1.8 print a file

package main

	import (
		"encoding/json"
		"fmt"
	)

	var IVar int
	var SVar string
	var I64Var int64
	var UIVar uint64

	type Example17 struct {
		A int
		B string
	}

	var E17 Example17
  var I64E17 int64							// Add int64 and uint64 types
	var U64E17 uint64

	var SliceOfString []string
	var MapOfString map[string]string
	var MapOfBool map[string]bool

	// init will initilize data before main() runs.  You can have more than one init() function.
	func init() {
		SliceOfString = make([]string, 0, 10)
		MapOfString = make(map[string]string)
		MapOfBool = make(map[string]bool)
	}

	func main() {
		SliceOfString = append(SliceOfString, "AAA", "BBB")
		MapOfString["mark"] = "first"
		MapOfString["twain"] = "last"
		MapOfBool["mark"] = true
		MapOfBool["twain"] = false

		fmt.Printf("IVar   = %d, type of IVar %T\n", IVar, IVar)
		fmt.Printf("SVar   = %v, type of SVar %T\n", SVar, SVar)

//  add prints for your int64 and uint64 types
		fmt.Printf("I64E17 = %d, type of IVar %T\n", I64E17, I64E17)
		fmt.Printf("U64E17 = %v, type of IVar %T\n", U64E17, U64E17)

		fmt.Printf("SVar   = %s, type of SVar %T\n", SVar, SVar)
		fmt.Printf("Address of SVar = %s, type of SVar %T\n", &SVar, &SVar)
		fmt.Printf("E17 = %s, type of E17 %T\n", &E17, &E17)
		fmt.Printf("    E17 = %+v, E17 as JSON: %s\n", &E17, IndentJSON(E17))

// add prints for the other types above - so you can see them printed out.
		fmt.Printf("I64Var = %d, type of IVar %T\n", I64Var, I64Var)
		fmt.Printf("UIVar = %v, type of IVar %T\n", UIVar, UIVar)
		fmt.Printf("SliceOfString = %s, type of IVar %T\n", SliceOfString, SliceOfString)
		fmt.Printf("MapOfString = %s, type of IVar %T\n", MapOfString, MapOfString)
		fmt.Printf("MapOfBool = %s, type of IVar %T\n", MapOfBool, MapOfBool)

// Print out each of them with the IndentJSON function.
JSliceOfString :=IndentJSON(SliceOfString)
fmt.Printf("JSliceOfString[JSON] = %s, type of IVar %T\n", JSliceOfString, JSliceOfString)
JMapOfString :=IndentJSON(MapOfString)
fmt.Printf("JMapOfString[JSON] = %s, type of IVar %T\n", JMapOfString, JMapOfString)
JMapOfBool :=IndentJSON(MapOfBool)
fmt.Printf("JMapOfBool[JSON] = %s, type of IVar %T\n", JMapOfBool, JMapOfBool)

	}

	func IndentJSON(v interface{}) string {
		s, err := json.MarshalIndent(v, "", "\t")
		if err != nil {
			return fmt.Sprintf("Error:%s", err)
		} else {
			return string(s)
		}
	}
