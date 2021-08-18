package main

import "fmt"

func main() {
	//NPC description
	longitude := 33.4484
	latitude := 112.0740
	isAgressive := false
	dialogAvailable := true
	name := "Bella Goth"
	class := "sim"

	fmt.Println(name, class, longitude, latitude, isAgressive, dialogAvailable)
}
