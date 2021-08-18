package main

import "fmt"

func main() {
	//first movie
	var title1 string = "Diplomatie"
	var length1 = 1.24
	var director1 string; director1 = "Volker Schlöndorff"
	var basedOnRealEvents1 = bool(true)
	year1 := 2014
	fmt.Println(title1, length1, director1, basedOnRealEvents1, year1)

        //second movie
	title2 := "Heartbreakers"
	length2 := 2.03
	director2 := "David Mirkin"
	basedOnRealEvents2 := false
	year2 := 2001
	fmt.Println(title2, length2, director2, basedOnRealEvents2, year2)
}
