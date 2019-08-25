package main

import "fmt"

func main() {
	// x := make(map[string]map[string]int)
	// y := make(map[string]int)
	// y["key"] = 10
	// x["a"] = y
	// fmt.Println(x)
	// name, ok := x["b"]
	// name2, notok := name["sas"]
	// fmt.Println(name, ok, len(name))
	// fmt.Println(name2, notok)
	//
	const gas = "GAS"
	const solid = "SOLID"

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": gas,
		},
		"He": {
			"name":  "Helium",
			"state": gas,
		},
		"Li": {
			"name":  "Lithium",
			"state": solid,
		},
		"Be": {
			"name":  "Beryllium",
			"state": solid,
		},
		"B": {
			"name":  "Boron",
			"state": solid,
		},
		"C": {
			"name":  "Carbon",
			"state": solid,
		},
		"N": {
			"name":  "Nitrogen",
			"state": gas,
		},
		"O": {
			"name":  "Oxygen",
			"state": gas,
		},
		"F": {
			"name":  "Fluorine",
			"state": gas,
		},
		"Ne": {
			"name":  "Neon",
			"state": gas,
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
